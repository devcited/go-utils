package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// Brand ...
type Brand struct {
	Token          string `json:"token"`
	Name           string `json:"name"`
	URL            string `json:"url"`
	Logo           string `json:"logo"`
	LogoDark       string `json:"logo_dark"`
	LogoMobile     string `json:"logo_mobile"`
	LogoMobileDark string `json:"logo_mobile_dark"`
}

// GetBrand ...
func GetBrand() Brand {
	jsonFile, err := ioutil.ReadFile(getSecretPath())
	if err != nil {
		fmt.Println(err)
	}

	var brand Brand

	err = json.Unmarshal(jsonFile, &brand)
	if err != nil {
		fmt.Println(err)
	}

	return brand
}

func getSecretPath() string {
	isDev := os.Getenv("NODE_ENV") == "development"

	secretPath := "/run/secrets/brand_secret"
	if isDev {
		root := os.Getenv("ROOT")
		token := os.Getenv("BRAND_TOKEN")
		secretPath = path.Join(root, "./projects", token, "./.configs/secrets/brand_secret.json")
	}

	return secretPath
}
