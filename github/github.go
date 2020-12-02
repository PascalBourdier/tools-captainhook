package github

import (
	"net/http"
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/webhooks.v5/github"
)

const (
	path = "/webhooks"
)

func onPush(w http.ResponseWriter, payload github.PushPayload) {
	if payload.Ref != "refs/heads/master" {
		log.Printf("No-op: ref '%s' isn't of interest", payload.Ref)
		return
	}

	log.Printf("Master was updated: '%s'", payload.Ref)
	// TODO
}

func WebhookReceiver() error {

	secret := os.Getenv("GITHUB_WEBHOOK_SECRET")
	if secret == "" {
		log.Println("Please provide a webhook secret with the GITHUB_WEBHOOK_SECRET environment variable.")
		return nil
	}
	hook, err := github.New(github.Options.Secret(secret))
	if err != nil {
		if err != github.ErrEventNotFound {
			log.Printf("Error: %#v", err)
		}
		return errors.Wrapf(err, "Unable to init the Github webhook")
	}

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.PushEvent)
		if err != nil {
			return
		}

		switch typedPayload := payload.(type) {
		case github.PushPayload:
			onPush(w, typedPayload)
		default:
			log.Printf("Unknown payload type.\n%#v\n", typedPayload)
		}
	})
	log.Println("Listening")
	err = http.ListenAndServe(":3000", nil)
	return err
}
