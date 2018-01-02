package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const apiEndpoint = "https://yts.ag/api/v2"

// Sort options
const (
	SortByTitle     = "title"
	SortByYear      = "year"
	SortByRating    = "rating"
	SortByPeers     = "peers"
	SortBySeeds     = "seeds"
	SortByDownload  = "download_count"
	SortByLike      = "like_count"
	SortByDateAdded = "date_added"
)

// Order options
const (
	OrderAsc  = "asc"
	OrderDesc = "desc"
)

// Movie information struct
type Movie struct {
	DateUploaded     string    `json:"date_uploaded"`
	DateUploadedUnix int64     `json:"date_uploaded_unix"`
	Genres           []string  `json:"genres"`
	ID               int       `json:"id"`
	ImdbID           string    `json:"imdb_code"`
	Language         string    `json:"language"`
	MediumCover      string    `json:"medium_cover_image"`
	Rating           float64   `json:"rating"`
	Runtime          int       `json:"runtime"`
	SmallCover       string    `json:"small_cover_image"`
	State            string    `json:"state"`
	Title            string    `json:"title"`
	TitleLong        string    `json:"title_long"`
	Torrents         []Torrent `json:"torrents"`
	Year             int       `json:"year"`
}

// Torrent information struct
type Torrent struct {
	DateUploaded     string `json:"date_uploaded"`
	DateUploadedUnix int64  `json:"date_uploaded_unix"`
	Hash             string `json:"hash"`
	Peers            int    `json:"peers"`
	Quality          string `json:"quality"`
	Seeds            int    `json:"seeds"`
	Size             string `json:"size"`
	SizeBytes        int    `json:"size_bytes"`
	URL              string `json:"url"`
}

// Data is movie informations
type Data struct {
	PageNumber int     `json:"int"`
	Movies     []Movie `json:"movies"`
}

// Result is http request status
type Result struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Data          Data   `json:"data"`
}

func getMovieList(URL string) ([]Movie, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result *Result
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.Movies, nil
}

// GetNewMovies gets a list of movies
func GetNewMovies(limit string) ([]Movie, error) {
	v := url.Values{}
	v.Set("limit", limit)
	v.Set("sort_by", SortByDateAdded)
	v.Set("order_by", OrderDesc)
	URL := fmt.Sprintf("%s/list_movies.json?%s", apiEndpoint, v.Encode())
	return getMovieList(URL)
}
