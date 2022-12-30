package main

import (
	"fmt"
	"net/http"

	"github.com/shynggys9219/greenlight/internal/data"
)

func (app *application) createTrailerHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID           int64  `json:"id"`
		Trailer_name string `json:"trailer_name"`
		Duration     int64  `json:"duration"`
		Premier_date string `json:"premier_date"`
	}

	// if there is error with decoding, we are sending corresponding message
	err := app.readJSON(w, r, &input) //non-nil pointer as the target decode destination
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	}

	trailer := &data.Trailer{
		Trailer_name: input.Trailer_name,
		Duration:     input.Duration,
		Premier_date: input.Premier_date,
	}

	err = app.models.Trailers.Insert(trailer)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/trailers/%d", trailer.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"trailer": trailer}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	// // Dump the contents of the input struct in a HTTP response.
	// fmt.Fprintf(w, "%+v\n", input) //+v here is adding the field name of a value // https://pkg.go.dev/fmt
}
