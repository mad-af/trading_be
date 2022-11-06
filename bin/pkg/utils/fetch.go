package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	r "trading_be/bin/pkg/response"
	"trading_be/config"
)

type FetchRequest struct {
	Method        string
	Url           string
	Authorization string
	Body          interface{}
}

func Fetch(p *FetchRequest) ([]byte, error) {
	var b, _ = json.Marshal(p.Body)
	var request, err = http.NewRequest(p.Method, p.Url, bytes.NewBuffer(b))
	if err != nil {
		return nil, r.ReplyError("Url not found", http.StatusConflict)
	}

	var client = &http.Client{}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", p.Authorization)
	response, err := client.Do(request)
	if err != nil {
		return nil, r.ReplyError("Failed to fetch url", http.StatusConflict)
	}
	defer response.Body.Close()

	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, r.ReplyError("Failed to fetch url", http.StatusConflict)
	}

	return res, nil
}

func FetchModule(p *FetchRequest) (r.ReplySend, error) {
	var data r.ReplySend
	p.Url = fmt.Sprintf("%s%s", config.Env.ApiUrl, p.Url) 
	var fetch, err = Fetch(p)
	if err != nil {
		return data, err
	}

	json.Unmarshal(fetch, &data)
	return data, nil	
}