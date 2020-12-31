package config

import (
	"bufio"
	"os"
	"strings"
)

// EnvConfig represents all environment variables used by this app.
type EnvConfig struct {

	// App
	AppName string `envconfig:"APP_NAME" required:"true"`

	// AWS
	AWSRegion  string `envconfig:"AWS_REGION" required:"true"`
	AWSProfile string `envconfig:"AWS_PROFILE" required:"true"`
}

// ExportEnvVars will export environment variables from given file.
func ExportEnvVars(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "export ") {
			line = line[7:]
		}

		token := strings.SplitN(line, "=", 2)
		// Remove spaces, ' and "
		key := strings.Trim(strings.Trim(strings.TrimSpace(token[0]), "'"), "\"")
		value := strings.Trim(strings.Trim(strings.TrimSpace(token[1]), "'"), "\"")

		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
