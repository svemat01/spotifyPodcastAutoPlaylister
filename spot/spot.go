package spot

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"log"
	"spotifyPodcastAutoplaylister/spot/tokencache"
	"sync"
	"time"
)

const (
	state     = "abc123"
	tokenFile = "token.json"
)

var (
	Client     *spotify.Client
	auth       *spotifyauth.Authenticator
	TokenCache tokencache.Tokencache
	wg         sync.WaitGroup
)

func Setup(_auth *spotifyauth.Authenticator) {
	auth = _auth
	setupHttp()

	TokenCache = tokencache.New(tokenFile)
	token, err := TokenCache.Read()
	if err != nil {
		url := auth.AuthURL(state)
		fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)
		wg.Add(1)
		wg.Wait()

		newToken, err := Client.Token()
		if err != nil {
			panic(err)
		}
		TokenCache.Write(*newToken)
	} else {
		fmt.Printf("Configured Spotify access token, expires at %s\n", token.Expiry.Format(time.RFC1123))

		httpClient := auth.Client(context.Background(), token)
		fmt.Println("oAuth loaded")

		Client = spotify.New(httpClient)
		fmt.Println("Spot client made")

		newToken, err := Client.Token()
		if err != nil {
			fmt.Println("Failed getting token from yes")
			panic(err)
		}
		TokenCache.Write(*newToken)
	}

	fmt.Println("Loading user")
	user, err := Client.CurrentUser(context.Background())
	if err != nil {
		fmt.Println("Failed loading user")
		log.Fatal(err)
	}
	fmt.Println("Logged in as:", user.ID)
}
