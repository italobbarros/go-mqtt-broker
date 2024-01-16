package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/italobbarros/go-mqtt-broker/docs"
)

// Endpoints for Subscription

// createSubscription cria uma nova assinatura.
// @Summary Create a new subscription
// @Description Create a new subscription
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param input body Subscription true "Subscription object that needs to be added"
// @Success 201 {object} Subscription
// @Router /subscriptions [post]
func (a *API) createSubscription(c *gin.Context) {
	var subscription Subscription
	if err := c.ShouldBindJSON(&subscription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.db.Create(&subscription)
	c.JSON(http.StatusCreated, subscription)
}

// getAllSubscriptions obtém todas as assinaturas.
// @Summary Get all subscriptions
// @Description Get all subscriptions
// @Tags Subscriptions
// @Produce json
// @Success 200 {array} Subscription
// @Router /subscriptions [get]
func (a *API) getAllSubscriptions(c *gin.Context) {
	var subscriptions []Subscription
	a.db.Find(&subscriptions)
	c.JSON(http.StatusOK, subscriptions)
}

// getSubscriptionByID obtém uma assinatura pelo ID.
// @Summary Get a subscription by ID
// @Description Get a subscription by ID
// @Tags Subscriptions
// @Produce json
// @Param id path int true "Subscription ID"
// @Success 200 {object} Subscription
// @Router /subscriptions/{id} [get]
func (a *API) getSubscriptionByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var subscription Subscription
	if err := a.db.First(&subscription, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, subscription)
}

// updateSubscription atualiza uma assinatura pelo ID.
// @Summary Update a subscription by ID
// @Description Update a subscription by ID
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param id path int true "Subscription ID"
// @Param input body SubscriptionRequest true "Updated subscription object"
// @Success 200 {object} SubscriptionRequest
// @Router /subscriptions/{id} [put]
func (a *API) updateSubscription(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var subscription Subscription
	if err := a.db.First(&subscription, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&subscription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.db.Save(&subscription)
	c.JSON(http.StatusOK, subscription)
}
