package helpers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetTheme ...
func GetTheme(ctx *fiber.Ctx) string {
	theme := ctx.Cookies("USER_THEME")

	if theme == "" {
		theme = "light"
	}

	return theme
}

// GetEntry ...
func GetEntry() string {
	entry := os.Getenv("ENTRY")

	if entry == "" {
		entry = "/main.js"
	}

	return entry
}

// GetStorageHost ...
func GetStorageHost() string {
	host := os.Getenv("STORAGE_EMULATOR_HOST")

	if host == "" {
		host = "https://firebasestorage.googleapis.com"
	}

	return host
}

// ProlongSession ...
func ProlongSession(ctx *fiber.Ctx, UID string) {
	ctx.Cookie(&fiber.Cookie{
		Name:     "SES_UID",
		Value:    UID,
		HTTPOnly: true,
		SameSite: "Lax",
		Expires:  time.Now().Add(30 * 24 * time.Hour),
	})
}
