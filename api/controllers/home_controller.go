package controllers

import (
	"net/http"

	"github.com/ebenovya/poster/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "#evdekal")

}
