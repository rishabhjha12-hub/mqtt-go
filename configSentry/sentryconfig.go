package configsentry

import (
	"rishabhjha12-hub/constants"
	"time"

	"log"

	"github.com/getsentry/sentry-go"
)

func SentryConfig() bool {

	//sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn: constants.SentryDsn,

		ServerName:       constants.SentryServer,
		TracesSampleRate: constants.TracesSampleRate,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
		sentry.CaptureException(err)
	}

	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!!")
	return err != nil

}
