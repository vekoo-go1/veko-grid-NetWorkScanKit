package config

import (
	"strconv"
	"strings"
	"time"
)

// Config menyimpan konfigurasi untuk Veko Grid
type Config struct {
	InputFile   string
	OutputFile  string
	ProxyAddr   string
	UseTor      bool
	DelayRange  string
	Timeout     int
	DNSMode     string
	Silent      bool
	JSONOutput  bool
	Debug       bool
	MaxThreads  int
}

// GetDelayRange mengparsing delay range menjadi min dan max milliseconds
func (c *Config) GetDelayRange() (time.Duration, time.Duration, error) {
	parts := strings.Split(c.DelayRange, "-")
	if len(parts) != 2 {
		return 100 * time.Millisecond, 500 * time.Millisecond, nil
	}

	min, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 100 * time.Millisecond, 500 * time.Millisecond, nil
	}

	max, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 100 * time.Millisecond, 500 * time.Millisecond, nil
	}

	return time.Duration(min) * time.Millisecond, 
		   time.Duration(max) * time.Millisecond, nil
}

// GetTimeout mengkonversi timeout ke time.Duration
func (c *Config) GetTimeout() time.Duration {
	return time.Duration(c.Timeout) * time.Second
}

// IsTorEnabled mengecek apakah TOR diaktifkan
func (c *Config) IsTorEnabled() bool {
	return c.UseTor || strings.Contains(strings.ToLower(c.ProxyAddr), "tor")
}

// IsDoHEnabled mengecek apakah DNS over HTTPS diaktifkan
func (c *Config) IsDoHEnabled() bool {
	return strings.ToLower(c.DNSMode) == "doh"
}
