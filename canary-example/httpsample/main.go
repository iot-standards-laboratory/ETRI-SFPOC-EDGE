package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/segmentio/encoding/json"
)

func main() {

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
