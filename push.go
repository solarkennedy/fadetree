package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func sendPushNotification(title string, message string) {
	fmt.Printf("%s: %s", title, message)

	gotifyURL := getGotifyURL()
	if gotifyURL == "" {
		return
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Post(gotifyURL,
		"application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf("title=%s&message=%s", title, message)),
	)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(body))

}

func getGotifyURL() string {
	return os.Getenv("FADETREE_GOTIFY_URL")
}
