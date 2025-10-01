package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	BrokerURL   string
	TopicPrefix string
	ListenAddr  string
	DT          time.Duration // optional: if API needs to know sim dt
	// Add more as needed.
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func parseDurEnv(key string, def time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
		// Allow plain milliseconds
		if n, err := strconv.Atoi(v); err == nil {
			return time.Duration(n) * time.Millisecond
		}
	}
	return def
}

func Load() Config {
	return Config{
		BrokerURL:   getenv("BROKER_URL", "tcp://localhost:1883"),
		TopicPrefix: strings.TrimSuffix(getenv("TOPIC_PREFIX", "grid/v1"), "/"),
		ListenAddr:  getenv("LISTEN_ADDR", ":8080"),
		DT:          parseDurEnv("DT", 100*time.Millisecond),
	}
}
