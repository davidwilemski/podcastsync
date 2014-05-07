package podcast

// PodcastDownloadArgs contains nessecary data for the service to download a file and find out which users are subscribed to the originating feed
type PodcastDownloadArgs struct {
	FeedID      int    // podcastsync internal id of the originating feed
	FeedURL     string // Originating podcast feed's URL
	PodcastName string // podcast name: name of the folder the podcast is put into
	PodcastURL  string // URL with the media file's location
	AccessToken string // Dropbox access token for making API requets
}

// PodcastDownloadReply response containing success or failure and corresponding message
type PodcastDownloadReply struct {
	Message string
	Success bool
}
