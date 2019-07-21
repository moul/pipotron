package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/ajg/form"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gohugoio/hugo/common/maps"
	yaml "gopkg.in/yaml.v2"
	"moul.io/pipotron/dict"
	"moul.io/pipotron/pipotron"
)

func reply(request events.APIGatewayProxyRequest, statusCode int, contentType string, body string) (*events.APIGatewayProxyResponse, error) {
	if contentType == "" {
		contentType = `text/plain; charset=utf-8`
	}
	ret := &events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers: map[string]string{
			"Content-Type": contentType,
		},
	}

	if callback := request.QueryStringParameters["callback"]; callback != "" {
		out, _ := json.Marshal(body)
		ret.Headers["Content-Type"] = `application/javascript; charset=utf-8`
		ret.Body = fmt.Sprintf("%s(%s)", callback, string(out))
	}

	return ret, nil
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var dictFile []byte
	var err error

	switch request.HTTPMethod {
	case "GET":
		dictFile, err = dict.Box.Find(request.QueryStringParameters["dict"] + ".yml")
		if err != nil {
			return reply(request, 404, "", "No such dictionary (?dict=)")
		}
	case "POST":
		type input struct {
			Dict   string `form:"dict"`
			Source string `form:"source"`
		}
		var i input
		d := form.NewDecoder(nil)
		if err := d.DecodeString(&i, request.Body); err != nil {
			return reply(request, 400, "", "invalid input")
		}
		dictFile = []byte(i.Source)
	default:
		return reply(request, 400, "", "No such method (GET or POST only)")
	}

	var context pipotron.Context
	if err = yaml.Unmarshal(dictFile, &context.Dict); err != nil {
		return reply(request, 500, "", fmt.Sprintf("%v", err))
	}
	context.Scratch = maps.NewScratch()

	out, err := pipotron.Generate(&context)
	if err != nil {
		return reply(request, 500, "", fmt.Sprintf("%v", err))
	}

	if request.QueryStringParameters["show-source"] == "1" {
		return reply(request, 200, "", string(dictFile))
	}

	contentType := `text/plain; charset=utf-8`
	if override := context.Scratch.Get("content-type"); override != nil {
		contentType = override.(string)
	}
	if override := request.QueryStringParameters["content-type"]; override != "" {
		contentType = override
	}
	return reply(request, 200, contentType, out)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lambda.Start(handler)
}
