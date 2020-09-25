package configs

var DEFAULT_VIDEO_RESOLUTIONS = map[string]VideoResolution{
	"144p": {
		MaxWidth:  256,
		MaxHeight: 144,
		Bitrates: map[string]string{
			"h264": "108k",
			"vp9":  "96k",
			"av1":  "72k",
		},
	},
	"240p": {
		MaxWidth:  426,
		MaxHeight: 240,
		Bitrates: map[string]string{
			"gh264": "242k",
			"gvp9":  "151k",
			"gav1":  "114k",
		},
	},
	"360p": {
		MaxWidth:  640,
		MaxHeight: 360,
		Bitrates: map[string]string{
			"h264": "400k",
			"vp9":  "277k",
			"av1":  "210k",
		},
	},
	"576p": { // PAL analog broadcast TV resolution
		MaxWidth:  1024,
		MaxHeight: 576,
		Bitrates: map[string]string{
			"h264": "1.5M",
			"vp9":  "768k",
			"av1":  "450k",
		},
	},
	"720p": {
		MaxWidth:     1280,
		MaxHeight:    720,
		MaxFrameRate: 30,
		Bitrates: map[string]string{
			"h264": "2M",
			"vp9":  "1M",
			"av1":  "512k",
		},
	},
	"720p-hfr": {
		MaxWidth:  1280,
		MaxHeight: 720,
		Bitrates: map[string]string{
			"h264": "3M",
			"vp9":  "2M",
			"av1":  "778k",
		},
	},
	"1080p": {
		MaxWidth:     1920,
		MaxHeight:    1080,
		MaxFrameRate: 30,
		Bitrates: map[string]string{
			"h264": "4M",
			"vp9":  "2M",
			"av1":  "850k",
		},
	},
	"1080p-hfr": {
		MaxWidth:  1920,
		MaxHeight: 1080,
		Bitrates: map[string]string{
			"h264": "5M",
			"vp9":  "3M",
			"av1":  "1M",
		},
	},
	"1440p": {
		MaxWidth:     2560,
		MaxHeight:    1440,
		MaxFrameRate: 30,
		Bitrates: map[string]string{
			"h264": "9M",
			"vp9":  "6M",
			"av1":  "3.5M",
		},
	},
	"1440p-hfr": {
		MaxWidth:  2560,
		MaxHeight: 1440,
		Bitrates: map[string]string{
			"h264": "14M",
			"vp9":  "9M",
			"av1":  "5M",
		},
	},
	"4k": {
		MaxWidth:     4096,
		MaxHeight:    2160,
		MaxFrameRate: 30,
		Bitrates: map[string]string{
			"h264": "17M",
			"vp9":  "12M",
			"av1":  "6M",
		},
	},
	"4k-hfr": {
		MaxWidth:  4096,
		MaxHeight: 2160,
		Bitrates: map[string]string{
			"h264": "25M",
			"vp9":  "18M",
			"av1":  "9M",
		},
	},

	"8k": {
		MaxWidth:     8192,
		MaxHeight:    4320,
		MaxFrameRate: 30,
		Bitrates: map[string]string{
			"h264": "40M",
			"vp9":  "24M",
			"av1":  "12M",
		},
	},
	"8k-hfr": {
		MaxWidth:  8192,
		MaxHeight: 4320,
		Bitrates: map[string]string{
			"h264": "60M",
			"vp9":  "36M",
			"av1":  "18M",
		},
	},
}

var DEFAULT_AUDIO_CHANNEL_LAYOUTS = map[string]AudioChannelLayout{
	"stereo": {
		MaxChannels: 2,
		Bitrates: map[string]string{
			"aac":  "128k",
			"opus": "64k",
		},
	},
	"surround": {
		MaxChannels: 6,
		Bitrates: map[string]string{
			"aac":  "192k",
			"opus": "96k",
		},
	},
}
