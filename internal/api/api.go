package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/italobbarros/go-mqtt-broker/docs"
	"github.com/italobbarros/go-mqtt-broker/pkg/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// API representa a interface da API do servidor
type API struct {
	db     *gorm.DB
	logger *logger.Logger
}

// Init inicializa a API com as rotas e inicia o servidor
func NewAPI() *API {
	return &API{
		db:     &gorm.DB{},
		logger: logger.NewLogger("API"),
	}
}

func (a *API) initDB() error {
	db, err := gorm.Open(postgres.Open("postgresql://teste:teste@localhost:5433/broker?sslmode=disable"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		return fmt.Errorf("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Container{}, &Topic{}, &Subscription{}, &Publication{}, &Session{})
	if err != nil {
		return err
	}
	a.db = db
	return nil
}

func (a *API) Init() {
	err := a.initDB()
	if err != nil {
		a.logger.Error("Erro init API:%s", err.Error())
		return
	}
	r := gin.Default()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Endpoints for ContainerPost
	r.POST("/containers", a.createContainer)
	r.GET("/containers", a.getAllContainers)
	r.GET("/containers/:id", a.getContainerByID)
	r.PUT("/containers/:id", a.updateContainer)

	// Endpoints for TopicPost
	r.POST("/topics", a.createTopic)
	r.GET("/topics", a.getAllTopics)
	r.GET("/topics/:id", a.getTopicByID)
	r.PUT("/topics/:id", a.updateTopic)

	// Endpoints for Subscription
	r.POST("/subscriptions", a.createSubscription)
	r.GET("/subscriptions", a.getAllSubscriptions)
	r.GET("/subscriptions/:id", a.getSubscriptionByID)
	r.PUT("/subscriptions/:id", a.updateSubscription)

	// Endpoints for Publication
	r.POST("/publications", a.createPublication)
	r.GET("/publications", a.getAllPublications)
	r.GET("/publications/:id", a.getPublicationByID)
	r.PUT("/publications/:id", a.updatePublication)

	// Endpoints for Session
	r.POST("/sessions", a.createSession)
	r.GET("/sessions", a.getAllSessions)
	r.GET("/sessions/:id", a.getSessionByID)
	r.PUT("/sessions/:id", a.updateSession)

	r.Run(":8080")
}

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
// @Success 200 {array} ContainerPostTable
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

// createTopic cria um novo tópico.
// @Summary Create a new topic
// @Description Create a new topic
// @Tags Topic
// @Accept json
// @Produce json
// @Param input body TopicPost true "Topic object that needs to be added"
// @Success 201 {object} TopicPost
// @Router /topics [post]
func (a *API) createTopic(c *gin.Context) {
	var topic Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.db.Create(&topic)
	c.JSON(http.StatusCreated, topic)
}

// getAllTopics obtém todos os tópicos.
// @Summary Get all topics
// @Description Get all topics
// @Tags Topic
// @Produce json
// @Success 200 {array} Topic
// @Router /topics [get]
func (a *API) getAllTopics(c *gin.Context) {
	var topics []Topic
	a.db.Find(&topics)
	c.JSON(http.StatusOK, topics)
}

// getTopicByID obtém um tópico pelo ID.
// @Summary Get a topic by ID
// @Description Get a topic by ID
// @Tags Topic
// @Produce json
// @Param id path int true "Topic ID"
// @Success 200 {object} Topic
// @Router /topics/{id} [get]
func (a *API) getTopicByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var topic Topic
	if err := a.db.First(&topic, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, topic)
}

// updateTopic atualiza um tópico pelo ID.
// @Summary Update a topic by ID
// @Description Update a topic by ID
// @Tags Topic
// @Accept json
// @Produce json
// @Param id path int true "Topic ID"
// @Param input body TopicPost true "Updated topic object"
// @Success 200 {object} TopicPost
// @Router /topics/{id} [put]
func (a *API) updateTopic(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var topic Topic
	if err := a.db.First(&topic, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.db.Save(&topic)
	c.JSON(http.StatusOK, topic)
}

// Endpoints for Subscription

// createSubscription cria uma nova assinatura.
// @Summary Create a new subscription
// @Description Create a new subscription
// @Tags Subscription
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
// @Param input body SubscriptionPost true "Updated subscription object"
// @Success 200 {object} SubscriptionPost
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

// Endpoints for Publication

// createPublication cria uma nova publicação.
// @Summary Create a new publication
// @Description Create a new publication
// @Tags Publications
// @Accept json
// @Produce json
// @Param input body PublicationPost true "Publication object that needs to be added"
// @Success 201 {object} PublicationPost
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
// @Param input body PublicationPost true "Updated publication object"
// @Success 200 {object} PublicationPost
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

// Endpoints for Session

// createSession cria uma nova sessão.
// @Summary Create a new session
// @Description Create a new session
// @Tags Sessions
// @Accept json
// @Produce json
// @Param input body SessionPost true "Session object that needs to be added"
// @Success 201 {object} SessionPost
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
// @Param input body SessionPost true "Updated session object"
// @Success 200 {object} SessionPost
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
