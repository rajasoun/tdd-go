// Licensed under the Creative Commons License.

package sentry

import (
	"log"
	"testing"
	"time"

	"github.com/getsentry/sentry-go"

	"github.com/rajasoun/tdd-go/configurator"
)

func TestSentry(t *testing.T) {
	t.Run("Send Message To Sentry", func(t *testing.T) {
		err := sentry.Init(sentry.ClientOptions{
			Dsn: configurator.GetEnvWithKey("SENTRY_DSN"),
		})
		if err != nil {
			log.Fatalf("sentry.Init: %s", err)
		}
		// Flush buffered events before the program terminates.
		defer sentry.Flush(2 * time.Second)

		sentry.CaptureMessage("It works!")
	})
}
