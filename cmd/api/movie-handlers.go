package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(err.Error())
		app.errorJSON(w, http.StatusBadRequest, errors.New("invalid movie id"))
		return
	}
	movie, err := app.models.DB.Get(id)
	if err != nil {
		app.logger.Print(err.Error())
		app.errorJSON(w, http.StatusInternalServerError, errors.New("could not get movie"))
		return
	}
	err = app.writeJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, http.StatusInternalServerError, err)
		return
	}
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.All()
	if err != nil {
		app.logger.Print(err.Error())
		app.errorJSON(w, http.StatusInternalServerError, errors.New("could not get movies"))
		return
	}
	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, http.StatusInternalServerError, err)
		return
	}
}

func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {

}
func (app *application) insertMovie(w http.ResponseWriter, r *http.Request) {

}
func (app *application) updateMovie(w http.ResponseWriter, r *http.Request) {

}
func (app *application) searchMovie(w http.ResponseWriter, r *http.Request) {

}
