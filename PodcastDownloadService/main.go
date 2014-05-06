package main

import (
	"net/http"

	"expvar"

	"github.com/davidwilemski/podcastsync/shared/podcast"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

var numHealthCalls = expvar.NewInt("numHealthCalls")

// PodcastDownloadService type for RPC methods used to download a file - and upload to dropbox
type PodcastDownloadService struct{}

// Health returns a dummy successful response for testing RPC service
func (p *PodcastDownloadService) Health(r *http.Request, args *podcast.PodcastDownloadArgs, reply *podcast.PodcastDownloadReply) error {
	numHealthCalls.Add(1)
	reply.Success = true
	reply.Message = "HI, I'm a feed parsing service!"
	return nil
}

// Process does the real meat of this service
func (p *PodcastDownloadService) Process(r *http.Request, args *podcast.PodcastDownloadArgs, reply *podcast.PodcastDownloadReply) error {
	// first, fetch the file
	reply.Success = true
	reply.Message = "HI, I'm a feed parsing service!"
	return nil
}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(PodcastDownloadService), "")
	http.Handle("/", s)
	http.ListenAndServe(":"+"9999", nil)
}
