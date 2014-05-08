package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"text/template"

	"github.com/davidwilemski/podcastsync/shared/jsonrpc"
	"github.com/davidwilemski/podcastsync/shared/podcast"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

type downloadReq struct {
	UID         string
	PodcastURL  string
	AccessToken string
}

func splash(c web.C, w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("PodcastSyncAPI/splash.html")
	t.Execute(w, "Hello World")
}

func podcastFileDownload(c web.C, w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", 400)
		return
	}

	var dl downloadReq
	err = json.Unmarshal(data, &dl)

	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	u, err := url.Parse(dl.PodcastURL)

	if err != nil || u.Scheme != "http" || u.Scheme != "https" {
		http.Error(w, "Podcast URL is invalid", 400)
		return
	}

	log.Println(u)

	go func() {
		var reply podcast.PodcastDownloadReply
		err := jsonrpc.Request("http://localhost:9999/", "PodcastDownloadService.Process", podcast.PodcastDownloadArgs{PodcastName: "SystemsLive", PodcastURL: dl.PodcastURL, AccessToken: dl.AccessToken}, &reply)
		if err != nil {
			log.Printf("Error with RPC call to PodcastDownloadService: %s\n", err)
			return
		}
		log.Printf("Response: %s\n", reply.Message)
	}()

	fmt.Fprintf(w, "success!")
}

func static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "PodcastSyncAPI/"+r.URL.Path[1:])
}

func main() {
	goji.Get("/", splash)
	goji.Post("/podcast/download", podcastFileDownload)

	goji.Get(regexp.MustCompile("^/static/(.+)$"), static)
	goji.Serve()
}
