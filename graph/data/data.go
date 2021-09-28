package data

import "pc-compositor/graph/model"

var playbacks = []*model.Playback{
	{
		ContentID: "1",
		Match:     false,
		PlaybackSet: []*model.PlaybackSet{
			{
				PlaybackURL:     "http://www.hotstar.com/videos/movies/english/captainmarvel/1",
				PlaybackCDNType: "INTERNAL",
			},
		},
	},
	{
		ContentID: "2",
		Match:     false,
		PlaybackSet: []*model.PlaybackSet{
			{
				PlaybackURL:     "http://www.hotstar.com/videos/movies/english/captainmarvel/2",
				PlaybackCDNType: "INTERNAL",
			},
		},
	},
}

func GetPlaybacks(contentid string) *model.Playback {
	for _, p := range playbacks {
		if p.ContentID == contentid {
			return p
		}
	}
	return nil
}
