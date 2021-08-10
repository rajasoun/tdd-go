// Licensed under the Creative Commons License.

package monitor

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"

	"github.com/rajasoun/tdd-go/configurator"
)

func sentrySetup() {
	envErr := configurator.LoadDotEnv()
	if envErr != nil {
		log.Fatalf("configurator.LoadDotEnv: %s", envErr)
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: configurator.GetEnvWithKey("SENTRY_DSN"),
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	sentry.AddBreadcrumb(&sentry.Breadcrumb{
		Category: "Setup",
		Message:  "Integration Test",
		Level:    sentry.LevelInfo,
	})
	sentry.CaptureMessage("It works!")
	log.Println("Sent Event to Sentry")
}
