package fourbytes

import "time"

type Config struct {
	// Version has no meaning until a later version.
	Version string

	// Seconds, 0 means no timeout.
	Timeout time.Duration
}

func DefaultConfig() *Config {
	return &Config{Version, 0}
}
