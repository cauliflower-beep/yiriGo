package main

import (
	"crypto/tls"
	"encoding/csv"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Create a file to write to
	file, err := os.Create("douban_top250.csv")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("\xEF\xBB\xBF")) // Write the UTF-8 BOM to the file
	defer file.Close()

	// Write CSV headers
	writer := csv.NewWriter(file)
	writer.Write([]string{"Rank", "Title", "Year", "Director", "Stars", "Rating", "Number of Ratings"})

	// Scrape the top 250 movies from Douban
	for i := 0; i < 250; i += 25 {
		client := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}

		req, err := http.NewRequest("GET", "https://movie.douban.com/top250", nil)
		if err != nil {
			panic(err)
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			panic(err)
		}

		// Extract movie information from the HTML
		doc.Find(".item").Each(func(j int, s *goquery.Selection) {
			rank := i + j + 1
			title := s.Find(".title").Eq(0).Text()
			year := strings.Trim(s.Find(".bd p").Eq(0).Text(), "()")
			director := strings.Split(s.Find(".bd p").Eq(0).Text(), " ")[1]
			stars := s.Find(".bd p").Eq(1).Text()
			rating := s.Find(".rating_num").Eq(0).Text()
			numRatings := strings.Trim(s.Find(".star span").Eq(3).Text(), "()")

			// Write movie information to CSV
			writer.Write([]string{strconv.Itoa(rank), title, year, director, stars, rating, numRatings})
		})
	}

	writer.Flush()
	fmt.Println("Done!")
}
