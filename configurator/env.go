// Licensed under the Creative Commons License.

package configurator

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

// loadDotEnv : Loads env vars from a file - typically from .env.
var loadEnv = godotenv.Load

func LoadDotEnv() error {
	err := loadEnv(getBasePath() + ".env")
	return err
}

// Get OS Env Value Function - To Enable Mocking.
var osGetEnv = os.Getenv

// GetEnvWithKey : Get Environment Value with Key.
func GetEnvWithKey(key string) string {
	return osGetEnv(key)
}

// Get Working directory function - To Enable Mocking.
var osGetWd = os.Getwd

// Get current working dir.
func getCurrentWorkingDir() (string, error) {
	workingDir, err := osGetWd()
	if err != nil {
		return "", errors.New("could not get current working directory")
	}

	return workingDir, nil
}

// Get Base Path - Based on Environment.
func getBasePath() string {
	currentDir, _ := getCurrentWorkingDir()
	// remove base directory from the workingDir when run from test
	baseDir := filepath.Base(currentDir)

	isInTest := os.Getenv("GO_ENV") == "test"
	if isInTest {
		currentDir = strings.ReplaceAll(currentDir, baseDir, "")
	}
	return currentDir
}
