package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/italobbarros/go-mqtt-broker/docs"
)

// Example for ContainerPost model

// createContainer cria um novo container.
// @Summary Create a new container
// @Description Create a new container
// @Tags Container
// @Accept json
// @Produce json
// @Param input body ContainerPost true "Container object that needs to be added"
// @Success 201 {object} ContainerPost
// @Router /containers [post]
func (a *API) createContainer(c *gin.Context) {
	var container Container
	if err := c.ShouldBindJSON(&container); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.db.Create(&container)
	c.JSON(http.StatusCreated, container)
}

// getAllContainers obtém todos os containers.
// @Summary Get all containers
// @Description Get all containers
// @Tags Container
// @Produce json
// @Success 200 {array} Container
// @Router /containers [get]
func (a *API) getAllContainers(c *gin.Context) {
	var containers []Container
	a.db.Find(&containers)
	c.JSON(http.StatusOK, containers)
}

// getContainerByID obtém um container pelo ID.
// @Summary Get a container by ID
// @Description Get a container by ID
// @Tags Container
// @Produce json
// @Param id path int true "Container ID"
// @Success 200 {object} Container
// @Router /containers/{id} [get]
func (a *API) getContainerByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var container Container
	if err := a.db.First(&container, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, container)
}

// updateContainer atualiza um container pelo ID.
// @Summary Update a container by ID
// @Description Update a container by ID
// @Tags Container
// @Accept json
// @Produce json
// @Param id path int true "Container ID"
// @Param input body ContainerPost true "Updated container object"
// @Success 200 {object} ContainerPost
// @Router /containers/{id} [put]
func (a *API) updateContainer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var container Container
	if err := a.db.First(&container, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&container); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.db.Save(&container)
	c.JSON(http.StatusOK, container)
}
