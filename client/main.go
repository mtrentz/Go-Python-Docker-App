package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func generator() string {
	// Generates a string with 5 random characters
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, 5)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Each second generates a random string of letters and POSTs into the server application
	for {
		time.Sleep(time.Second)
		randomStr := generator()

		var jsonStr = []byte(fmt.Sprintf(`{"path":"%s"}`, randomStr))

		// POST
		url := "http://testserver:8001"
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		if err != nil {
			panic(err)
		}
		req.Header.Set("X-Custom-Header", "Random String")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))

	}

}
