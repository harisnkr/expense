package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

const (
	baseURL = "http://localhost:8801"
)

var (
	client = resty.New()
)

func main() {
	var (
		hash = generateRandomHash(10)
		resp *resty.Response
		err  error
	)

	// performing health check
	resp, err = client.R().EnableTrace().
		Get(baseURL + "/health")
	printTest(resp, err)

	// enter email and get OTP
	resp, err = client.R().EnableTrace().
		SetBody(fmt.Sprintf(`{
			"username": "%s_username",
			"firstName": "%s",
			"lastName": "%s",
			"password": "12345678",
			"email": "%s@gmail.com"
		}`, hash, hash, hash, hash)).
		Post(baseURL + "/user/register")
	printTest(resp, err)

	// internal endpoint to get OTP for email
	resp, err = client.R().EnableTrace().
		SetQueryParams(map[string]string{"email": hash + "@gmail.com"}).
		Get(baseURL + "/internal/user/otp")
	printTest(resp, err)
	otp := getField(resp, "otp")

	resp, err = client.R().EnableTrace().
		SetBody(fmt.Sprintf(`{"email": "%s@gmail.com","verificationCode" : "%s"}`, hash, otp)).
		Post(baseURL + "/user/email/verify")
	printTest(resp, err)
	sessionToken := getField(resp, "sessionToken")

	resp, err = client.R().EnableTrace().
		SetHeader("Authorization", sessionToken).
		Patch(baseURL + "/me")
	printTest(resp, err)
}

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
