package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/italobbarros/go-mqtt-broker/docs"
)

// Endpoints for Session

// createSession cria uma nova sessão.
// @Summary Create a new session
// @Description Create a new session
// @Tags Sessions
// @Accept json
// @Produce json
// @Param input body SessionRequest true "Session object that needs to be added"
// @Success 201 {object} SessionRequest
// @Router /sessions [post]
func (a *API) createSession(c *gin.Context) {
	var session Session
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.db.Create(&session)
	c.JSON(http.StatusCreated, session)
}

// getAllSessions obtém todas as sessões.
// @Summary Get all sessions
// @Description Get all sessions
// @Tags Sessions
// @Produce json
// @Success 200 {array}  Session
// @Router /sessions [get]
func (a *API) getAllSessions(c *gin.Context) {
	var sessions []Session
	a.db.Find(&sessions)
	c.JSON(http.StatusOK, sessions)
}

// getSessionByID obtém uma sessão pelo ID.
// @Summary Get a session by ID
// @Description Get a session by ID
// @Tags Sessions
// @Produce json
// @Param id path int true "Session ID"
// @Success 200 {object}  Session
// @Router /sessions/{id} [get]
func (a *API) getSessionByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var session Session
	if err := a.db.First(&session, id).Error; err != nil {
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
// @Param input body SessionRequest true "Updated session object"
// @Success 200 {object} SessionRequest
// @Router /sessions/{id} [put]
func (a *API) updateSession(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var session Session
	if err := a.db.First(&session, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.db.Save(&session)
	c.JSON(http.StatusOK, session)
}
