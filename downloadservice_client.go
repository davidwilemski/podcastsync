package main

import (
	"fmt"
	"log"

	"github.com/davidwilemski/podcastsync/shared/jsonrpc"
	"github.com/davidwilemski/podcastsync/shared/podcast"
)

func main() {
	var reply podcast.PodcastDownloadReply
	err := jsonrpc.Request("http://localhost:9999/", "PodcastDownloadService.Health", podcast.PodcastDownloadArgs{}, &reply)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	fmt.Printf("Response: %s\n", reply.Message)
}
