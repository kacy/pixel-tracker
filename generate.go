package main

import (
	"bufio"
	"net/url"
	"os"
	"strings"
)

func cleanHost(host string) string {
	if len(host) == 0 {
		host = "localhost"
	}

	if host[:4] != "http" {
		host = "https://" + host
	}

	return host
}

func main() {
	host := os.Getenv("HOST")

	println("Generating the link...")

	reader := bufio.NewReader(os.Stdin)

	println("What's the email address?")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSuffix(email, "\n")

	println("What's the subject?")
	subject, _ := reader.ReadString('\n')
	subject = strings.TrimSuffix(subject, "\n")

	escapedSubject := url.QueryEscape(subject)

	host = cleanHost(host)

	link := host + "/p?email=" + email + "&subject=" + escapedSubject

	println("Here's the link:")
	println(link)
}
