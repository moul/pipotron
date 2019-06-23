package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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
	dictFile, err := dict.Box.Find(request.QueryStringParameters["dict"] + ".yml")
	if err != nil {
		return reply(request, 404, "", "No such dictionary (?dict=)")
	}

	var dict pipotron.Dict
	if err = yaml.Unmarshal(dictFile, &dict); err != nil {
		return reply(request, 500, "", fmt.Sprintf("%v", err))
	}

	out, err := pipotron.Generate(&dict)
	if err != nil {
		return reply(request, 500, "", fmt.Sprintf("%v", err))
	}

	if request.QueryStringParameters["show-source"] == "1" {
		return reply(request, 200, "", string(dictFile))
	}

	return reply(request, 200, "", out)
}

func main() {
	// rand.Seed(time.Now().UTC().UnixNano())
	lambda.Start(handler)
}
