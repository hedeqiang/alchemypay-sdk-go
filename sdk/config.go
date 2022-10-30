package sdk

import (
	"time"
)

type Config struct {
	Scheme       string `json:"scheme"`
	Endpoint     string `json:"endpoint"`
	MerchantCode string `json:"merchantCode"`
	PrivateKey   string `json:"privateKey"`
	Timeout      time.Duration
}

// NewConfig returns a pointer of Config
// scheme only accepts http or https
// endpoint is the host to access, the connection could not be created if it's error
func NewConfig() *Config {
	return &Config{
		Scheme:  SchemeHTTP,
		Timeout: 30 * time.Second,
	}
}

func (c *Config) WithScheme(scheme string) *Config {
	c.Scheme = scheme
	return c
}

func (c *Config) WithEndpoint(endpoint string) *Config {
	c.Endpoint = endpoint
	return c
}

func (c *Config) WithMerchantCode(merchantCode string) *Config {
	c.MerchantCode = merchantCode
	return c
}

func (c *Config) WithPrivateKey(privateKey string) *Config {
	c.PrivateKey = privateKey
	return c
}

func (c *Config) WithTimeout(timeout time.Duration) *Config {
	c.Timeout = timeout
	return c
}
