package main

import (
	"encoding/json"
	"fmt"
)

type HTTPResult struct {
	//Code int
	//Msg  string
	code int
	msg  string
}

func main() {
	var results = []HTTPResult{
		{code: 200, msg: "ok"},
		{code: 206, msg: "okok"},
	}
	data, err := json.Marshal(results)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%s\n", data)
}
