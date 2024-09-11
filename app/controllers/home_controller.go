package controllers

import (
	"net/http"

	"github.com/unrolled/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	rndr := render.New(render.Options{Layout: "layout"})

	rndr.HTML(w, http.StatusOK, "home", map[string]interface{}{
		"title":       "Home",
		"description": "Home Description",
	})
}
