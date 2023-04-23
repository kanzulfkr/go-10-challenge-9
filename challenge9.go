package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// Create seed
	rand.NewSource(time.Now().UnixNano())

	// Create JSON
	data := map[string]interface{}{
		"wind":  rand.Intn(100),
		"water": rand.Intn(100),
	}

	requestJson, err := json.Marshal(data)
	client := &http.Client{}

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(requestJson))

	// POST Request
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(string(body)), &result)

	if err != nil {
		log.Fatalln(err)
	}

	// Condition for Wind
	if result["water"].(float64) < 5 {
		fmt.Println("status water : aman")
	} else if result["water"].(float64) > 5 && result["water"].(float64) <= 8 {
		fmt.Println("status water : siaga")
	} else {
		fmt.Println("status water : bahaya")
	}

	// Condition for Water
	if result["wind"].(float64) <= 6 {
		fmt.Println("status wind : aman")
	} else if result["wind"].(float64) > 6 && result["wind"].(float64) <= 15 {
		fmt.Println("status wind : siaga")
	} else {
		fmt.Println("status wind : bahaya")
	}
}
