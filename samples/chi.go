package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Teapot struct {
	ID          int    `json:"id"`
	Colour      string `json:"colour"`
	Temperature int    `json:"temperature"` // Celsius
}

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
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := NewMapTeapotHandler()
	SetupRoutes(r, h)

	addr := ":7735"
	fmt.Printf("Starting server on port %s. Try:\n", addr)
	fmt.Printf("  curl localhost%s/api/teapot -d ", addr)
	fmt.Println(`'{"colour": "blue","temperature": 88}'`)
	fmt.Printf("  curl localhost%s/api/teapot/1 \n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}

type MapTeapotHandler struct {
	// Race condition, don't use `map` in production for concurrent access
	store map[int]Teapot
	maxID int
}

func NewMapTeapotHandler() *MapTeapotHandler {
	return &MapTeapotHandler{store: map[int]Teapot{}}
}

func (mt *MapTeapotHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	teapot, ok := mt.store[id]
	if !ok {
		http.NotFound(w, r)
		return
	}
	render.JSON(w, r, teapot)
}

func (mt *MapTeapotHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	var teapot Teapot
	if err := render.DecodeJSON(r.Body, &teapot); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Race condition, don't do this in production
	mt.maxID++
	teapot.ID = mt.maxID
	mt.store[mt.maxID] = teapot
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, teapot)
}
