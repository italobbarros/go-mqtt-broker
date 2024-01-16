package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/italobbarros/go-mqtt-broker/docs"
	"github.com/italobbarros/go-mqtt-broker/internal/api/models"
	"github.com/italobbarros/go-mqtt-broker/internal/api/routes"
	"github.com/italobbarros/go-mqtt-broker/pkg/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// API representa a interface da API do servidor
type API struct {
	routes *routes.Routes
	logger *logger.Logger
}

// Init inicializa a API com as rotas e inicia o servidor
func NewAPI() *API {
	logger := logger.NewLogger("API")
	return &API{
		logger: logger,
	}
}

func (a *API) initDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("postgresql://teste:teste@localhost:5433/broker?sslmode=disable"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.Container{}, &models.Topic{}, &models.Subscription{}, &models.Publication{}, &models.Session{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (a *API) Init() {
	db, err := a.initDB()
	if err != nil {
		a.logger.Error("Erro init API:%s", err.Error())
		return
	}
	a.routes = routes.NewRoutes(a.logger, db)

	r := gin.Default()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Endpoints for ContainerPost
	r.POST("/containers", a.routes.CreateContainer)
	r.GET("/containers", a.routes.GetAllContainers)
	r.GET("/containers/:id", a.routes.GetContainerByID)
	r.PUT("/containers/:id", a.routes.UpdateContainer)

	// Endpoints for TopicRequest
	r.POST("/topics", a.routes.CreateTopic)
	r.GET("/topics", a.routes.GetAllTopics)
	r.GET("/topics/:id", a.routes.GetTopicByID)
	r.GET("/topics/ByIdContainer/:IdContainer", a.routes.GetAllTopicsByIdContainer)
	r.PUT("/topics/:id", a.routes.UpdateTopic)

	// Endpoints for Subscription
	r.POST("/subscriptions", a.routes.CreateSubscription)
	r.GET("/subscriptions", a.routes.GetAllSubscriptions)
	r.GET("/subscriptions/:id", a.routes.GetSubscriptionByID)
	r.PUT("/subscriptions/:id", a.routes.UpdateSubscription)

	// Endpoints for Publication
	r.POST("/publications", a.routes.CreatePublication)
	r.GET("/publications", a.routes.GetAllPublications)
	r.GET("/publications/:id", a.routes.GetPublicationByID)
	r.PUT("/publications/:id", a.routes.UpdatePublication)

	// Endpoints for Session
	r.POST("/sessions", a.routes.CreateSession)
	r.GET("/sessions", a.routes.GetAllSessions)
	r.GET("/sessions/:id", a.routes.GetSessionByID)
	r.PUT("/sessions/:id", a.routes.UpdateSession)

	r.Run(":8080")
}
