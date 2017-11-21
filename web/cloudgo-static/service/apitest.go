package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func apiTestHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			ID      string `json:"id"`
			Content string `json:"content"`
			Name    string `json:"name"`
			Date    string `json:"date"`
		}{ID: "15331145", Content: "lg", Name: "lg", Date: "2017/11/16"})
	}
}
