package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"errors"
	"expvar"

	"code.google.com/p/goauth2/oauth"
	"github.com/davidwilemski/podcastsync/shared/podcast"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

var numHealthCalls = expvar.NewInt("numHealthCalls")
var numProcessCalls = expvar.NewInt("numProcess")

var dropboxAppKey = os.Getenv("PODCAST_DBAPPKEY")
var dropboxAppSecret = os.Getenv("PODCAST_DBAPPSECRET")

// PodcastDownloadService type for RPC methods used to download a file - and upload to dropbox
type PodcastDownloadService struct{}

// Health returns a dummy successful response for testing RPC service
func (p *PodcastDownloadService) Health(r *http.Request, args *podcast.PodcastDownloadArgs, reply *podcast.PodcastDownloadReply) error {
	numHealthCalls.Add(1)
	reply.Success = true
	reply.Message = "HI, I'm a podcast downloading service!"
	return nil
}

// Process does the real meat of this service
func (p *PodcastDownloadService) Process(r *http.Request, args *podcast.PodcastDownloadArgs, reply *podcast.PodcastDownloadReply) error {
	numProcessCalls.Add(1)
	reply.Success = false
	reply.Message = ""
	// first, fetch the file
	log.Println(args.PodcastURL)
	resp, err := http.Get(args.PodcastURL)
	if err != nil {
		return err
	}
	file := resp.Body

	c := dropboxOauthClient(args.AccessToken)

	path := strings.Split(args.PodcastURL, "/")
	if len(path) < 1 {
		return errors.New("args.PodcastURL is not actually a fodcast file URL")
	}

	contentDisposition := resp.Header.Get("Content-Disposition")
	var matchFileName = regexp.MustCompile(`filename="(.+)"`)

	matches := matchFileName.FindStringSubmatch(contentDisposition)
	var filename string
	if len(matches) < 2 {
		filename = path[len(path)-1]
	} else {
		filename = matches[1]
	}

	url := fmt.Sprintf("https://api-content.dropbox.com/1/files_put/sandbox/%s/%s", args.PodcastName, filename)
	resp, err = c.Post(url, "", file)
	if err != nil {
		return err
	}

	fmt.Printf("Dropbox file_put response: %s\n", resp.Status)
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Println(string(buff))
	log.Println(filename)
	reply.Success = true
	reply.Message = "Successful file download"
	return nil
}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(PodcastDownloadService), "")
	http.Handle("/", s)
	http.ListenAndServe(":"+"9999", nil)
}

func dropboxOauthClient(token string) *http.Client {

	var config = &oauth.Config{
		ClientId:     dropboxAppKey,
		ClientSecret: dropboxAppSecret,
		Scope:        "",
		AuthURL:      "https://www.dropbox.com/1/oauth2/authorize",
		TokenURL:     "https://api.dropbox.com/1/oauth2/token",
		RedirectURL:  "http://localhost:8888/dropbox_auth",
	}

	t := &oauth.Transport{Config: config}
	t.Token = &oauth.Token{AccessToken: token}
	return t.Client()

}
