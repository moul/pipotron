package pipotron // import "moul.io/pipotron/pipotron"

import (
	"bytes"
	"math/rand"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
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
		return opts[rand.Intn(len(opts))]
	}
	funcMap["randString"] = randStringBytes
	funcMap["title"] = strings.Title
	funcMap["lower"] = strings.ToLower
	funcMap["upper"] = strings.ToUpper
	funcMap["rand"] = rand.Float64
	funcMap["randIntn"] = rand.Intn
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
		return elem
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
