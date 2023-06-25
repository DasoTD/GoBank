package api

import (
	db "github.com/dasotd/gobank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	bank db.Bank
	router *gin.Engine
}

func NewServer(bank db.Bank) *Server {
	server := &Server{bank: bank}
	router := gin.Default()

	//routes
	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.DELETE("/account/:id", server.deleteAccount)
	router.PUT("/account/:id", server.updateAccount)

	router.POST("/entry", server.createEntry)
	router.GET("/entry/:id", server.getEntry)
	router.GET("/entries/:accountId", server.listEntry)

	router.POST("/transfer", server.createTransfer)

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// // // Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}