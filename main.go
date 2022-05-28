package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

const (
	defaultPlistFilename string = "GoogleService-Info.plist"
)

const (
	fbClientIDKey string = "FB_CLIENT_ID"

	fbReversedClientIDKey string = "FB_REVERSED_CLIENT_ID"

	fbAPIKeyKey string = "FB_API_KEY"

	gcmSenderIDKey string = "GCM_SENDER_ID"

	appBundleIDKey string = "APP_BUNDLE_ID"

	fbProjectIDKey string = "FB_PROJECT_ID"

	googleAppIDKey string = "GOOGLE_APP_ID"
)

const (
	clientIDValue string = "CLIENT_ID_HERE"

	revClientValue string = "REVERSED_CLIENT_ID_HERE"

	apiKeyValue string = "API_KEY_HERE"
)

type googleService struct {
	clientID    string
	revClientID string
	apiKey      string
	gcmSenderID string
	bundleID    string
	projectID   string
	googleAppID string
}

func newGoogleService() googleService {
	// try to load all variables from the current environment
	clientID := loadEnvVariable(fbClientIDKey)
	revClientID := loadEnvVariable(fbReversedClientIDKey)
	apiKey := loadEnvVariable(fbAPIKeyKey)
	gcmSenderID := loadEnvVariable(gcmSenderIDKey)
	bundleID := loadEnvVariable(appBundleIDKey)
	projectID := loadEnvVariable(fbProjectIDKey)
	googleAppID := loadEnvVariable(googleAppIDKey)
	return googleService{
		clientID:    clientID,
		revClientID: revClientID,
		apiKey:      apiKey,
		gcmSenderID: gcmSenderID,
		bundleID:    bundleID,
		projectID:   projectID,
		googleAppID: googleAppID,
	}
}

func execute(fptr *string) {
	// create a new struct from all environment variables
	// failing to load a single one will exit the program
	service := newGoogleService()

	var filename string
	if fptr != nil {
		filename = *fptr
	} else {
		log.Printf("No input plist specified. Using default name: %v\n", defaultPlistFilename)
		filename = defaultPlistFilename
	}
	// open plist file for reading and writing
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// parse each line in the plist file
		line := scanner.Text()
		rgxp := regexp.MustCompile(clientIDValue)
		result := rgxp.ReplaceAllString(line, service.clientID)
		fmt.Println(result)
	}
}

func loadEnvVariable(key string) string {
	if variable := os.Getenv(key); variable != "" {
		return variable
	}
	log.Fatalf("Unable to load environment variable: %v\n", key)
	return ""
}

func main() {
	var filePath *string
	// first argument is always the program name, so skip it
	argLength := len(os.Args[1:])
	if argLength >= 2 {
		*filePath = os.Args[1]
	}
	execute(filePath)
}
