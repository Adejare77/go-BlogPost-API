package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func validateRequired(keys ...string) error {
	var missing []string

	for _, key := range keys {
		if strings.TrimSpace(os.Getenv(key)) == "" {
			missing = append(missing, key)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}
	return nil
}

func optional(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))

	if value == "" {
		return fallback
	}

	return value
}

func getEnvInt(key string, fallback int) int {
	value := strings.TrimSpace(os.Getenv(key))
	if  value == "" {
		return fallback
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("invalid value for %s: %q; using default %d", key, value, fallback)
		return fallback
	}

	return intValue
}

func getEnvDuration(key string, fallback string) time.Duration {
	value := strings.TrimSpace(os.Getenv(key))

	timeValue, err := time.ParseDuration(value)
	if err != nil {
		log.Printf("invalid value for %s: %q; using default %q", key, value, fallback)
		timeValue, _ = time.ParseDuration(fallback)
	}

	return timeValue
}
