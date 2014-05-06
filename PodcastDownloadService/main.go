package main

import (
	"net/http"

	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

// PodcastDownloadService type for RPC methods used to download a file - and upload to dropbox
type PodcastDownloadService struct{}

// Health returns a dummy successful response for testing RPC service
func (p *PodcastDownloadService) Health(r *http.Request, args *podcastsync.PodcastDownloadArgs, reply *podcastsync.PodcastDownloadReply) error {
	reply.Success = true
	reply.Message = "HI, I'm a feed parsing service!"
	return nil
}

// Process does the real meat of this service
func (p *PodcastDownloadService) Process(r *http.Request, args *PodcastDownloadArgs, reply *PodcastDownloadReply) error {
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
