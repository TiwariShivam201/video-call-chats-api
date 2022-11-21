package main

import (
	"math/rand"
	"os"
	API "shivamtiwari/github.com/agora/api"
	CONFIG "shivamtiwari/github.com/agora/config"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano()) // seed for random generator

	CONFIG.LoadConfig()

	if strings.EqualFold(os.Getenv("lambda"), "1") {
		API.StartServer(true)
	} else {
		API.StartServer(false)
	}
}
