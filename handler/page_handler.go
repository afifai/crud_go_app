package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rysmaadit/go-template/common/responder"
	"github.com/rysmaadit/go-template/external/gorm_client"
)

func CreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var movie gorm_client.Movie
		payloads, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, err, nil)
			return
		}
		json.Unmarshal(payloads, &movie)

		err = gorm_client.CreateMovie(movie)

		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, err, nil)
			return
		}
		responder.NewHttpResponse(r, w, http.StatusCreated, movie, nil)
	}
}

func ReadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		movie, err := gorm_client.ReadMovie(args["slug"])
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		responder.NewHttpResponse(r, w, http.StatusOK, movie, nil)
	}
}

func ReadAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		movie, err := gorm_client.ReadAll()
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		responder.NewHttpResponse(r, w, http.StatusOK, movie, nil)
	}
}

func UpdateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var movie gorm_client.Movie
		args := mux.Vars(r)
		payloads, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, err, nil)
			return
		}
		json.Unmarshal(payloads, &movie)

		err = gorm_client.UpdateMovie(args["slug"], movie)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		responder.NewHttpResponse(r, w, http.StatusOK, movie, nil)
	}
}

func DeleteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		err := gorm_client.DeleteMovie(args["slug"])
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, err, nil)
			return
		}
		responder.NewHttpResponse(r, w, http.StatusOK, "success", nil)
	}
}
