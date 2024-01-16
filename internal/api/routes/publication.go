package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/italobbarros/go-mqtt-broker/docs"
	models "github.com/italobbarros/go-mqtt-broker/internal/api/models"
)

// Endpoints for Publication

// createPublication cria uma nova publicação.
// @Summary Create a new publication
// @Description Create a new publication
// @Tags Publications
// @Accept json
// @Produce json
// @Param input body models.PublicationRequest true "Publication object that needs to be added"
// @Success 201 {object} models.PublicationRequest
// @Router /publications [post]
func (r *Routes) CreatePublication(c *gin.Context) {
	var publication models.Publication
	if err := c.ShouldBindJSON(&publication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r.db.Create(&publication)
	c.JSON(http.StatusCreated, publication)
}

// getAllPublications obtém todas as publicações.
// @Summary Get all publications
// @Description Get all publications
// @Tags Publications
// @Produce json
// @Success 200 {array}  models.Publication
// @Router /publications [get]
func (r *Routes) GetAllPublications(c *gin.Context) {
	var publications []models.Publication
	r.db.Find(&publications)
	c.JSON(http.StatusOK, publications)
}

// getPublicationByID obtém uma publicação pelo ID.
// @Summary Get a publication by ID
// @Description Get a publication by ID
// @Tags Publications
// @Produce json
// @Param id path int true "Publication ID"
// @Success 200 {object}  models.Publication
// @Router /publications/{id} [get]
func (r *Routes) GetPublicationByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var publication models.Publication
	if err := r.db.First(&publication, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, publication)
}

// updatePublication atualiza uma publicação pelo ID.
// @Summary Update a publication by ID
// @Description Update a publication by ID
// @Tags Publications
// @Accept json
// @Produce json
// @Param id path int true "Publication ID"
// @Param input body models.PublicationRequest true "Updated publication object"
// @Success 200 {object} models.PublicationRequest
// @Router /publications/{id} [put]
func (r *Routes) UpdatePublication(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var publication models.Publication
	if err := r.db.First(&publication, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&publication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r.db.Save(&publication)
	c.JSON(http.StatusOK, publication)
}
