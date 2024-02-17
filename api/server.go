package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/yyht/simplebank/db/sqlc"
	"github.com/yyht/simplebank/token"
	"github.com/yyht/simplebank/util"
)

type Server struct {
	config     util.Config // when creating token, we need to use the duration data in config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine // send each API request to the correct handler
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err) // %w, to wrap original error
	}

	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

	// binding.Validator.Engine(), get currently used validator engine
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 1st param is validation tag, which can be used in JSON tags
		// 2nd param, validation funciton
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	// create a group of routes
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("accounts", server.createAccount)

	authRoutes.GET("accounts/:id", server.getAccount)

	authRoutes.GET("accounts", server.listAccount)

	authRoutes.POST("transfers", server.createTransfer)

	server.router = router
}

// Start runs the HTTP server on a specific address. to start listening for API requests
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H { // gin.H: type map[string]interface{}
	return gin.H{"error": err.Error()}
}
