package pipotron // import "moul.io/pipotron/pipotron"

import (
	"bytes"
	"fmt"
	"math/rand"
	"text/template"
)

func Generate(dict *Dict) (string, error) {
	return executeTemplate("{{pick .output}}", dict)
}

func executeTemplate(input string, dict *Dict) (string, error) {
	funcMap := template.FuncMap{
		"pick": func(opts []string) string {
			if len(opts) < 1 {
				return fmt.Sprintf("$$$ INVALID OPTION $$$")
			}
			return opts[rand.Intn(len(opts))]
		},
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
