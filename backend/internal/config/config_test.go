package config

import "testing"

func TestLoadUsesDefaultHTTPAddr(t *testing.T) {
	t.Setenv(httpAddrEnv, "")

	cfg := Load()

	if cfg.HTTPAddr != defaultHTTPAddr {
		t.Fatalf("HTTPAddr = %q, want %q", cfg.HTTPAddr, defaultHTTPAddr)
	}
}

func TestLoadUsesHTTPAddrFromEnvironment(t *testing.T) {
	const want = "127.0.0.1:9090"
	t.Setenv(httpAddrEnv, want)

	cfg := Load()

	if cfg.HTTPAddr != want {
		t.Fatalf("HTTPAddr = %q, want %q", cfg.HTTPAddr, want)
	}
}
