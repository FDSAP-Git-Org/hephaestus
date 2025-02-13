package test

import (
	"fmt"
	"testing"

	"github.com/FDSAP-Git-Org/hephaestus/encryption"
	"github.com/gofiber/fiber/v3"

	"github.com/stretchr/testify/require"
)

func TestProcess(t *testing.T) {
	// apilogs.SystemLog("testPath", "HealthTest", "FileName", "HEalth Checking", nil, nil)
	// hash, hashErr := encryption.GenerateHash("")
	// fmt.Println("HASH ERROR:", hashErr)
	// fmt.Println("HASH:", hash)

	// types := []string{utils.NumericString, utils.UpperString}
	// sKey := []byte("s3cr3tK3y")

	// ssKey := []byte("sadadasd")
	stringBody := fiber.Map{"retCode": 2233, "message": "test"}

	// resToken, jwterr := utils.GenerateJWTSignedString(sKey, 1, stringBody)
	// fmt.Println("JWT ERROR:", jwterr)
	// fmt.Println("TOKEN:", resToken)

	// claims, err := utils.ReadJWTToken(sKey, resToken)
	// fmt.Println("ERROR:", err)

	// mbody, _ := json.Marshal(claims)

	// fmt.Println("CLAIMS:", string(mbody))

	// fmt.Println("RESULT:", utils.VerifyAuthorization(resToken, sKey))

	secKey := "testSecret"
	result, _ := encryption.CreateSeal(stringBody, secKey)

	
	fmt.Println("TEST:", result)
	require.Equal(t, result, "")
}
