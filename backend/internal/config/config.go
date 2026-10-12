package config

import "os"

const (
	httpAddrEnv     = "DOCMIND_HTTP_ADDR"
	defaultHTTPAddr = ":8080"
)

type Config struct {
	HTTPAddr string
}

func Load() Config {
	httpAddr := os.Getenv(httpAddrEnv)
	if httpAddr == "" {
		httpAddr = defaultHTTPAddr
	}

	return Config{
		HTTPAddr: httpAddr,
	}
}
