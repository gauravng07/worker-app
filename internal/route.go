package internal

import (
	"hello-fresh-app/internal/handler"

	"github.com/gorilla/mux"
)

func Configure(r *mux.Router)  {
	r.HandleFunc("/recipe",handler.RecipeTemplateHandler)
	r.HandleFunc("/recipe/stats", handler.RecipeStatsHandler)
}