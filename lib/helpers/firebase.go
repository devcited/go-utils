package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// FirebaseConfig ...
type FirebaseConfig struct {
	ApiKey            string `json:"apiKey"`
	AuthDomain        string `json:"authDomain"`
	ProjectID         string `json:"projectId"`
	StorageBucket     string `json:"storageBucket"`
	MessagingSenderId string `json:"messagingSenderId"`
	AppId             string `json:"appId"`
}

// GetFirebaseConfig ...
func GetFirebaseConfig() FirebaseConfig {
	jsonFile, err := ioutil.ReadFile(getFBSecretPath())
	if err != nil {
		fmt.Println(err)
	}

	var config FirebaseConfig

	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		fmt.Println(err)
	}

	return config
}

func getFBSecretPath() string {
	isDev := os.Getenv("NODE_ENV") == "development"

	secretPath := "/run/secrets/firebase_secret"
	if isDev {
		root := os.Getenv("ROOT")
		token := os.Getenv("BRAND_TOKEN")
		secretPath = path.Join(root, "./projects", token, "./.configs/secrets/firebase_secret.json")
	}

	return secretPath
}
