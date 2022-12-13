package configsentry

import (
	"rishabhjha12-hub/constants"
	"time"

	"log"

	"github.com/getsentry/sentry-go"
)

func SentryConfig() {

	//sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://69c6e82b1f544d63980fca02d6d2485d@sentry.parkzap.com/18",

		ServerName:       "172.27.15.2",
		TracesSampleRate: constants.TracesSampleRate,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
		sentry.CaptureException(err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!!")

}
