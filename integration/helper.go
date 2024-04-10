package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

func getField(resp *resty.Response, s string) string {
	var respBody map[string]string
	err := json.Unmarshal(resp.Body(), &respBody)
	if err != nil {
		log.Fatal("Failed to unmarshal response body with err: ", err)
	}
	return respBody[s]
}

func printTest(resp *resty.Response, err error) {
	fmt.Println("Request Info:")
	fmt.Println("  Endpoint: ", fmt.Sprintf("%s %s", resp.Request.Method, resp.Request.URL))
	if len(resp.Request.Header["Authorization"]) > 0 {
		fmt.Println("  Authorization : ", resp.Request.Header["Authorization"])
	}
	fmt.Println("  Body    :\n", resp.Request.Body)
	fmt.Println("Response Info:")
	fmt.Println("  Error  :", err)
	fmt.Println("  Status :", resp.Status())
	fmt.Println("  Time   :", resp.Time())
	fmt.Println("  Body   :\n", resp)
	fmt.Println()
}

func generateRandomHash(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	hash := make([]byte, length)
	for i := range hash {
		hash[i] = charset[rand.Intn(len(charset))]
	}
	return string(hash)
}
