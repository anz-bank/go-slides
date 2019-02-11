package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type TeapotHandler interface {
	handlePost(w http.ResponseWriter, r *http.Request)
	handleGet(w http.ResponseWriter, r *http.Request)
}

func SetupRoutes(r chi.Router, h TeapotHandler) {
	r.Post("/api/teapot", h.handlePost)
	r.Get("/api/teapot/{id}", h.handleGet)
}

func main() {
	r := chi.NewRouter()
	h := NewMapTeapotHandler()
	SetupRoutes(r, h)
	http.ListenAndServe(":7735", r)
}

type Teapot struct {
	ID          int
	Colour      string
	Temperature int
}

// MapTeapotHandler implements TeapotHandler
type MapTeapotHandler struct {
	store map[int]Teapot // Race condition: don't use in production
	maxID int
}

func NewMapTeapotHandler() *MapTeapotHandler {
	return &MapTeapotHandler{store: map[int]Teapot{}}
}

func (mt *MapTeapotHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	teapot, _ := mt.store[id]
	render.JSON(w, r, teapot)
}

func (mt *MapTeapotHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	var teapot Teapot
	render.DecodeJSON(r.Body, &teapot)
	mt.maxID++ // Race condition: don't use in production
	teapot.ID = mt.maxID
	mt.store[mt.maxID] = teapot
	render.JSON(w, r, teapot)
}
