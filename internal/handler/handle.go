package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"name_api/internal/controller"
	"name_api/internal/models"
	"net/http"
	"strconv"
)

type handlerInterface interface {
	GetInfo(c *gin.Context)
	DeletePerson(c *gin.Context)
	PutPerson(c *gin.Context)
	PostPerson(c *gin.Context)
}

type Handler struct {
	controller controller.ControllerInterface
}

func NewHandler(controller controller.ControllerInterface) *Handler {
	return &Handler{
		controller: controller,
	}
}

func (h Handler) GetInfo(c *gin.Context) {
	log.Info("Start get info")

	ageQ, ok := c.GetQuery("age")
	log.Debugf("handler.GetInfo; age: %s", ageQ)
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "wrong age format"})
		return
	}

	genderQ, ok := c.GetQuery("gender")
	log.Debugf("handler.GetInfo; gender: %s", genderQ)
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "wrong gender format"})
		return
	}

	page := c.DefaultQuery("page", "1")
	log.Debugf("handler.GetInfo; page: %s", page)

	params, err := models.NewRequestParam(ageQ, genderQ, page)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	people, err := h.controller.GetInfo(params)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	jsonPeople, err := json.Marshal(people)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, jsonPeople)
	log.Info("Get success")
}

func (h Handler) DeletePerson(c *gin.Context) {
	log.Infof("Start delete person")

	idStr := c.Param("id")
	log.Debugf("handler.DeletePerson; id: %s", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.controller.DeleteEntity(id); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "person deleted"})
	log.Info("Delete success")
}

func (h Handler) PutPerson(c *gin.Context) {
	log.Info("Start put person")

	var person models.FIO
	idStr := c.Param("id")
	log.Debugf("handler.PutPerson; id: %s", idStr)

	if err := c.ShouldBindJSON(&person); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.controller.UpdateEntity(id, person); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "update person"})
	log.Info("Put success")
}

func (h Handler) PostPerson(c *gin.Context) {
	log.Info("Start post new person")

	var person models.FIO
	if err := c.BindJSON(&person); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	log.Debugf("handler.PostPerson; Person: %v", person)
	if err := h.controller.CreateEntity(person); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "person created"})
	log.Info("Post success")
}
