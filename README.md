# Spotify Podcast AutoPlaylister

The **Spotify Podcast AutoPlaylister** is a GO application that automates the process of fetching the latest episodes of a list of podcasts and adding them to a Spotify playlist.

## Installation

To install the program, you can either download the binary from the releases page or build it yourself using the Go build command.

### Building the Program

To build the program yourself, follow these steps:

1. Clone the repository
2. Navigate to the root directory of the cloned repository
3. Run the command `go build`

## Usage

The **Spotify Podcast AutoPlaylister** uses environment variables to configure itself. These variables can either be set in a `.env` file in the same directory as the binary or set in the environment.

The following variables are required for the program to run:

- `SPOTIFY_ID`: The client ID for the Spotify OAuth2 API
- `SPOTIFY_SECRET`: The client secret for the Spotify OAuth2 API
- `SPOTIFY_REDIRECT_URI`: The redirect URI for the Spotify OAuth2 API
- `SPOTIFY_SHOWS`: The list of shows to fetch episodes for. This should be in the format `SHOW_ID:COUNT` where `COUNT` is the number of episodes to fetch.
- `SPOTIFY_PLAYLIST_ID`: The ID of the Spotify playlist to add the fetched episodes to.

### Example .env File

```env
SPOTIFY_ID=your_client_id
SPOTIFY_SECRET=your_client_secret
SPOTIFY_REDIRECT_URI=http://localhost:8080/callback

# Format: SHOW_ID:COUNT (COUNT is the number of episodes to fetch)
SPOTIFY_SHOWS=12345:5,67890:10

SPOTIFY_PLAYLIST_ID=your_playlist_id
```

## Automating the Program

You can automate the **Spotify Podcast AutoPlaylister** by setting up a cron job. To do this, create a cron job that runs the program every day at a specific time.

### Example Cron Job
```bash
0 21 * * * /path/to/spotify-podcast-autoplaylister
```

This will run the program every day at 9:00 PM.

