package main

import (
	"context"
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

	applicationData, err := readApplicationData(applicationPath, config)
	if err != nil {
		panic(err)
	}

	//establish PubSub client and topic
	pubSubClient, err := EstablishpubSubClient(ctx, applicationData.Project, userAgent)
	if err != nil {
		panic(err)
	}

	pubSubTopic := EstablishPubSubTopic(ctx, pubSubClient, pubSubTopic)

	//establish GSMSecret client
	gsmClient, err := EstalisbGSMSecretClient(ctx, applicationData.Project)

	encryptionToken, err := RetrieveGSMSecret(ctx, gsmClient, applicationData.GsmSecretRef)

	payloadIsValid := ValidatePayload(r.Header, encryptionToken)

}
