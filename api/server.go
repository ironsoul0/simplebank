package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/ironsoul0/simplebank/db/sqlc"
	"github.com/ironsoul0/simplebank/token"
	"github.com/ironsoul0/simplebank/util"
)

type Server struct {
	config     util.Config
	store      *db.SQLStore
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store *db.SQLStore) *Server {
	tokenMaker, _ := token.NewJWTMaker(config.TokenSymmetricKey)

	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

	server.setupRouter()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.POST("/transfers", server.createTransfer)

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
