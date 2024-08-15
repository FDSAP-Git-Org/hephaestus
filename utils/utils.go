package utils

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"

	"net/mail"
	"os"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v4"

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
	connTime := c.Context().ConnTime()
	return connTime.Format(time.DateTime)
}


// GenerateResponse generates a response object with the process time and request
// It takes in the response and the fiber context as parameters
// and returns an interface{} as the response object
func GenerateResponse(response interface{}, c fiber.Ctx) interface{} {
	// Get the process time from the fiber context
	processTime := GetResponseTime(c)

	// Create a new EPResponse object with the process time and request
	return EPResponse{
		ProcessTime: processTime, // Set the process time
		Request:     response,     // Set the request
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

// CustomClaims is a custom struct to store JWT claims
type CustomClaims struct {
	Body interface{} `json:"body"`
	jwt.RegisteredClaims
}

func ReadJWTToken(secretKey []byte, tokenString string) (*CustomClaims, error) {
	// Initialize a new instance of CustomClaims
	claims := &CustomClaims{}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Check if there was an error parsing the token
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
