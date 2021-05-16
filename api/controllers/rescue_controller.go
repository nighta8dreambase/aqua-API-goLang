package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/TanutN/Aqua/api/models"
	"github.com/TanutN/Aqua/api/responses"
	"github.com/gin-gonic/gin"
)

func (server *Server) CreateSos(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.Input_Wearing{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	SosDevice, err := user.FindSosDevicesByImei(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	SosCreated, err := SosDevice.Create(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, SosCreated)
}

func (server *Server) GetRescuser(c *gin.Context) {

	rescuser := models.Rescuser{}

	rescusers, err := rescuser.FindAllRescusers(server.DB)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusInternalServerError, err)
		return
	}
	var response []map[string]interface{}
	resjs, _ := json.Marshal(rescusers)
	json.Unmarshal(resjs, &response)
	fmt.Println(rescusers)
	responses.JSON(c.Writer, http.StatusOK, responses.SuccessList(response))
}

func (server *Server) GetRescuserByID(c *gin.Context) {

	uid := c.Param("id")
	rescuser := models.Rescuser{}
	userGotten, err := rescuser.FindRescueByID(server.DB, uid)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusBadRequest, err)
		return
	}
	var response map[string]interface{}
	resjs, _ := json.Marshal(userGotten)
	json.Unmarshal(resjs, &response)
	responses.JSON(c.Writer, http.StatusOK, responses.Success(response))
}

func (server *Server) CreateRescuser(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
	}
	rescue := models.Rescuser{}
	err = json.Unmarshal(body, &rescue)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	rescuser, err := rescue.CreateRescuser(server.DB)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusInternalServerError, err)
		return
	}
	var response map[string]interface{}
	resjs, _ := json.Marshal(rescuser)
	json.Unmarshal(resjs, &response)
	responses.JSON(c.Writer, http.StatusOK, responses.Success(response))
}

func (server *Server) UpdateRescuser(c *gin.Context) {

	uid := c.Param("id")
	rescuser := models.Rescuser{}
	userGotten, err := rescuser.FindRescueByID(server.DB, uid)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	rescueInput := models.RescuserUpdate{}
	err = json.Unmarshal(body, &rescueInput)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUser, err := rescueInput.UpdateARescuser(server.DB, uid, userGotten)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusInternalServerError, err)
		return
	}
	var response map[string]interface{}
	resjs, _ := json.Marshal(updatedUser)
	json.Unmarshal(resjs, &response)
	responses.JSON(c.Writer, http.StatusOK, responses.Success(response))
}

func (server *Server) DeleteRescue(c *gin.Context) {

	uid := c.Param("id")
	rescuser := models.Rescuser{}
	_, err := rescuser.DeleteARescuser(server.DB, uid)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusInternalServerError, err)
		return
	}
	c.Writer.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(c.Writer, http.StatusOK, responses.Success(nil))
}

func (server *Server) ChooseRescue(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
	}
	choose := models.Sos_Rescuser{}
	err = json.Unmarshal(body, &choose)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	/*doo, err := commit.FindAll(server.DB)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusInternalServerError, err)
		return
	}*/
	responses.JSON(c.Writer, http.StatusOK, choose.UUID)
}

func (server *Server) RescuerCommit(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
	}
	commit := models.Rescuser_Input{}
	err = json.Unmarshal(body, &commit)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	doo, err := commit.FindAll(server.DB)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(c.Writer, http.StatusOK, doo)
}
