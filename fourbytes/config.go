package fourbytes

import "time"

type Config struct {
	// Seconds, 0 means no timeout.
	Timeout time.Duration
}

func DefaultConfig() *Config {
	return &Config{0}
}
