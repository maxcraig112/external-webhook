package main

import (
	"context"

	"cloud.google.com/go/pubsub"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"google.golang.org/api/option"
)

func EstablishpubSubClient(ctx context.Context, projectID string, userAgent string) (*pubsub.Client, error) {
	client, err := pubsub.NewClient(ctx, projectID, option.WithUserAgent(userAgent))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func EstablishPubSubTopic(ctx context.Context, client *pubsub.Client, pubsubTopic string) *pubsub.Topic {

	topic := client.Topic(pubsubTopic)
	return topic
}

func EstalisbGSMSecretClient(ctx context.Context, projectID string) (*secretmanager.Client, error) {
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func RetrieveGSMSecret(ctx context.Context, client *secretmanager.Client, gsmRef string) (payload string, err error) {
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: gsmRef,
	}

	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		return "", err
	}
	return string(result.Payload.Data), nil
}

func ValidatePayload(header map[string][]string, token string) bool {
	//TODO
	return true
}
