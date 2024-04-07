package main

import (
	"context"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

var (
	configFile  string = "config.yaml"
	userAgent   string = "NA"
	pubSubTopic string = "NA"
)

func init() {

	functions.HTTP("Json2Pubsub", RequestHandler)
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	applicationPath := r.URL.Path

	config, err := getConfigStruct(configFile)
	if err != nil {
		panic(err)
	}
	log.Println("✅ Config yaml retrieved")

	applicationData, err := readApplicationData(applicationPath, config)
	if err != nil {
		panic(err)
	}
	log.Println("✅ application data retrieved")

	//establish PubSub client and topic
	pubSubClient, err := EstablishpubSubClient(ctx, applicationData.Project, userAgent)
	if err != nil {
		panic(err)
	}
	log.Println("✅ pubsub client retrieved")

	pubSubTopic := EstablishPubSubTopic(ctx, pubSubClient, pubSubTopic)

	//establish GSMSecret client
	gsmClient, err := EstalisbGSMSecretClient(ctx, applicationData.Project)
	if err != nil {
		panic(err)
	}
	log.Println("✅ gsm client retrieved")

	encryptionToken, err := RetrieveGSMSecret(ctx, gsmClient, applicationData.GsmSecretRef)
	if err != nil {
		panic(err)
	}
	log.Println("✅ gsm secret retrieved")

	payloadIsValid := ValidatePayload(r.Header, encryptionToken)

	if payloadIsValid {
		//Send payload to pubsub
	} else {
		log.Println("Unable to verify payload sent with path", r.URL.Path)
	}
}
