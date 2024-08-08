package main 

import (
    "bufio"
    "io"
    "fmt"
    "log"
    "os"
    "os/exec"
    "net/http"
	"strings"
    "encoding/json"
)

type Video struct {
    Url string `json:"url"`
    Title string `json:"title"`
}

func download(urls []string) {
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

func main() {
    
    var searchQuery string
    var urls[] string 
    url := "http://localhost:3001/search?q="

    for {
        fmt.Printf("Enter song: ")
        reader := bufio.NewReader(os.Stdin)
        line, err := reader.ReadString('\n')
        if err != nil {
            log.Print(err)
        }

        searchQuery = strings.TrimSpace(line)
        if searchQuery == "Exit" {
            break
        }

        searchQuery = strings.ReplaceAll(searchQuery, " ", "%20")
        urlWithQuery := fmt.Sprintf("%s%s", url, searchQuery)
        
        resp, err := http.Get(urlWithQuery) 
        if err != nil{
            log.Print(err)
        }
        defer resp.Body.Close()

        body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Print("Error reading response body:", err)
			continue
		}

		var video Video
        if err := json.Unmarshal(body, &video); err != nil {
			log.Print("Error decoding JSON response:", err)
			continue
		}

        urls = append(urls, video.Url)
    }
    
    download(urls)
}
