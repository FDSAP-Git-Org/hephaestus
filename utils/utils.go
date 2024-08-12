package utils

import (
	"fmt"
	"log"

	"net/mail"
	"os"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v3"

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

func GenerateResponse(response interface{}, c fiber.Ctx) interface{} {
	return EPResponse{
		ProcessTime: GetResponseTime(c),
		Response:    response,
	}
}