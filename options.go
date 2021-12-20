package radiogarden

import "net/http"

// WithFollowRedirectsDisabled prevents the underlying HTTP client from automatically
// following redirects. This is useful when getting a radio station's live
// broadcast stream URL via the following methods:
//
// - GetAraContentListenChannelIdChannelMp3
//
// - HeadAraContentListenChannelIdChannelMp3
//
// - GetAraContentListenChannelIdChannelMp3WithResponse
//
// - HeadAraContentListenChannelIdChannelMp3WithResponse
//
// These methods invoke an endpoint that responds with a redirect to the URL of
// an audio stream. When redirect following is enabled (this is the default
// behavior), the typed WithResponse methods produce errors because they attempt
// to consume and parse the entire response body.
func WithFollowRedirectsDisabled() ClientOption {
	return WithHTTPClient(
		&http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	)
}
