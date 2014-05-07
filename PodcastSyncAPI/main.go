package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
		return
	}

	var dl downloadReq
	err = json.Unmarshal(data, &dl)

	if err != nil {
		return
	}

	go func() {
		var reply podcast.PodcastDownloadReply
		err := jsonrpc.Request("http://localhost:9999/", "PodcastDownloadService.Process", podcast.PodcastDownloadArgs{PodcastName: "SystemsLive", PodcastURL: dl.PodcastURL, AccessToken: dl.AccessToken}, &reply)
		if err != nil {
			log.Fatal("dialing:", err)
		}
		log.Printf("Response: %s\n", reply.Message)
	}()

	fmt.Fprintf(w, "success!")
}

func main() {
	goji.Get("/", splash)
	goji.Post("/podcast/download", podcastFileDownload)
	goji.Serve()
}
