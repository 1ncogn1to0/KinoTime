package movies

import (
	"PracticeServer/db"
	"PracticeServer/models"
	"encoding/json"
	"fmt"
	"net/http"
)

const TMDB_API_KEY = "2d5c7316"

func FetchMovies() error {
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/popular?api_key=%s&language=en-US&page=1", TMDB_API_KEY)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch movies: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result struct {
		Results []struct {
			Title       string `json:"title"`
			Overview    string `json:"overview"`
			ReleaseDate string `json:"release_date"`
			PosterPath  string `json:"poster_path"`
		} `json:"results"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	for _, movie := range result.Results {
		newMovie := models.Movie{
			Title:       movie.Title,
			Description: movie.Overview,
			Price:       float64(5 + len(movie.Title)%10), // Случайная цена
			Genre:       "Popular",
			ReleaseDate: movie.ReleaseDate,
			ImageURL:    fmt.Sprintf("https://image.tmdb.org/t/p/w500%s", movie.PosterPath),
		}
		db.DB.Create(&newMovie)
	}

	return nil
}
