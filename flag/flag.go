package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Flags struct {
	Host string
	Port int
}

func ParseFlags() (Flags, error) {

	f := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	host := f.String("host", getStringEnv("GOX_HOST", "localhost"), "go examples host")
	port := f.Int("port", getIntEnv("GOX_PORT", 8080), "go examples port")
	if err := f.Parse(os.Args[1:]); err != nil {
		return Flags{}, err
	}

	flags := Flags{
		Host: stringValue(host),
		Port: intValue(port),
	}

	err := flags.validate()
	return flags, err
}

func (f Flags) validate() error {

	if f.Host == "" {
		return errors.New("host cannot be empty")
	}
	if f.Port < 1 || f.Port > 65535 {
		return fmt.Errorf("invalid port %d", f.Port)
	}
	return nil
}

func getStringEnv(envName string, defaultValue string) string {

	env, ok := os.LookupEnv(envName)
	if !ok {
		return defaultValue
	}
	return env
}

func getIntEnv(envName string, defaultValue int) int {

	env, ok := os.LookupEnv(envName)
	if !ok {
		return defaultValue
	}

	if intValue, err := strconv.Atoi(env); err == nil {
		return intValue
	}
	return defaultValue
}

func intValue(v *int) int {

	if v == nil {
		return 0
	}
	return *v
}

func stringValue(v *string) string {

	if v == nil {
		return ""
	}
	return *v
}
