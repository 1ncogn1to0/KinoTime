package movies

import (
	"PracticeServer/db"
	"PracticeServer/models"
	"log"
)

func SeedMovies() {
	movies := []models.Movie{
		{Title: "Inception", Description: "A mind-bending thriller", Price: 9.99, Genre: "Sci-Fi", ReleaseDate: "2010-07-16", ImageURL: "https://via.placeholder.com/150"},
		{Title: "The Dark Knight", Description: "A tale of a hero in Gotham", Price: 8.99, Genre: "Action", ReleaseDate: "2008-07-18", ImageURL: "https://via.placeholder.com/150"},
		{Title: "Interstellar", Description: "A journey beyond the stars", Price: 10.99, Genre: "Sci-Fi", ReleaseDate: "2014-11-07", ImageURL: "https://via.placeholder.com/150"},
		{Title: "The Matrix", Description: "Reality is not what it seems", Price: 7.99, Genre: "Sci-Fi", ReleaseDate: "1999-03-31", ImageURL: "https://via.placeholder.com/150"},
		{Title: "Avatar", Description: "A visually stunning world", Price: 12.99, Genre: "Fantasy", ReleaseDate: "2009-12-18", ImageURL: "https://via.placeholder.com/150"},
		{Title: "Titanic", Description: "A timeless love story", Price: 11.99, Genre: "Drama", ReleaseDate: "1997-12-19", ImageURL: "https://via.placeholder.com/150"},
		{Title: "The Avengers", Description: "Earth's mightiest heroes", Price: 10.99, Genre: "Action", ReleaseDate: "2012-05-04", ImageURL: "https://via.placeholder.com/150"},
		{Title: "Gladiator", Description: "A fight for honor and glory", Price: 8.99, Genre: "Historical", ReleaseDate: "2000-05-05", ImageURL: "https://via.placeholder.com/150"},
		{Title: "Jurassic Park", Description: "Dinosaurs brought back to life", Price: 9.49, Genre: "Adventure", ReleaseDate: "1993-06-11", ImageURL: "https://via.placeholder.com/150"},
		{Title: "Forrest Gump", Description: "Life is like a box of chocolates", Price: 7.49, Genre: "Drama", ReleaseDate: "1994-07-06", ImageURL: "https://via.placeholder.com/150"},
		{Title: "The Shawshank Redemption", Description: "Hope is a good thing, maybe the best of things.", Price: 8.99, Genre: "Drama", ReleaseDate: "1994-09-23", ImageURL: "https://via.placeholder.com/150"},
		{Title: "Inception", Description: "Your mind is the scene of the crime.", Price: 9.99, Genre: "Sci-Fi", ReleaseDate: "2010-07-16", ImageURL: "https://via.placeholder.com/150"},
		{Title: "The Godfather", Description: "An offer you can't refuse.", Price: 12.49, Genre: "Crime", ReleaseDate: "1972-03-24", ImageURL: "https://via.placeholder.com/150"},
		{Title: "The Dark Knight", Description: "Why so serious?", Price: 11.99, Genre: "Action", ReleaseDate: "2008-07-18", ImageURL: "https://via.placeholder.com/150"},
		{Title: "Pulp Fiction", Description: "Just because you are a character doesn’t mean you have character.", Price: 10.49, Genre: "Crime", ReleaseDate: "1994-10-14", ImageURL: "https://via.placeholder.com/150"},
		{Title: "Fight Club", Description: "The first rule of Fight Club is: You do not talk about Fight Club.", Price: 8.49, Genre: "Drama", ReleaseDate: "1999-10-15", ImageURL: "https://via.placeholder.com/150"},
		{Title: "Schindler's List", Description: "Whoever saves one life, saves the world entire.", Price: 10.99, Genre: "Historical", ReleaseDate: "1993-12-15", ImageURL: "https://via.placeholder.com/150"},
		{Title: "The Matrix", Description: "Free your mind.", Price: 8.99, Genre: "Sci-Fi", ReleaseDate: "1999-03-31", ImageURL: "https://via.placeholder.com/150"},
		{Title: "Gladiator", Description: "What we do in life echoes in eternity.", Price: 9.49, Genre: "Action", ReleaseDate: "2000-05-05", ImageURL: "https://via.placeholder.com/150"},
		{Title: "Titanic", Description: "Nothing on Earth could come between them.", Price: 7.99, Genre: "Romance", ReleaseDate: "1997-12-19", ImageURL: "https://via.placeholder.com/150"},
		{Title: "Braveheart", Description: "Every man dies, not every man really lives.", Price: 9.99, Genre: "Historical", ReleaseDate: "1995-05-24", ImageURL: "https://via.placeholder.com/150"},
		{Title: "The Lion King", Description: "Remember who you are.", Price: 6.99, Genre: "Animation", ReleaseDate: "1994-06-15", ImageURL: "https://via.placeholder.com/150"},
		{Title: "Goodfellas", Description: "As far back as I can remember, I always wanted to be a gangster.", Price: 8.49, Genre: "Crime", ReleaseDate: "1990-09-19", ImageURL: "https://via.placeholder.com/150"},
		{Title: "Star Wars: Episode IV - A New Hope", Description: "May the Force be with you.", Price: 10.99, Genre: "Sci-Fi", ReleaseDate: "1977-05-25", ImageURL: "https://via.placeholder.com/150"},
		{Title: "The Avengers", Description: "Avengers, assemble!", Price: 11.99, Genre: "Action", ReleaseDate: "2012-05-04", ImageURL: "https://via.placeholder.com/150"},
		{Title: "Jurassic Park", Description: "Life finds a way.", Price: 7.49, Genre: "Sci-Fi", ReleaseDate: "1993-06-11", ImageURL: "https://via.placeholder.com/150"},
		{Title: "The Silence of the Lambs", Description: "I do wish we could chat longer, but I’m having an old friend for dinner.", Price: 8.99, Genre: "Thriller", ReleaseDate: "1991-02-14", ImageURL: "https://via.placeholder.com/150"},
		{Title: "Saving Private Ryan", Description: "This is why we fight, for a brother at your side.", Price: 9.99, Genre: "War", ReleaseDate: "1998-07-24", ImageURL: "https://via.placeholder.com/150"},
		{Title: "Avatar", Description: "I see you.", Price: 12.49, Genre: "Sci-Fi", ReleaseDate: "2009-12-18", ImageURL: "https://via.placeholder.com/150"},
		{Title: "The Lord of the Rings: The Fellowship of the Ring", Description: "Even the smallest person can change the course of the future.", Price: 10.99, Genre: "Fantasy", ReleaseDate: "2001-12-19", ImageURL: "https://via.placeholder.com/150"},
	}

	for i := len(movies); i < 30; i++ {
		movie := models.Movie{
			Title:       "Movie " + string(i+1),
			Description: "This is a placeholder movie",
			Price:       float64(5 + i%10),
			Genre:       "Drama",
			ReleaseDate: "2025-01-01",
			ImageURL:    "https://via.placeholder.com/150",
		}
		movies = append(movies, movie)
	}

	for _, movie := range movies {
		result := db.DB.Create(&movie)
		if result.Error != nil {
			log.Printf("Error creating movie: %v", result.Error)
		}
	}

	log.Println("30 movies successfully added to the database!")
}
