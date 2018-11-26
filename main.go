package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"encoding/json"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	out, _ := json.MarshalIndent(name, " ", "\t")
	fmt.Println("\n name", string(out))

	two, _ := json.MarshalIndent(ctx, " ", "\t")
	fmt.Println("\n ctx", string(two))


	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
