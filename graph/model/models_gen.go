// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Input struct {
	ContentID string `json:"content_id"`
}

type Playback struct {
	ContentID        string         `json:"content_id"`
	Match            bool           `json:"match"`
	DrmClass         string         `json:"drm_class"`
	DownloadDrmClass string         `json:"download_drm_class"`
	PlaybackSet      []*PlaybackSet `json:"playback_set"`
}

type PlaybackSet struct {
	PlaybackURL     string `json:"PlaybackURL"`
	PlaybackCDNType string `json:"PlaybackCDNType"`
}
