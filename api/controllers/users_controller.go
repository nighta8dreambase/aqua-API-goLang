package controllers

import (
	"net/http"

	"github.com/TanutN/Aqua/api/models"
	"github.com/TanutN/Aqua/api/responses"
)

func (server *Server) GetAllSos(w http.ResponseWriter, r *http.Request) {

	user := models.Tracker_sos{}

	users, err := user.FindAllSOS(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func (server *Server) GetDevices(w http.ResponseWriter, r *http.Request) {

	user := models.Wearing_devices{}

	users, err := user.FindAllDevices(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}
