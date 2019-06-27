package pipotron // import "moul.io/pipotron/pipotron"

import (
	"bytes"
	"math/rand"
	"strings"
	"text/template"
)

func Generate(dict *Dict) (string, error) {
	return executeTemplate("{{pick .output}}", dict)
}

func executeTemplate(input string, dict *Dict) (string, error) {
	already_picked := map[string]interface{}{}

	funcMap := template.FuncMap{}
	pickFunc := func(opts []string) string {
		if len(opts) < 1 {
			return "$$$ INVALID OPTION $$$"
		}
		return opts[rand.Intn(len(opts))]
	}
	funcMap["title"] = strings.Title
	funcMap["lower"] = strings.ToLower
	funcMap["upper"] = strings.ToUpper
	funcMap["pick"] = pickFunc
	funcMap["pick_once"] = func(opts []string) string {
		// FIXME: find a better way to do this :)
		for i := 0; i < 100; i++ {
			picked := pickFunc(opts)
			if _, found := already_picked[picked]; !found {
				already_picked[picked] = nil
				return picked
			}
		}
		return "$$$ NO MORE UNIQUE PICKABLE ITEM $$$"
	}

	tmpl, err := template.New("").Funcs(funcMap).Parse(input)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer

	if err = tmpl.Execute(&tpl, dict); err != nil {
		return "", err
	}

	if tpl.String() == input {
		return input, nil
	}
	return executeTemplate(tpl.String(), dict)
}
