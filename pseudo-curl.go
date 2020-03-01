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
	var url string
	if os.Args[1][0:4] == "http" {
		url = os.Args[1]
	} else {
		// 決め打ちで申し訳ないけれど、とりあえず
		url = os.Args[2]
	}
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
		os.Exit(1)
	}
	body, reserr := ioutil.ReadAll(res.Body)
	if reserr != nil {
		log.Fatal(reserr)
		os.Exit(1)
	}
	fmt.Println(string(body))
	os.Exit(0)
}
