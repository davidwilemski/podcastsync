package podcast

// PodcastDownloadArgs contains nessecary data for the service to download a file and find out which users are subscribed to the originating feed
type PodcastDownloadArgs struct {
	FeedID     int    // podcastsync internal id of the originating feed
	FeedURL    string // Originating podcast feed's URL
	PodcastURL string // URL with the media file's location
}

// PodcastDownloadReply response containing success or failure and corresponding message
type PodcastDownloadReply struct {
	Message string
	Success bool
}
