package http

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Get(uri string, username, password interface{}) (io.ReadCloser, error) {
	var nop []byte
	body := ioutil.NopCloser(bytes.NewBuffer(nop))

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Printf("Create NewRequest error: %s", err.Error())
		return body, err
	}
	uname, ok1 := username.(string)
	pw, ok2 := password.(string)
	if ok1 && ok2 {
		req.SetBasicAuth(uname, pw)
	}
	req.Header.Set("Content-Type", "application/json")

	tc := &tls.Config{InsecureSkipVerify: true}
	tp := &http.Transport{TLSClientConfig: tc}
	client := &http.Client{Transport: tp}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Do Get error: %s", err.Error())
		return resp.Body, err
	}

	// TODO: 200, 201,...
	if resp.StatusCode != 200 {
		log.Printf("Get failed: %s", resp.Status)
		return resp.Body, errors.New(resp.Status)
	}
	return resp.Body, nil
}

func Post(url string, payload []byte, username, password interface{}) (io.ReadCloser, error) {
	var nop []byte
	body := ioutil.NopCloser(bytes.NewBuffer(nop))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("Create NewRequest error: %s", err.Error())
		return body, err
	}
	uname, ok1 := username.(string)
	pw, ok2 := password.(string)
	if ok1 && ok2 {
		req.SetBasicAuth(uname, pw)
	}
	req.Header.Set("Content-Type", "application/json")

	tc := &tls.Config{InsecureSkipVerify: true}
	tp := &http.Transport{TLSClientConfig: tc}
	client := &http.Client{Transport: tp}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Do Post error: %s", err.Error())
		return resp.Body, err
	}

	if resp.StatusCode != 200 {
		log.Printf("Post failed: %s", resp.Status)
		return resp.Body, errors.New(resp.Status)
	}
	return resp.Body, nil
}
