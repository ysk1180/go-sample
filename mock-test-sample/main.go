package main

import (
	"fmt"
	"go-sample/mock-test-sample/api"
)

func main() {
	client := &api.Client{
		HC: &api.HogeClient{},
	}
	res := client.Fetch()
	fmt.Println(res)
}
