package helpers

import "os"

// Brand ...
type Brand struct {
	Token    string `json:"token"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	Logo     string `json:"logo"`
	LogoDark string `json:"logo_dark"`
}

// GetBrand ...
func GetBrand() Brand {
	brand := Brand{
		Token:    os.Getenv("BRAND_TOKEN"),
		Name:     os.Getenv("BRAND_NAME"),
		URL:      os.Getenv("BRAND_URL"),
		Logo:     os.Getenv("BRAND_LOGO"),
		LogoDark: os.Getenv("BRAND_LOGO_DARK"),
	}

	if brand.Token == "" {
		brand.Token = "localhost"
	}

	if brand.Name == "" {
		brand.Name = "localhost"
	}

	if brand.URL == "" {
		brand.URL = "http://localhost:8080"
	}

	if brand.Logo == "" {
		brand.Logo = "/statics/logo_default.svg"
	}

	if brand.LogoDark == "" {
		brand.LogoDark = "/statics/logo_white.svg"
	}

	return brand
}
