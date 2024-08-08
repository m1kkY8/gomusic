package utils

import (
    "os"
    "os/exec"
    "log"
    "time"
    "net/http"
)

func IsServerRunning() bool {
    url := "http://localhost:3001/health"
	client := &http.Client{
		Timeout: 2 * time.Second, 
	}

	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	// Check if the response status code is 200 OK
	return resp.StatusCode == http.StatusOK
}

func CheckDependencies() bool {
    _ , err := exec.LookPath("yt-dlp")
    if err != nil{
        return false
    }

    return true 
}

func Downloader(urls []string) {
    for _ , url := range urls {
        args := []string{
            "--no-write-description",
            "-q",
            "--no-playlist",
            "--extract-audio",
            "--add-metadata",
            "--audio-format", "mp3",
            "--audio-quality", "0",
            url,
        }

        cmd := exec.Command("yt-dlp", args...)

        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        err := cmd.Run()
        if err != nil {
            log.Print("error downloading")
        }
    }
}

