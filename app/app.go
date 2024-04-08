package app

import (
    "github.com/gin-gonic/gin"
)

type GinServer struct {
    router *gin.Engine
}

func NewGinServer() *GinServer {
    // Initialize Gin with the default settings
    router := gin.Default()
    
    // You can add middleware, routes, etc. here
    
    return &GinServer{router: router}
}

func (s *GinServer) Run(address string) error {
    // Start the Gin server
    return s.router.Run(address)
}
