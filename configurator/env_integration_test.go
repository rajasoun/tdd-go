// Licensed under the Creative Commons License.

package configurator

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("GO_ENV", "test")
}

func TestLoadEnv(t *testing.T) {
	t.Run("Load .env File", func(t *testing.T) {
		err := LoadDotEnv()
		if err != nil {
			log.Println("Error Loading .env")
		}
		assert.NotContains(t, GetEnvWithKey("SENTRY_DSN"), "FILL_ME")
		assert.NotContains(t, GetEnvWithKey("GITGUARDIAN_API_KEY"), "FILL_ME")
		assert.NotContains(t, GetEnvWithKey("GITHUB_TOKEN"), "FILL_ME")
	})
	t.Run("Load invalid File", func(t *testing.T) {
		loadEnv = func(filenames ...string) (err error) {
			return errors.New("Error Loading .env")
		}
		err := LoadDotEnv()
		assert.Error(t, err)
	})
}
