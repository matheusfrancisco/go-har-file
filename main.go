package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vvakame/go-harlog"
)

func main() {
	err := realMain()
	if err != nil {
		panic(err)
	}
}

func realMain() error {
	har := &harlog.Transport{}
	hc := &http.Client{
		Transport: har,
	}

	{
		_, err := hc.Get("https://blog.vvaka.me/")
		if err != nil {
			return err
		}
	}


	b, err := json.MarshalIndent(har.HAR(), "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))

	return nil
}
