package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/italobbarros/go-mqtt-broker/docs"
	models "github.com/italobbarros/go-mqtt-broker/internal/api/models"
)

// Endpoints for Session

// createSession cria uma nova sessão.
// @Summary Create a new session
// @Description Create a new session
// @Tags Sessions
// @Accept json
// @Produce json
// @Param input body models.SessionRequest true "Session object that needs to be added"
// @Success 201 {object} models.SessionRequest
// @Router /sessions [post]
func (r *Routes) CreateSession(c *gin.Context) {
	var session models.Session
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r.db.Create(&session)
	c.JSON(http.StatusCreated, session)
}

// getAllSessions obtém todas as sessões.
// @Summary Get all sessions
// @Description Get all sessions
// @Tags Sessions
// @Produce json
// @Success 200 {array}  models.Session
// @Router /sessions [get]
func (r *Routes) GetAllSessions(c *gin.Context) {
	var sessions []models.Session
	r.db.Find(&sessions)
	c.JSON(http.StatusOK, sessions)
}

// getSessionByID obtém uma sessão pelo ID.
// @Summary Get a session by ID
// @Description Get a session by ID
// @Tags Sessions
// @Produce json
// @Param id path int true "Session ID"
// @Success 200 {object}  models.Session
// @Router /sessions/{id} [get]
func (r *Routes) GetSessionByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var session models.Session
	if err := r.db.First(&session, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, session)
}

// updateSession atualiza uma sessão pelo ID.
// @Summary Update a session by ID
// @Description Update a session by ID
// @Tags Sessions
// @Accept json
// @Produce json
// @Param id path int true "Session ID"
// @Param input body models.SessionRequest true "Updated session object"
// @Success 200 {object} models.SessionRequest
// @Router /sessions/{id} [put]
func (r *Routes) UpdateSession(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var session models.Session
	if err := r.db.First(&session, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r.db.Save(&session)
	c.JSON(http.StatusOK, session)
}
