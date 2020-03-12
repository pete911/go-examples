package main

import (
	"os"
	"testing"
)

func TestDefaultFlags(t *testing.T) {

	rollback := setInput(nil, nil)
	defer rollback()

	flags, err := ParseFlags()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if flags.Port != 8080 {
		t.Errorf("port: want 8080, got %d", flags.Port)
	}
	if flags.Host != "localhost" {
		t.Errorf("host: want localhost, got %s", flags.Host)
	}
}

func TestFlags(t *testing.T) {

	rollback := setInput([]string{"flag",
		"--host", "test",
		"--port", "443",
	}, nil)
	defer func() { rollback() }()

	flags, err := ParseFlags()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if flags.Port != 443 {
		t.Errorf("port: want 443, got %d", flags.Port)
	}
	if flags.Host != "test" {
		t.Errorf("host: want test, got %s", flags.Host)
	}
}

func TestFlagsFromEnvVar(t *testing.T) {

	rollback := setInput([]string{"flag"}, map[string]string{
		"GOX_PORT": "80",
		"GOX_HOST": "env-host",
	})
	defer func() { rollback() }()

	flags, err := ParseFlags()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if flags.Port != 80 {
		t.Errorf("port: want 80, got %d", flags.Port)
	}
	if flags.Host != "env-host" {
		t.Errorf("host: want env-host, got %s", flags.Host)
	}
}

func TestFlagsOverrideEnvVar(t *testing.T) {

	rollback := setInput([]string{"flag",
		"--host", "flag-host",
		}, map[string]string{
		"GOX_PORT": "80",
		"GOX_HOST": "env-host",
	})
	defer func() { rollback() }()

	flags, err := ParseFlags()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if flags.Port != 80 {
		t.Errorf("port: want 80, got %d", flags.Port)
	}
	if flags.Host != "flag-host" {
		t.Errorf("host: want flag-host, got %s", flags.Host)
	}
}

// --- helper functions ---

func setInput(args []string, env map[string]string) (rollback func()) {

	osArgs := os.Args
	rollback = func() {
		os.Args = osArgs
		for k := range env {
			os.Unsetenv(k)
		}
	}

	if args == nil {
		args = []string{"test"}
	}

	os.Args = args
	for k, v := range env {
		os.Setenv(k, v)
	}
	return rollback
}
