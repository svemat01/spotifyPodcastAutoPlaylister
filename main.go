// This example demonstrates how to authenticate with Spotify using the authorization code flow.
// In order to run this example yourself, you'll need to:
//
//  1. Register an application at: https://developer.spotify.com/my-applications/
//     - Use "http://localhost:8080/callback" as the redirect URI
//  2. Set the SPOTIFY_ID environment variable to the client ID you got in step 1.
//  3. Set the SPOTIFY_SECRET environment variable to the client secret from step 1.
package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	"github.com/zmb3/spotify/v2/auth"
	"spotifyPodcastAutoplaylister/env"
	"spotifyPodcastAutoplaylister/spot"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found")
	}

	env.Validate()

	spot.Setup(spotifyauth.New(
		spotifyauth.WithRedirectURL(env.RedirectURI),
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate, spotifyauth.ScopePlaylistModifyPrivate, spotifyauth.ScopePlaylistModifyPublic),
		spotifyauth.WithClientID(env.ClientID),
		spotifyauth.WithClientSecret(env.ClientSecret)))

	items := make([]spotify.URI, 0)

	for showID, limit := range env.Shows {
		episodes, err := spot.Client.GetShowEpisodes(context.Background(), showID, spotify.Limit(limit))
		if err != nil {
			panic(err)
		}
		for _, episode := range episodes.Episodes {
			items = append(items, episode.URI)
			fmt.Printf("Added %s from %s to playlist\n", episode.Name, episode.Show.Name)
		}

		time.Sleep(100 * time.Millisecond)
	}

	_, err = spot.Client.ReplacePlaylistItems(context.Background(), env.PlaylistID, items...)
	if err != nil {
		panic(err)
	}

}
