package controllers

import (
	"net/http"

	"github.com/BlackBoyZoovie/fullstack/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "welcome to this wonderful Api")
}
