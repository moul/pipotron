package pipotron // import "moul.io/pipotron/pipotron"

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"moul.io/godev"
)

func Generate(ctx *Context) (string, error) {
	_, err := executeTemplate(`{{pick "init"}}`, ctx)
	if err != nil {
		return "", err
	}
	return executeTemplate(`{{pick "output"}}`, ctx)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func executeTemplate(input string, ctx *Context) (string, error) {
	funcMap := template.FuncMap{}
	for k, v := range sprig.FuncMap() {
		funcMap[k] = v
	}
	funcMap["pick"] = func(key string) string {
		opts := ctx.Dict[key]
		if len(opts) < 1 {
			return "$$$ INVALID OPTION $$$"
		}
		elem := opts[rand.Intn(len(opts))]
		evaluated, err := executeTemplate(elem, ctx)
		if err != nil {
			return elem
		}
		return evaluated
	}
	funcMap["randString"] = randStringBytes
	funcMap["title"] = strings.Title
	funcMap["lower"] = strings.ToLower
	funcMap["upper"] = strings.ToUpper
	funcMap["rand"] = rand.Float64
	funcMap["randIntn"] = rand.Intn
	funcMap["debug"] = func(v interface{}) string {
		fmt.Println(godev.PrettyJSON(v))
		return fmt.Sprintf("%#v", v)
	}
	funcMap["N"] = func(n int) (stream chan int) {
		stream = make(chan int)
		go func() {
			for i := 0; i <= n; i++ {
				stream <- i
			}
			close(stream)
		}()
		return
	}
	funcMap["randMinMax"] = func(min, max int) int { return rand.Intn(max-min) + min + 1 }
	funcMap["pick_once"] = func(key string) string {
		opts := ctx.Dict[key]
		if len(opts) < 1 {
			return "$$$ INVALID OPTION $$$"
		}
		i := rand.Intn(len(opts))
		elem := opts[i]
		opts = append(opts[:i], opts[i+1:]...)
		evaluated, err := executeTemplate(elem, ctx)
		if err != nil {
			return elem
		}
		return evaluated
	}

	tmpl, err := template.New("").Funcs(funcMap).Parse(input)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer

	if err = tmpl.Execute(&tpl, ctx); err != nil {
		return "", err
	}

	if tpl.String() == input {
		return input, nil
	}
	return executeTemplate(tpl.String(), ctx)
}
