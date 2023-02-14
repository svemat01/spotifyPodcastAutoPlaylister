package env

import (
	"fmt"
	"github.com/zmb3/spotify/v2"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var (
	// ClientID is the client ID for the application
	ClientID string
	// ClientSecret is the client secret for the application
	ClientSecret string
	// RedirectURI is the OAuth redirect URI for the application
	RedirectURI string
	// Shows is a map of show names to the number of episodes to download
	Shows map[string]int
	// Playlist ID is the ID of the playlist to add the episodes to
	PlaylistID spotify.ID
)

// Get and verify each environment variable
func Validate() {
	envErrors := make([]string, 0)

	ClientID = os.Getenv("SPOTIFY_ID")
	if ClientID == "" {
		envErrors = append(envErrors, "SPOTIFY_ID not set")
	}

	ClientSecret = os.Getenv("SPOTIFY_SECRET")
	if ClientSecret == "" {
		envErrors = append(envErrors, "SPOTIFY_SECRET not set")
	}

	RedirectURI = os.Getenv("SPOTIFY_REDIRECT_URI")
	// Check if the redirect URI is set and if it is a valid URI
	if RedirectURI == "" {
		envErrors = append(envErrors, "SPOTIFY_REDIRECT_URI not set")
	} else if _, err := url.Parse(RedirectURI); err != nil {
		envErrors = append(envErrors, "SPOTIFY_REDIRECT_URI is not a valid URI")
	}

	// Check if the shows environment variable is set
	shows := os.Getenv("SPOTIFY_SHOWS")
	if shows == "" {
		envErrors = append(envErrors, "SPOTIFY_SHOWS not set")
	} else {
		//	Shows format is the following: "show1:5,show2:10,show3:15"
		//	Where show1, show2, and show3 are the names of the shows and 5, 10, and 15 are the number of episodes to download
		//	We need to parse this into a map[string]int
		Shows = make(map[string]int)
		for _, show := range strings.Split(shows, ",") {
			s := strings.Split(show, ":")
			if len(s) != 2 {
				envErrors = append(envErrors, "SPOTIFY_SHOWS is not in the correct format")
				break
			}
			if i, err := strconv.Atoi(s[1]); err != nil {
				envErrors = append(envErrors, "SPOTIFY_SHOWS is not in the correct format")
				break
			} else {
				Shows[s[0]] = i
			}

		}
	}

	PlaylistID = spotify.ID(os.Getenv("SPOTIFY_PLAYLIST_ID"))
	if PlaylistID == "" {
		envErrors = append(envErrors, "SPOTIFY_PLAYLIST_ID not set")
	}

	// If there are any errors, print them and exit
	if len(envErrors) > 0 {
		for _, err := range envErrors {
			fmt.Printf("Error: %s\n", err)
		}
		os.Exit(1)
	}
}
