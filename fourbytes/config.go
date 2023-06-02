package fourbytes

import "time"

type Config struct {
	// Seconds, 0 means no timeout.
	Timeout time.Duration
}

func (c *Config) validate() error {
	return nil
}

func DefaultConfig() *Config {
	return &Config{0}
}
