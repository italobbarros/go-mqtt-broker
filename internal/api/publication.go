package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/italobbarros/go-mqtt-broker/docs"
)

// Endpoints for Publication

// createPublication cria uma nova publicação.
// @Summary Create a new publication
// @Description Create a new publication
// @Tags Publications
// @Accept json
// @Produce json
// @Param input body PublicationRequest true "Publication object that needs to be added"
// @Success 201 {object} PublicationRequest
// @Router /publications [post]
func (a *API) createPublication(c *gin.Context) {
	var publication Publication
	if err := c.ShouldBindJSON(&publication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.db.Create(&publication)
	c.JSON(http.StatusCreated, publication)
}

// getAllPublications obtém todas as publicações.
// @Summary Get all publications
// @Description Get all publications
// @Tags Publications
// @Produce json
// @Success 200 {array}  Publication
// @Router /publications [get]
func (a *API) getAllPublications(c *gin.Context) {
	var publications []Publication
	a.db.Find(&publications)
	c.JSON(http.StatusOK, publications)
}

// getPublicationByID obtém uma publicação pelo ID.
// @Summary Get a publication by ID
// @Description Get a publication by ID
// @Tags Publications
// @Produce json
// @Param id path int true "Publication ID"
// @Success 200 {object}  Publication
// @Router /publications/{id} [get]
func (a *API) getPublicationByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var publication Publication
	if err := a.db.First(&publication, id).Error; err != nil {
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
// @Param input body PublicationRequest true "Updated publication object"
// @Success 200 {object} PublicationRequest
// @Router /publications/{id} [put]
func (a *API) updatePublication(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var publication Publication
	if err := a.db.First(&publication, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&publication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.db.Save(&publication)
	c.JSON(http.StatusOK, publication)
}
