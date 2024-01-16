package api

import (
	"fmt"

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

	// Endpoints for TopicRequest
	r.POST("/topics", a.createTopic)
	r.GET("/topics", a.getAllTopics)
	r.GET("/topics/:id", a.getTopicByID)
	r.GET("/topics/ByIdContainer/:IdContainer", a.getAllTopicsByIdContainer)
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
