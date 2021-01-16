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

	j := fmt.Sprintf(`
	{
		"title": "%s",
		"message": "%s",
		"extras": {
		  "client::display": {
			"contentType": "text/markdown"
		  }
		}
	  }`, title, message)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Post(gotifyURL,
		"application/json",
		strings.NewReader(j),
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
