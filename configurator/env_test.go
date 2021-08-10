// Licensed under the Creative Commons License.
package configurator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test Current Working Directory
func TestGetCurrentWorkingDir(t *testing.T) {
	t.Run("Check Current Dir Erors Out Gracefully", func(t *testing.T) {
		// Mocked function for os.Getwd and Update the var to this mocked function
		osGetWd = func() (string, error) {
			err := errors.New("Simulated error")

			return "", err
		}
		// This will return error
		_, err := getCurrentWorkingDir()
		assert.Error(t, err)
	})
	t.Run("Check Current Working Directory Path Contains .env", func(t *testing.T) {
		want := ".env"
		// Mocked function for os.Getwd
		osGetWd = func() (string, error) {
			return want, nil
		}
		got, _ := getCurrentWorkingDir()
		assert.Contains(t, got, want)
	})
}

func TestGetEnv(t *testing.T) {
	t.Run("Get Environment Variable via Mock", func(t *testing.T) {
		want := "dummy_token"
		// Mock function for os.Getenv
		osGetEnv = func(key string) string {
			return want
		}
		got := GetEnvWithKey("DUMMY_AUTH")
		assert.Equal(t, got, want)
	})
}
