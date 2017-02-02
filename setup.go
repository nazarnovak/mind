package main

import (
	"bufio"
	"io"
	"flag"
	"os"
	"strings"
)

var port *string

func LoadEnvFile() error {
	port = flag.String("port", "8080", "Server port")
	dEnvFile := ".env"
	envFile := flag.String("env-file", dEnvFile, "Environment file")
	flag.Parse()

	os.Setenv("PORT", *port)

	f, err := os.Open(*envFile)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		ln, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		parts := strings.SplitN(ln, "=", 2)
		if len(parts) != 2 {
			continue
		}
		err = os.Setenv(parts[0], strings.TrimSpace(parts[1]))
		if err != nil {
			return err
		}
	}

	return nil
}
