// Licensed under the Creative Commons License.

package monitor

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("GO_ENV", "test")
}

func TestSetup(t *testing.T) {
	t.Run("Send Message To Sentry", func(t *testing.T) {
		sentrySetup()
	})
}
