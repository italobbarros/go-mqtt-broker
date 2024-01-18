package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/italobbarros/go-mqtt-broker/docs"
	models "github.com/italobbarros/go-mqtt-broker/internal/api/models"
)

// Endpoints for Subscription

// createSubscription cria uma nova assinaturr.
// @Summary Create a new subscription
// @Description Create a new subscription
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param input body models.Subscription true "Subscription object that needs to be added"
// @Success 201 {object} models.Subscription
// @Router /subscriptions [post]
func (r *Routes) CreateSubscription(c *gin.Context) {
	var subscription models.Subscription
	if err := c.ShouldBindJSON(&subscription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r.db.Create(&subscription)
	c.JSON(http.StatusCreated, subscription)
}

// getAllSubscriptions obtém todas as assinaturas.
// @Summary Get all subscriptions
// @Description Get all subscriptions
// @Tags Subscriptions
// @Produce json
// @Success 200 {array} models.Subscription
// @Router /subscriptions [get]
func (r *Routes) GetAllSubscriptions(c *gin.Context) {
	var subscriptions []models.Subscription
	r.db.Find(&subscriptions)
	c.JSON(http.StatusOK, subscriptions)
}

// getSubscriptionByID obtém uma assinatura pelo ID.
// @Summary Get a subscription by ID
// @Description Get a subscription by ID
// @Tags Subscriptions
// @Produce json
// @Param id path int true "Subscription ID"
// @Success 200 {object} models.Subscription
// @Router /subscriptions/{id} [get]
func (r *Routes) GetSubscriptionByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var subscription models.Subscription
	if err := r.db.First(&subscription, id).Error; err != nil {
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
// @Param input body models.SubscriptionRequest true "Updated subscription object"
// @Success 200 {object} models.SubscriptionRequest
// @Router /subscriptions/{id} [put]
func (r *Routes) UpdateSubscription(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var subscription models.Subscription
	if err := r.db.First(&subscription, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&subscription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r.db.Save(&subscription)
	c.JSON(http.StatusOK, subscription)
}
