package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const apiKey = "********"
const baseURL = "http://www.omdbapi.com/?apikey=" + apiKey

type Movie struct {
	Title      string   `json:"Title"`
	Year       string   `json:"Year"`
	Rated      string   `json:"Rated"`
	Released   string   `json:"Released"`
	Runtime    string   `json:"Runtime"`
	Genre      string   `json:"Genre"`
	Director   string   `json:"Director"`
	Writer     string   `json:"Writer"`
	Actors     string   `json:"Actors"`
	Plot       string   `json:"Plot"`
	Language   string   `json:"Language"`
	Country    string   `json:"Country"`
	Awards     string   `json:"Awards"`
	Poster     string   `json:"Poster"`
	Ratings    []Rating `json:"Ratings"`
	Metascore  string   `json:"Metascore"`
	ImdbRating string   `json:"imdbRating"`
	ImdbVotes  string   `json:"imdbVotes"`
	ImdbID     string   `json:"imdbID"`
	Type       string   `json:"Type"`
	DVD        string   `json:"DVD"`
	BoxOffice  string   `json:"BoxOffice"`
	Production string   `json:"Production"`
	Website    string   `json:"Website"`
	Response   string   `json:"Response"`
}

type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

func fetchMovie(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("%s&t=%s", baseURL, title)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func fetchMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	genre := r.URL.Path[len("/api/movies/"):]
	if genre == "" {
		http.Error(w, "Genre is required", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://api.sampleapis.com/movies/%s", genre)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var movies []Movie
	if err := json.NewDecoder(resp.Body).Decode(&movies); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	http.HandleFunc("/api/movie", fetchMovie)
	http.HandleFunc("/api/movies/", fetchMoviesByGenre)

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
