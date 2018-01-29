package main

import (
	"log"
	"net/http"
	"gopkg.in/fsnotify.v1"
	"strings"
	"path/filepath"
	"fmt"
)

var config Config

func main() {
	config = DoConfig()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go readEventsFrom(watcher)

	for _, d := range config.VolumeDirs {
		log.Printf("Watching directory: %q", d)
		err = watcher.Add(d)
		if err != nil {
			log.Fatal(err)
		}
	}

	<-done
}

func readEventsFrom(watcher *fsnotify.Watcher) {
	for {
		select {
		case event := <-watcher.Events:
			if isMatch(event) {
				log.Println("Event :", event)
				err := triggerWebHook()
				if err != nil {
					log.Println(err)
				} else {
					log.Println("successfully triggered reload")
				}
			}

		case err := <-watcher.Errors:
			log.Println("error:", err)
		}
	}
}

func triggerWebHook() error {
	req, err := http.NewRequest(config.WebhookMethod, config.Webhook, nil)
	if err != nil {
		return fmt.Errorf("error: %s", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error: %s", err)
	}

	resp.Body.Close()

	if resp.StatusCode != config.WebhookStatusCode {
		return fmt.Errorf("error: Received response code %d, expected %d", resp.StatusCode, config.WebhookStatusCode)
	}

	return nil
}

func isMatch(event fsnotify.Event) bool {
	if event.Op&fsnotify.Chmod == fsnotify.Chmod {
		return false
	}

	if strings.HasPrefix(filepath.Base(event.Name), ".") {
		return false
	}

	return true
}
