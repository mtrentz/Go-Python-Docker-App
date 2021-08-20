package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// // GET
	// resp, err := http.Get("http://localhost:8001")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer resp.Body.Close()

	// responseData, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// responseString := string(responseData)
	// fmt.Println(responseString)

	// POST
	url := "http://localhost:8001"
	var jsonStr = []byte(`{"path":" OPA TEXTO TESTE "}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Custom-Header", "myvalue")
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
