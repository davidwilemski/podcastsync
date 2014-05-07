package main

import (
	"fmt"
	"log"

	"github.com/davidwilemski/podcastsync/shared/jsonrpc"
	"github.com/davidwilemski/podcastsync/shared/podcast"
)

func main() {
	var reply podcast.PodcastDownloadReply
	err := jsonrpc.Request("http://localhost:9999/", "PodcastDownloadService.Process", podcast.PodcastDownloadArgs{PodcastName: "SystemsLive", PodcastURL: "https://s3.amazonaws.com/SystemsLive/Episode42.mp3"}, &reply)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	fmt.Printf("Response: %s\n", reply.Message)
}
