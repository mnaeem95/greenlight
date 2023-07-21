package data

import (
	"time"

	"greenlight.alexedwards.net/internal/validator"
)

// Annotate the Movie struct with struct tags to control how the keys appear in the JSON-encoded output.
type Movie struct {
	ID        int64     `json:"id"`             // Unique integer ID for the movie
	CreatedAt time.Time `json:"-"`              // Timestamp for when the movie is added to our database [Use the - directive]
	Title     string    `json:"title"`          // Movie title
	Year      int32     `json:"year,omitempty"` // Movie release year [Add the omitempty directive]
	// if the Runtime field has the underlying value 0, then it will be considered empty and omitted and the MarshalJSON() method won't be called at all.
	Runtime Runtime  `json:"runtime,omitempty"` // Movie runtime (in minutes) [Add the omitempty directive]
	Genres  []string `json:"genres,omitempty"`  // Slice of genres for the movie (romance, comedy, etc.) [Add the omitempty directive]
	Version int32    `json:"version"`           // The version number starts at 1 and will be incremented each time the movie information is updated
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}
