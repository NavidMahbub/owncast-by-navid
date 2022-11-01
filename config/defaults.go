package config

import (
	"os"
	"time"

	"github.com/owncast/owncast/models"
)

// Defaults will hold default configuration values.
type Defaults struct {
	Name                 string
	Title                string
	Summary              string
	ServerWelcomeMessage string
	Logo                 string
	Tags                 []string
	PageBodyContent      string

	DatabaseFilePath string
	WebServerPort    int
	WebServerIP      string
	RTMPServerPort   int
	StreamKey        string

	YPEnabled bool
	YPServer  string

	SegmentLengthSeconds int
	SegmentsInPlaylist   int
	StreamVariants       []models.StreamOutputVariant

	FederationUsername      string
	FederationGoLiveMessage string

	ChatEstablishedUserModeTimeDuration time.Duration
}

// GetDefaults will return default configuration values.
func GetDefaults() Defaults {

	baseUrl := os.Getenv("REMOTE_LEARNING_BASE_URL")

	return Defaults{
		Name:                 "IPEMIS STUDIO",
		Title:                "Your IPEMIS STUDIO Server",
		Summary:              "This is brief summary of whom you are or what your stream is. You can edit this description in the admin.",
		ServerWelcomeMessage: "",
		Logo:                 "logo.svg",
		Tags: []string{
			"DPE",
			"IPEMIS",
		},

		PageBodyContent: "# This is your page content that can be edited from the admin.",

		DatabaseFilePath: "data/owncast.db",

		YPEnabled: false,
		YPServer:  "https://directory.owncast.online",

		WebServerPort:  8080,
		WebServerIP:    baseUrl,
		RTMPServerPort: 1935,
		StreamKey:      "ZHNpbm5vdmF0b3Jz",

		ChatEstablishedUserModeTimeDuration: time.Minute * 15,

		StreamVariants: []models.StreamOutputVariant{
			{
				Name:          "360p",
				VideoBitrate:  1200,
				ScaledWidth:   360,
				ScaledHeight:  360,
				Framerate:     24,
				CPUUsageLevel: 2,
			},
			{
				Name:          "720p",
				VideoBitrate:  3000,
				ScaledWidth:   1280,
				ScaledHeight:  720,
				Framerate:     24,
				CPUUsageLevel: 2,
			},
			{
				Name:          "1080p",
				VideoBitrate:  4000,
				ScaledWidth:   1920,
				ScaledHeight:  1080,
				Framerate:     30,
				CPUUsageLevel: 2,
			},
		},

		FederationUsername:      "streamer",
		FederationGoLiveMessage: "I've gone live!",
	}
}
