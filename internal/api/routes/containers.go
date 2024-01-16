package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/italobbarros/go-mqtt-broker/docs"
	models "github.com/italobbarros/go-mqtt-broker/internal/api/models"
)

// Example for Container model

// createContainer cria um novo container.
// @Summary Create a new container
// @Description Create a new container
// @Tags Container
// @Accept json
// @Produce json
// @Param input body models.ContainerRequest true "Container object that needs to be added"
// @Success 201 {object} models.ContainerRequest
// @Router /containers [post]
func (r *Routes) CreateContainer(c *gin.Context) {
	var container models.Container
	if err := c.ShouldBindJSON(&container); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r.db.Create(&container)
	c.JSON(http.StatusCreated, container)
}

// getAllContainers obtém todos os containers.
// @Summary Get all containers
// @Description Get all containers
// @Tags Container
// @Produce json
// @Success 200 {array} models.Container
// @Router /containers [get]
func (r *Routes) GetAllContainers(c *gin.Context) {
	var containers []models.Container
	r.db.Find(&containers)
	c.JSON(http.StatusOK, containers)
}

// getContainerByID obtém um container pelo ID.
// @Summary Get a container by ID
// @Description Get a container by ID
// @Tags Container
// @Produce json
// @Param id path int true "Container ID"
// @Success 200 {object} models.Container
// @Router /containers/{id} [get]
func (r *Routes) GetContainerByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var container models.Container
	if err := r.db.First(&container, id).Error; err != nil {
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
// @Param input body models.ContainerRequest true "Updated container object"
// @Success 200 {object} models.ContainerRequest
// @Router /containers/{id} [put]
func (r *Routes) UpdateContainer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var container models.Container
	if err := r.db.First(&container, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&container); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r.db.Save(&container)
	c.JSON(http.StatusOK, container)
}
