package data

import "time"

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
