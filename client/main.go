package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
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

var wg sync.WaitGroup

func main() {

	wg.Add(10)

	// Each second generates a random string of letters and POSTs into the server application
	for i := 0; i <= 10; i++ {
		go func() {
			// time.Sleep(time.Second)
			randomStr := generator()

			// Sends a JSON with only one key:val pair. The key that Python app is waiting for is 'msg'.
			var jsonStr = []byte(fmt.Sprintf(`{"msg":"%s"}`, randomStr))

			// POST
			url := "http://testserver:8001"
			// url := "http://localhost:8001"
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
			wg.Done()
		}()
	}
	wg.Wait()

}
