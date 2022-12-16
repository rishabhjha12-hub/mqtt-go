package constants

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Gate string

var Redisadress string
var RedisPass string
var BrokerIpp string
var Brokerport string
var SentryDsn string
var Fastag_Status_Url string
var SentryServer string

//reference https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func SetEnv() {

	gate := goDotEnvVariable("GATE")

	log.Printf("godotenv : %s = %s \n", "GATE", gate)

	redisaddress := goDotEnvVariable("REDIS_ADD")
	log.Printf("godotenv : %s = %s \n", "redis", redisaddress)

	sentrydsn := goDotEnvVariable("SENTRY_DSN")
	log.Printf("godotenv : %s = %s \n", "SENTRY", sentrydsn)

	redispass := goDotEnvVariable("REDIS_PASSWORD")
	log.Printf("godotenv : %s = %s \n", "redispas", redispass)

	brokeripp := goDotEnvVariable("BROKER_IP")
	log.Printf("godotenv : %s = %s \n", "brokeripp", brokeripp)

	brokerport := goDotEnvVariable("BROKER_PORT")
	log.Printf("godotenv : %s = %s \n", "brokerport", brokerport)

	SentryDsn := goDotEnvVariable("SENTRY_DSN")
	log.Printf("godotenv : %s = %s \n", "sentrydsn", SentryDsn)

	SentryServer := goDotEnvVariable("SENTRY_SERVER_NAME")
	log.Printf("godotenv : %s = %s \n", "sentrydsn", SentryServer)

	Fastag_Status_Url := goDotEnvVariable("Fastag_Status_Url")
	log.Printf("godotenv : %s = %s \n", "FastagStatusUrl", Fastag_Status_Url)

	BrokerIpp = brokeripp
	Brokerport = brokerport
	Gate = gate

	Redisadress = redisaddress
	RedisPass = redispass

}
