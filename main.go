package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// base64Image is a 1x1 transparent pixel
const base64Image = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNkYAAAAAYAAjCB0C8AAAAASUVORK5CYII="

func checkEnv() {
	if len(os.Getenv("WEBHOOK")) == 0 {
		log.Fatal("WEBHOOK environment variable must be set")
	}
}

func sendPush(values url.Values) {
	endpoint := os.Getenv("WEBHOOK")

	email := values.Get("email")
	subject := values.Get("subject")

	// exit if email or subject are empty
	if len(email) == 0 || len(subject) == 0 {
		return
	}

	println("Sending push notification to " + email + " with subject " + subject)

	// send webhook
	resp, err := http.PostForm(endpoint, url.Values{
		"email":   {email},
		"subject": {subject},
	})

	if err != nil {
		log.Println(err.Error())
	}

	defer resp.Body.Close()
}

func main() {
	checkEnv()

	println("Starting server...")

	http.HandleFunc("/p", func(w http.ResponseWriter, r *http.Request) {
		go sendPush(r.URL.Query())

		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
		w.Header().Set("Expires", "Wed, 11 Nov 1999 11:11:11 GMT")
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
		w.Header().Set("Pragma", "no-cache")

		unbased, err := base64.StdEncoding.DecodeString(base64Image)
		if err != nil {
			log.Println("Error decoding string: ", err.Error())
			return
		}

		err = ioutil.WriteFile("image.png", unbased, 0666)
		if err != nil {
			log.Println("Error writing file: ", err.Error())
			return
		}

		http.ServeFile(w, r, "image.png")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
