package api

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/italobbarros/go-mqtt-broker/docs"
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
	dns := os.Getenv("DB_ADDRESS")
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database")
	}

	// Migrate the schema
	//err = db.AutoMigrate(&models.Container{}, &models.Topic{}, &models.Subscription{}, &models.Publish{}, &models.Session{})
	//if err != nil {
	//	return nil, err
	//}
	sqlDB, err := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(50)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(1000)
	return db, nil
}

func (a *API) Init() {
	db, err := a.initDB()
	if err != nil {
		a.logger.Error("Erro init API:%s", err.Error())
		return
	}
	a.routes = routes.NewRoutes(a.logger, db)
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Endpoints for ContainerPost
	r.POST("/containers", a.routes.CreateContainer)
	r.GET("/containers/all", a.routes.GetAllContainers)
	r.GET("/containers/:id", a.routes.GetContainerByID)
	r.DELETE("/containers/:Name", a.routes.DeleteContainerByName)

	// Endpoints for TopicRequest
	r.GET("/topics/all", a.routes.GetAllTopics)
	r.GET("/topics/:id", a.routes.GetTopicByID)
	r.GET("/topics", a.routes.GetTopicsByName)
	r.PUT("/topics/:id", a.routes.UpdateTopic)

	// Endpoints for Subscription
	r.POST("/subscriptions", a.routes.CreateSubscription)
	r.GET("/subscriptions", a.routes.GetAllSubscriptions)
	r.GET("/subscriptions/:id", a.routes.GetSubscriptionByID)
	r.PUT("/subscriptions/:id", a.routes.UpdateSubscription)

	// Endpoints for Publish
	r.POST("/publisher", a.routes.CreatePublish)
	r.GET("/publisher/all", a.routes.GetAllPublisher)
	r.GET("/publisher/historic", a.routes.GetPublishByTopicName)

	// Endpoints for Session
	r.POST("/sessions", a.routes.CreateSession)
	r.GET("/sessions", a.routes.GetSessionByClientId)
	r.GET("/sessions/all", a.routes.GetAllSessions)
	r.PUT("/sessions", a.routes.UpdateSession)
	r.DELETE("/sessions", a.routes.DeleteSessionByClientId)

	r.Run(os.Getenv("API_ADDRESS"))
}
