package data

import (
	"database/sql"
)

// By default, the keys in the JSON object are equal to the field names in the struct ( ID,
// CreatedAt, Title and so on).
type Trailer struct {
	ID           int64  `json:"id"`
	Trailer_name string `json:"trailer_name"`
	Duration     int64  `json:"duration"`
	Premier_date string `json:"premier_date"`
}

// Define a MovieModel struct type which wraps a sql.DB connection pool.
type TrailerModel struct {
	DB *sql.DB
}

// method for inserting a new record in the movies table.
func (m TrailerModel) Insert(trailer *Trailer) error {
	query := `
		INSERT INTO trailers(trailer_name, duration, premier_date)
		VALUES ($1, $2, $3)
		RETURNING id`

	return m.DB.QueryRow(query, &trailer.Trailer_name, &trailer.Duration, &trailer.Premier_date).Scan(&trailer.ID)
}
