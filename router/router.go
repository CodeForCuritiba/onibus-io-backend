package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/codeforcuritiba/onibus-io-backend/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func NewRouter(s store.Storer) chi.Router {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Mount("/", appRoutes())
	r.Mount("/api/linhas", linhasRoutes(s))
	r.Mount("/api/veiculos", veiculosRoutes(s))
	return r
}

func appRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/versao", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET /versao")
		json.NewEncoder(w).Encode(map[string]string{
			"versao":    "0.0.1",
			"descricao": "API  do onibus-io",
		})
	})
	return r
}

func linhasRoutes(s store.Storer) chi.Router {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		linhas, err := s.Linhas()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(linhas)
	})
	r.Get("/{codigo}", func(w http.ResponseWriter, r *http.Request) {
		codigo := chi.URLParam(r, "codigo")
		linha, err := s.Linha(codigo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(linha)
	})
	return r
}

func veiculosRoutes(s store.Storer) chi.Router {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		veiculos, err := s.Veiculos()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(veiculos)
	})
	r.Get("/{codigo}", func(w http.ResponseWriter, r *http.Request) {
		codigo := chi.URLParam(r, "codigo")
		veiculo, err := s.Veiculo(codigo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(veiculo)
	})
	return r
}
