package youtube

// OAuth2 scopes used by the YouTube Content ID Partner API.
const (
	// YoutubepartnerScope grants access to view and manage your assets and
	// associated content on YouTube.
	YoutubepartnerScope = "https://www.googleapis.com/auth/youtubepartner"

	// YoutubepartnerContentOwnerReadonlyScope grants read-only access to
	// view content owner account details from YouTube.
	YoutubepartnerContentOwnerReadonlyScope = "https://www.googleapis.com/auth/youtubepartner-content-owner-readonly"
)

// OAuth2 scopes used by the YouTube Data API v3.
const (
	// YoutubeScope grants access to manage your YouTube account.
	YoutubeScope = "https://www.googleapis.com/auth/youtube"

	// YoutubeForceSslScope grants access to manage your YouTube account,
	// requiring SSL for all requests.
	YoutubeForceSslScope = "https://www.googleapis.com/auth/youtube.force-ssl"

	// YoutubeReadonlyScope grants read-only access to your YouTube account.
	YoutubeReadonlyScope = "https://www.googleapis.com/auth/youtube.readonly"
)
