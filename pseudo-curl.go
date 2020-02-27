package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"net/http"
	"crypto/tls"
)

func main() {
	url := os.Args[1]
	if len(url) == 0 {
		log.Fatal("first argument is required")
		os.Exit(1)
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	res, geterr := http.Get(url)
	if geterr != nil {
		log.Fatal(geterr)
		os.Exit(1)
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		log.Fatal("HTTP Response is not valid: " + res.Status)
	}
	body, reserr := ioutil.ReadAll(res.Body)
	if reserr != nil {
		log.Fatal(reserr)
		os.Exit(1)
	}
	fmt.Println(string(body))
	os.Exit(0)
}
