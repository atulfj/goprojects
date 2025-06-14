package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	url = "http://localhost:8081/add"
)

type CalcRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

func makeAddRequest() {
	c := CalcRequest{
		A: 2,
		B: 5,
	}

	body, err := json.Marshal(&c)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, _ = io.ReadAll(resp.Body)
	fmt.Println("response body: ", string(body))
}

func main() {
	makeAddRequest()
}
