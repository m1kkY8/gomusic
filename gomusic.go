package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/m1kkY8/gomusic/utils"
)

type Video struct {
    Url string `json:"url"`
    Title string `json:"title"`
}

func main() {
    
    var searchQuery string
    var urls[] string 
    url := "http://localhost:3001/search?q="
    
    if !utils.CheckDependencies(){
        fmt.Println("yt-dlp is not installed")
        return
    }

    if !utils.IsServerRunning() {
        fmt.Println("Server is not running")
        return
    }

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
    
    utils.Downloader(urls)
}
