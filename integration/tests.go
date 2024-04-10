package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
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
