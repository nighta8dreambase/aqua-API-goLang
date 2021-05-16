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

func (server *Server) GetReasons(c *gin.Context) {
	projectKey := c.Request.Header["Project-Key"][0]
	reason := models.RescuerResons{}
	reasons, err := reason.FindAllReason(server.DB, projectKey)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusInternalServerError, err)
		return
	}
	var response []map[string]interface{}
	resjs, _ := json.Marshal(reasons)
	json.Unmarshal(resjs, &response)
	responses.JSON(c.Writer, http.StatusOK, responses.SuccessList(response))
}

func (server *Server) GetReasonsByID(c *gin.Context) {
	projectKey := c.Request.Header["Project-Key"][0]
	uid := c.Param("id")
	reason := models.RescuerResons{}
	userGotten, err := reason.FindReasonByID(server.DB, uid, projectKey)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusBadRequest, err)
		return
	}
	var response map[string]interface{}
	resjs, _ := json.Marshal(userGotten)
	json.Unmarshal(resjs, &response)
	responses.JSON(c.Writer, http.StatusOK, responses.Success(response))
}

func (server *Server) CreateReasons(c *gin.Context) {
	projectKey := c.Request.Header["Project-Key"][0]
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
	}
	reason := models.RescuerResons{}
	err = json.Unmarshal(body, &reason)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	createdReason, err := reason.CreateReason(server.DB,projectKey)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusInternalServerError, err)
		return
	}
	var response map[string]interface{}
	resjs, _ := json.Marshal(createdReason)
	json.Unmarshal(resjs, &response)
	responses.JSON(c.Writer, http.StatusOK, responses.Success(response))
}

func (server *Server) UpdateReasons(c *gin.Context) {
	projectKey := c.Request.Header["Project-Key"][0]
	uid := c.Param("id")
	reason := models.RescuerResons{}
	userGotten, err := reason.FindReasonByID(server.DB, uid, projectKey)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	reasonInput := models.ReasonInput{}
	err = json.Unmarshal(body, &reasonInput)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUser, err := reasonInput.UpdateAReason(server.DB, uid, userGotten)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusInternalServerError, err)
		return
	}
	var response map[string]interface{}
	resjs, _ := json.Marshal(updatedUser)
	json.Unmarshal(resjs, &response)
	responses.JSON(c.Writer, http.StatusOK, responses.Success(response))
}

func (server *Server) DeleteReasons(c *gin.Context) {

	uid := c.Param("id")
	reason := models.RescuerResons{}
	_, err := reason.DeleteAReason(server.DB, uid)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusInternalServerError, err)
		return
	}
	c.Writer.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(c.Writer, http.StatusOK, responses.Success(nil))
}
