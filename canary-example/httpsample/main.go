package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/segmentio/encoding/json"
)

func cookieSample() {

	obj := map[string]interface{}{
		"email":    "email@email.com",
		"password": "password",
	}

	b_obj, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	resp, _ := http.Post("http://localhost:8080/login", "application/json", bytes.NewReader(b_obj))
	fmt.Println(resp.StatusCode)
	cookie := resp.Header.Get("Set-Cookie")
	fmt.Println("cookie:", cookie)

}

func localAddrSample() {
	resp, err := http.Get("https://naver.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Request.URL)
}

func main() {
	localAddrSample()
}
