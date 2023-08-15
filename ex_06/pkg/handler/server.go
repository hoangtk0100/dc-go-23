package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/business"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/db"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/repository"
	router "github.com/hoangtk0100/dc-go-23/ex_06/pkg/route"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/token"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/util"
	custom_validator "github.com/hoangtk0100/dc-go-23/ex_06/pkg/validation"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	config     util.Config
	tokenMaker token.TokenMaker
	router     *gin.Engine
	store      *db.DB
	repo       repository.Repository
	business   business.Business
}

func NewServer(config util.Config) *Server {
	server := &Server{
		config: config,
		router: gin.Default(),
	}

	tokenMaker, err := token.NewJWTMaker(config.SecretKey, config.AccessTokenExpiresIn, config.RefreshTokenExpiresIn)
	if err != nil {
		log.Fatal(err)
	}

	server.tokenMaker = tokenMaker
	server.store = db.NewDB()
	server.repo = repository.NewRepository(server.store)
	server.business = business.NewBusiness(server.repo, tokenMaker)

	// Add custom validator for Currency, WeightUnit
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", custom_validator.ValidateCurrency)
		v.RegisterValidation("weight_unit", custom_validator.ValidateWeightUnit)
	}

	router.SetupRoutes(server)
	return server
}

func (server *Server) GetRouter() *gin.Engine {
	return server.router
}

func (server *Server) GetRepository() repository.Repository {
	return server.repo
}

func (server *Server) GetTokenMaker() token.TokenMaker {
	return server.tokenMaker
}

func (server *Server) Start() {
	srv := &http.Server{
		Addr:    server.config.ServerAddress,
		Handler: server.router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server closed unexpectedly: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exited")
}
