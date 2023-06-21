package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

func main() {
	// Set up the client
	for i := 0; i < 5; i++ {
		client := &http.Client{
			Transport: &transport.APIKey{Key: "AIzaSyANLk8H2wUQjcfzhAjFCqLj8_u5zQC8Eho"},
		}

		// Create a new YouTube service
		youtubeService, err := youtube.New(client)
		if err != nil {
			fmt.Printf("Error creating YouTube client: %v", err)
			os.Exit(1)
		}

		// Define the search query
		// Define the search query
		var part []string
		searchQuery := youtubeService.Search.List(part).
			Q("Top 5 Search Engine #coding #fortnite").
			Type("video").
			Order("date").
			PublishedAfter(time.Now().Add(-24 * time.Hour).Format(time.RFC3339))

		// Perform the search
		searchResponse, err := searchQuery.Do()
		if err != nil {
			fmt.Printf("Error searching for videos: %v", err)
			os.Exit(1)
		}

		// Check if any videos were found
		if len(searchResponse.Items) == 0 {
			fmt.Println("No videos found.")
			return
		}

		// Get the ID of the first video in the search results
		videoID := searchResponse.Items[0].Id.VideoId

		// Open the video in the default web browser

		err = openURL(fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID))
		if err != nil {
			fmt.Printf("Error opening video: %v", err)
			os.Exit(1)
		}

		duration := 2 * time.Minute
		fmt.Printf("Waiting for %v...\n", duration)
		time.Sleep(duration)
		closeChrome()
		fmt.Println("Done!")
	}
}

// Helper function to open a URL in the default web browser
func openURL(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("google-chrome", "--incognito", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}

func closeChrome() {
	cmd := exec.Command("pkill", "chrome")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
