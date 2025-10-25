package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var request FindServiceRequest
if err := c.ShouldBindJSON(&request); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
    return
}

