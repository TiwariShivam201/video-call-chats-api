package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// LoadConfig - load .env file from given path for local, else will be getting from env var
func LoadConfig() {
	// load .env file from given path for local, else will be getting from env var
	if !strings.EqualFold(os.Getenv("prod"), "true") {
		err := godotenv.Load(".test-env")
		if err != nil {
			panic("Error loading .env file")
		}
	}

	Log, _ = strconv.ParseBool(os.Getenv("LOG"))
	Migrate, _ = strconv.ParseBool(os.Getenv("MIGRATE"))
	AGORA_APP_ID = os.Getenv("AGORA_APP_ID")
	AGORA_APP_CERTIFICATE = os.Getenv("AGORA_APP_CERTIFICATE")
	AWSAccesKey = os.Getenv("AWSACCESSKEY")
	AWSSecretKey = os.Getenv("AWSSECRETKEY")
	AWSRegion = os.Getenv("AWSREGION")

}
