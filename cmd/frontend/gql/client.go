// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package gql

import (
	"context"
	"net/http"

	"github.com/Yamashou/gqlgenc/client"
)

type Client struct {
	Client *client.Client
}

func NewClient(cli *http.Client, baseURL string, options ...client.HTTPRequestOption) *Client {
	return &Client{Client: client.NewClient(cli, baseURL, options...)}
}

type Query struct {
	Greetings []string "json:\"greetings\" graphql:\"greetings\""
}
type Mutation struct {
	HelloWorld string "json:\"helloWorld\" graphql:\"helloWorld\""
}
type HelloWorld struct {
	HelloWorld string "json:\"helloWorld\" graphql:\"helloWorld\""
}
type GetGreetings struct {
	Greetings []string "json:\"greetings\" graphql:\"greetings\""
}

const HelloWorldDocument = `mutation HelloWorld ($input: Input!) {
	helloWorld(input: $input)
}
`

func (c *Client) HelloWorld(ctx context.Context, input Input, httpRequestOptions ...client.HTTPRequestOption) (*HelloWorld, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res HelloWorld
	if err := c.Client.Post(ctx, "HelloWorld", HelloWorldDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetGreetingsDocument = `query GetGreetings {
	greetings
}
`

func (c *Client) GetGreetings(ctx context.Context, httpRequestOptions ...client.HTTPRequestOption) (*GetGreetings, error) {
	vars := map[string]interface{}{}

	var res GetGreetings
	if err := c.Client.Post(ctx, "GetGreetings", GetGreetingsDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}