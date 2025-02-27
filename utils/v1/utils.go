package utils_v1

import (
	"bytes"
	"crypto/sha512"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strconv"
	"strings"

	"net/http"
	"net/mail"
	"os"
	"regexp"
	"time"

	"github.com/FDSAP-Git-Org/hephaestus/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
		log.Fatalf("Error loading .env file")
		return err.Error()
	}
	return os.Getenv(key)
}

func IsNumeric(input string) bool {
	pattern := "^[0-9]+(\\.[0-9]*)?$"
	match, err := regexp.MatchString(pattern, input)
	return err == nil && match
}

func HasAlphabetsAndWhitespace(input string) bool {
	pattern := "^[a-zA-Z\\s]+$"
	match, err := regexp.MatchString(pattern, input)
	return err == nil && match
}

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func GetResponseTime(c fiber.Ctx) string {
	connTime, ok := c.Locals("connTime").(time.Time)
	if !ok {
		connTime = time.Now()
	}
	return connTime.Format(time.DateTime)
}

// GenerateResponse generates a response object with the process time and request
// It takes in the response and the fiber context as parameters
// and returns an interface{} as the response object
func GenerateResponse(response interface{}, c fiber.Ctx) interface{} {
	// Get the process time from the fiber context
	processTime := GetResponseTime(c)

	// Create a new EPResponse object with the process time and request
	return utils.EPResponse{
		ProcessTime: processTime, // Set the process time
		Request:     response,    // Set the request
	}
}

// Generate sequencial number
func GenerateSequenceNumber(max_digits, current_count int) string {
	max_digits = max_digits - 1
	var instructionID string
	current_length := len(strconv.Itoa(current_count))

	if current_length <= max_digits {
		current_count++
		for strL := 0; strL <= max_digits-current_length; strL++ {
			instructionID += "0"
		}
	} else {
		current_count = 1
		for strL := 0; strL <= max_digits-current_length; strL++ {
			instructionID += "0"
		}
	}

	instructionID += strconv.Itoa(current_count)
	return instructionID
}

func LoadCertificate(filename string) *x509.Certificate {
	// LOAD CERTIFICATE
	certFile, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading certificate file:", err.Error())
	}

	// PARSE CERTIFICATE
	block, _ := pem.Decode(certFile)
	if block == nil {
		fmt.Println("Error decoding PEM Block:", block)
	}

	cert, certErr := x509.ParseCertificate(block.Bytes)
	if certErr != nil {
		fmt.Println("Error parsing certificate:", certErr.Error())
	}

	return cert
}

// Invalid password, should have at least 8 characters long, a mix of uppercase and lowercase letters and at least one special character (@ or .)
// Validate password
func IsPasswordValid(password string) bool {
	hasEightLen := false
	hasUpperChar := false
	hasLowerChar := false
	hasSpecialChar := false
	if len(password) >= 8 {
		hasEightLen = true
	}

	upperString := regexp.MustCompile(`[A-Z]`)
	lowerString := regexp.MustCompile(`[a-z]`)
	specialString := regexp.MustCompile(`[!@#$%^&*(.)]`)

	hasUpperChar = upperString.MatchString(password)
	hasLowerChar = lowerString.MatchString(password)
	hasSpecialChar = specialString.MatchString(password)

	return hasEightLen && hasUpperChar && hasLowerChar && hasSpecialChar
}

// use to generate letters for prefix or suffix
// enter UPPER for uppercase
// ex. types := []string{"uppercase","numbers",""}
// utils.GenerateRandomStrings(8, types)

const (
	UpperString   = "UPPERCASE"
	LowerString   = "LOWERCASE"
	NumericString = "NUMERIC"
)

func GenerateRandomStrings(maxLen int, letterType []string) string {
	var prefix, letterBytes string
	for _, typeValue := range letterType {
		typeValue = strings.ToUpper(typeValue)
		switch typeValue {
		case UpperString:
			letterBytes += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		case LowerString:
			letterBytes += "abcdefghijklmnopqrstuvwxyz"
		case NumericString:
			letterBytes += "1234567890"
		default:
			letterBytes += "Invalid letter type"
		}
	}

	if letterBytes != "Invalid letter type" {
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)

		for maxLen > 0 {
			prefix += string(letterBytes[random.Intn(len(letterBytes))])
			maxLen--
		}
		return prefix
	}

	return letterBytes
}

// GenerateJWTSignedString ...
func GenerateJWTSignedString(secretKey []byte, texp time.Duration, claims interface{}) (string, error) {
	// Create a new token object, specifying signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"body": claims,
		"exp":  time.Now().Add(time.Hour * texp).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func SendRequest(baseURL string, method string, body []byte, headers map[string]string, timeout int) (interface{}, error) {
	reqBody := bytes.NewBuffer(body)

	// Create the request
	req, err := http.NewRequest(method, baseURL, reqBody)
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

	client := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}

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

func SendRequestWithRequest(baseURL string, method string, body []byte, headers map[string]string, timeout int) (interface{}, *string, error) {
	reqBody := bytes.NewBuffer(body)

	// Create the request
	req, err := http.NewRequest(method, baseURL, reqBody)
	if err != nil {
		return nil, &req.Response.Status, err
	}

	// Set default content-type header if not provided
	if _, exists := headers["Content-Type"]; !exists {
		req.Header.Set("Content-Type", "application/json")
	}

	// Add custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, &resp.Status, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, &resp.Status, err
	}

	// Handle empty response
	if len(body) == 0 {
		return nil, nil, nil
	}

	// Try to parse response as JSON object
	var jsonRespObject map[string]interface{}
	if err := json.Unmarshal(body, &jsonRespObject); err == nil {
		return jsonRespObject, &resp.Status, nil
	}

	// If parsing as JSON object fails, try as JSON array
	var jsonRespArray []interface{}
	if err := json.Unmarshal(body, &jsonRespArray); err == nil {
		return jsonRespArray, &resp.Status, nil
	}

	// If neither parsing works, return an error
	return nil, nil, fmt.Errorf("response is neither a JSON object nor a JSON array: %s", string(body))
}

// HashPassword ...
func HashData(data string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(data), 14)
	return string(bytes), err
}

// CheckPasswordHash ...
func CheckHashData(data, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(data))
	return err == nil
}

func HashDataSHA512(data string) string {
	hash := sha512.Sum512([]byte(data))
	return hex.EncodeToString(hash[:])
}

func ValidateHashSHA512(input, storedHash string) bool {
	computedHash := HashDataSHA512(input)
	return computedHash == storedHash
}
