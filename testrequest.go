package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type request struct {
	url  string
	body []byte
}

var url string

func SendRequest() {

	log.Println("Sending request to Cloudflare")

	query := map[string]string{
		"sql": "SELECT email FROM users",
	}
	body, _ := json.Marshal(query)

	c := http.Client{}

	bodybuffer := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", url, bodybuffer)
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Auth-Key", CLOUDFLARE_API_KEY)
	req.Header.Set("X-Auth-Email", CLOUDFLARE_ACCOUNT_EMAIL)
	req.Header.Set("Content-Type", "application/json")

	log.Printf("Sending request to %s\n", url)
	log.Printf("Headers: \n")
	for k, v := range req.Header {
		log.Printf("%s: %s\n", k, v)
	}

	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	log.Printf("Response status: %s\n", resp.Status)

	respurl := resp.Request.URL
	e := resp.Proto
	log.Printf("Response URL: %s\n", respurl)
	log.Printf("Response Proto: %s\n", e)

	// empty body

	// go-staticcheck ignore body
	body = []byte("")
	body, _ = io.ReadAll(resp.Body)

	log.Printf("Response: %s\n", body)

	hdrs := resp.Header
	for k, v := range hdrs {
		log.Printf("%s: %s\n", k, v)
	}

	time.Sleep(5 * time.Second)
	os.Exit(0)

}

func CreateBody() {
	log.Println("Listing databases")

}
