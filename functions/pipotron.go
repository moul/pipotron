package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ultreme/pipotron/dict"
	"github.com/ultreme/pipotron/pipotron"
	yaml "gopkg.in/yaml.v2"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	dictFile, err := dict.Box.Find(request.QueryStringParameters["dict"] + ".yml")
	if err != nil {
		return &events.APIGatewayProxyResponse{StatusCode: 404, Body: "No such dictionary (?dict=)"}, nil
	}

	var dict pipotron.Dict
	if err = yaml.Unmarshal(dictFile, &dict); err != nil {
		return &events.APIGatewayProxyResponse{StatusCode: 500, Body: fmt.Sprintf("%v", err)}, nil
	}

	out, err := pipotron.Generate(&dict)
	if err != nil {
		return &events.APIGatewayProxyResponse{StatusCode: 500, Body: fmt.Sprintf("%v", err)}, nil
	}

	return &events.APIGatewayProxyResponse{StatusCode: 200, Body: out}, nil
}

func main() {
	// rand.Seed(time.Now().UTC().UnixNano())
	lambda.Start(handler)
}
