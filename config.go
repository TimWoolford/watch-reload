package main

import (
	"flag"
	"log"
	"os"
)
var volumeDirs volumeDirsFlag

var webhook = flag.String("webhook-url", "", "the url to send a request to when the specified config map volume directory has been updated")
var webhookMethod = flag.String("webhook-method", "POST", "the HTTP method url to use to send the webhook")
var webhookStatusCode = flag.Int("webhook-status-code", 200, "the HTTP status code indicating successful triggering of reload")

type Config struct {
	VolumeDirs volumeDirsFlag
	Webhook string
	WebhookMethod string
	WebhookStatusCode int
}

func DoConfig() Config {

	flag.Var(&volumeDirs, "volume-dir", "the config map volume directory to watch for updates; may be used multiple times")
	flag.Parse()

	if len(volumeDirs) < 1 {
		badArgs("Missing volume-dir")
	}

	if *webhook == "" {
		badArgs("Missing webhook")
	}

	return Config{
		VolumeDirs: volumeDirs,
		Webhook: *webhook,
		WebhookMethod: *webhookMethod,
		WebhookStatusCode: *webhookStatusCode,
	}
}

func badArgs(msg string) {
	log.Println(msg)
	log.Println()
	flag.Usage()
	os.Exit(1)
}
