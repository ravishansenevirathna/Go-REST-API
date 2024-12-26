package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/service"
	"strconv"
)

type RecommendationController struct {
	Service *service.RecommendationService
}

func NewRecommendationController(service *service.RecommendationService) *RecommendationController {
	return &RecommendationController{Service: service}
}

func (c *RecommendationController) GetRecommendations(ctx *gin.Context) {
	userID, exists := ctx.Get("userID") // Assuming userID is set via middleware
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	limitStr := ctx.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
		return
	}

	recommendations, err := c.Service.RecommendSongs(userID.(uint), limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, recommendations)
}
