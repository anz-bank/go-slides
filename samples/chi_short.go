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

func SetupRoutes(h TeapotHandler, r chi.Router) {
	r.Post("/api/teapot", h.handlePost)
	r.Get("/api/teapot/{id}", h.handleGet)
}

func main() {
	r := chi.NewRouter()
	th := NewMapTeapotHandler()
	SetupRoutes(th, r)
	http.ListenAndServe(":7735", r)
}

type Teapot struct {
	ID          int
	Colour      string
	Temperature int
}

// MapTeapotHandler implements TeapotHandler
type MapTeapotHandler struct {
	store map[int]Teapot
	maxID int
}

func NewMapTeapotHandler() *MapTeapotHandler {
	return &MapTeapotHandler{store: make(map[int]Teapot), maxID: 0}
}

func (mt *MapTeapotHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	teapot, _ := mt.store[id]
	render.JSON(w, r, teapot)
}

func (mt *MapTeapotHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	var teapot Teapot
	render.DecodeJSON(r.Body, &teapot)
	mt.maxID += 1
	teapot.ID = mt.maxID
	mt.store[mt.maxID] = teapot
	render.JSON(w, r, teapot)
}
