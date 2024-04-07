package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	baseURL = "http://localhost:8801"
)

var (
	client = &http.Client{}
)

func main() {
	var (
		err          error
		resp         *http.Response
		vals         map[string]string
		sessionToken string
	)

	log.Info("----------------------------1. Performing health check ----------------------------")
	resp, err = performRequest(http.MethodGet, baseURL+"/health", nil, "")
	if err != nil {
		log.Warn("/health failed: ", err)
		return
	}
	if _, err = readResponse(resp, "message"); err != nil {
		log.Warn("Error reading response:", err)
		return
	}

	log.Info("----------------------------2. Registering user, getting OTP ----------------------------")
	hash := generateRandomHash(10)
	resp, err = performRequest(http.MethodPost, baseURL+"/user/register",
		map[string]string{
			"username":  fmt.Sprintf("%s_username", hash),
			"firstName": fmt.Sprintf("%s", hash),
			"lastName":  fmt.Sprintf("%s_last_name", hash),
			"email":     fmt.Sprintf("%s@gmail.com", hash),
			"password":  "12345678",
		}, "")
	if err != nil {
		log.Warn("/user/register failed: ", err)
		return
	}
	if vals, err = readResponse(resp, "otp"); err != nil {
		log.Warn("Error reading response:", err)
		return
	}

	log.Info("---------------------------- 3. Verifying OTP, getting JWT ----------------------------")
	resp, err = performRequest(http.MethodPost, baseURL+"/user/email/verify",
		map[string]string{
			"email":            fmt.Sprintf("%s@gmail.com", hash),
			"verificationCode": vals["otp"],
		}, "")
	if err != nil {
		log.Warn("/user/email/verify failed: ", err)
		return
	}
	if vals, err = readResponse(resp, "sessionToken"); err != nil {
		log.Warn("Error reading response:", err)
		return
	}
	sessionToken = vals["sessionToken"]

	log.Info("---------------------------- 3. Calling authenticated endpoint PATCH /me ----------------------------")
	resp, err = performRequest(http.MethodPatch, baseURL+"/me", nil, sessionToken)
	if err != nil {
		log.Warn("/user/email/verify failed: ", err)
		return
	}
	if vals, err = readResponse(resp, "sessionToken"); err != nil {
		log.Warn("Error reading response:", err)
		return
	}
	sessionToken = vals["sessionToken"]

}

// performRequest performs an HTTP request and returns the response or an error
func performRequest(method, url string, payload map[string]string, sessionToken string) (*http.Response, error) {
	log.Info(fmt.Sprintf("Calling %s %s", method, url))
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Warn("Error encoding JSON:", err)
		return nil, err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Warn("Error creating request: ", err)
		return nil, err
	}
	req.Header.Add("Authorization", sessionToken)

	return client.Do(req)
}

func readResponse(resp *http.Response, fieldName ...string) (map[string]string, error) {
	result := make(map[string]string)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Warn("Error reading response body:", err)
		return result, err
	}

	if resp.StatusCode >= http.StatusBadRequest && resp.StatusCode <= http.StatusNetworkAuthenticationRequired {
		log.Warn("statusCode: ", resp.StatusCode, " API call failed")
		log.Warn("respBody: ", string(body))
		return result, errors.New(resp.Status)
	}

	log.Info("Succeeded with: ", resp.StatusCode)

	var responseMap map[string]interface{}

	if err := json.Unmarshal(body, &responseMap); err != nil {
		log.Warn("Error parsing response body:", err)
		return result, err
	}

	for _, field := range fieldName {
		if value, ok := responseMap[field]; ok {
			result[field] = fmt.Sprintf("%v", value)
		}
	}
	return result, nil
}

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

func generateRandomHash(length int) string {
	hash := make([]byte, length)
	for i := range hash {
		hash[i] = charset[rand.Intn(len(charset))]
	}
	return string(hash)
}
