package utils_v2

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func SendRequest(baseURL string, method string, body []byte, headers map[string]string, queryParam map[string]interface{}, timeout int) (interface{}, error) {
	reqBody := bytes.NewBuffer(body)

	finalUrl := baseURL

	for qkey, qvalue := range queryParam {
		finalUrl = finalUrl + "?" + qkey + "=" + fmt.Sprint(qvalue)
	}

	fmt.Println("Final Url", finalUrl)
	// Create the request
	req, err := http.NewRequest(method, finalUrl, reqBody)
	if err != nil {
		return nil, err
	}

	// Set default content-type header if not provided
	if _, exists := headers["Content-Type"]; !exists {
		req.Header.Set("Content-Type", "application/json")
	}

	// Add custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{Timeout: time.Second * time.Duration(timeout)}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Handle empty response
	if len(body) == 0 {
		return nil, nil
	}

	// Try to parse response as JSON object
	var jsonRespObject map[string]interface{}
	if err := json.Unmarshal(body, &jsonRespObject); err == nil {
		return jsonRespObject, nil
	}

	// If parsing as JSON object fails, try as JSON array
	var jsonRespArray []interface{}
	if err := json.Unmarshal(body, &jsonRespArray); err == nil {
		return jsonRespArray, nil
	}

	// If neither parsing works, return an error
	return nil, fmt.Errorf("response is neither a JSON object nor a JSON array: %s", string(body))
}

func GenerateUIID(appName string) string {
	uiid := uuid.New()
	return fmt.Sprintf("%s-%s", appName, uiid.String())
}

func GenerateSeal(message any, signingKey string) string {
	convertedMessage, _ := json.Marshal(message)
	key := []byte(signingKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(convertedMessage))
	return hex.EncodeToString(h.Sum(nil))
}

func ValidateSeal(vendorKey, receivedSeal string, requestTimestamp string, requestBody any, signingKey string) bool {
	computedSeal := GenerateSeal(requestBody, signingKey)
	fmt.Println("COMPUTED SEAL:", computedSeal)
	// Use constant-time comparison to avoid timing attacks
	return hmac.Equal([]byte(computedSeal), []byte(receivedSeal))
}
