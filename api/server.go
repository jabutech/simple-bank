package api

import (
	db "github/jabutech/simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for out banking service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store *db.Store) *Server {
	// Create new server with input store
	server := &Server{store: store}
	// Create new router
	router := gin.Default()

	// Add route
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts", server.listAccount)
	router.GET("/accounts/:id", server.getAccount)

	// Save object router to server.router
	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// function for handle error response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
