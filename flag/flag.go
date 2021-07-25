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
	Args []string
}

func ParseFlags() (Flags, error) {

	flagSet := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	var flags Flags

	flagSet.StringVar(&flags.Host, "host", getStringEnv("GOX_HOST", "localhost"), "go examples host")
	flagSet.IntVar(&flags.Port, "port", getIntEnv("GOX_PORT", 8080), "go examples port")

	flagSet.Usage = func() {
		fmt.Fprint(flagSet.Output(), "Usage: flag [flags] [args]\n")
		flagSet.PrintDefaults()
	}

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return Flags{}, err
	}

	flags.Args = flagSet.Args()
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
