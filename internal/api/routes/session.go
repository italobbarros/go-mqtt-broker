package routes

import (
	"net/http"
	"time"

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
	c.JSON(http.StatusCreated, gin.H{"detail": "success.created.session"})
}

// getAllSessions obtém todas as sessões.
// @Summary Get all sessions
// @Description Get all sessions
// @Tags Sessions
// @Produce json
// @Success 200 {array}  models.Session
// @Router /sessions/all [get]
func (r *Routes) GetAllSessions(c *gin.Context) {
	var sessions []models.Session
	r.db.Find(&sessions)
	c.JSON(http.StatusOK, sessions)
}

// getSessionByID obtém uma sessão pelo ClientId.
// @Summary Get a session by ClientId
// @Description Get a session by ClientId
// @Tags Sessions
// @Produce json
// @Param ClientId query string true "Session ClientId"
// @Success 200 {object}  models.SessionResponse
// @Router /sessions [get]
func (r *Routes) GetSessionByClientId(c *gin.Context) {
	clientId := c.Query("ClientId")
	r.logger.Debug("clientId: %s", clientId)
	var sessionResponse models.SessionResponse
	if err := r.db.Model(&models.Session{}).
		Select("sessions.*, containers.*").
		Joins("join containers on sessions.\"IdContainer\"=containers.\"Id\"").
		Where("sessions.\"ClientId\" = ?", clientId).
		First(&sessionResponse).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Error getting session by ClientId"})
		r.logger.Error("Error: %s", err.Error())
		return
	}
	c.JSON(http.StatusOK, sessionResponse)
}

// updateSession atualiza uma sessão pelo ClientId.
// @Summary Update a session by ClientId
// @Description Update a session by ClientId
// @Tags Sessions
// @Accept json
// @Produce json
// @Param ClientId query string true "Session ClientId"
// @Param input body models.SessionUpdateRequest true "Updated session object"
// @Success 200 {object} models.GenericResponse
// @Router /sessions [put]
func (r *Routes) UpdateSession(c *gin.Context) {
	clientId := c.Query("ClientId")
	var session models.Session
	if err := r.db.Where("\"ClientId\" = ?", clientId).First(&session).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"detail": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}
	session.Updated = time.Now()
	r.db.Save(&session)
	c.JSON(http.StatusOK, gin.H{"detail": "session.updated"})
}
