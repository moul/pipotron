package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	yaml "gopkg.in/yaml.v2"
	"moul.io/pipotron/dict"
	"moul.io/pipotron/pipotron"
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

	if request.QueryStringParameters["show-source"] == "1" {
		return &events.APIGatewayProxyResponse{StatusCode: 200, Body: string(dictFile)}, nil
	}

	return &events.APIGatewayProxyResponse{StatusCode: 200, Body: out}, nil
}

func main() {
	// rand.Seed(time.Now().UTC().UnixNano())
	lambda.Start(handler)
}
