package controllers

import (
	"go-movie-reservation/internal/model"
	"go-movie-reservation/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShowtimeController struct {
	showtimeService *services.ShowtimeService
}

func NewShowtimeController(showtimeService *services.ShowtimeService) *ShowtimeController {
	return &ShowtimeController{
		showtimeService: showtimeService,
	}
}

func (sc *ShowtimeController) GetAvailableSeats(c *gin.Context) {
	showtimeID := c.Param("showtimeID")
	availableSeats, err := sc.showtimeService.GetAvailableSeats(showtimeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"available_seats": availableSeats})
}

func (s *ShowtimeController) CreateShowtime(c *gin.Context) {
	var showtime model.Showtime
	if err := c.ShouldBindJSON(&showtime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdShowtime, err := s.showtimeService.CreateShowtime(showtime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdShowtime)
}

func (s *ShowtimeController) GetShowtimes(c *gin.Context) {
	movieID := c.Param("movieID")
	showtimes, err := s.showtimeService.GetShowtimes(movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, showtimes)
}

func (s *ShowtimeController) UpdateShowtime(c *gin.Context) {
	var showtime model.Showtime
	if err := c.ShouldBindJSON(&showtime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedShowtime, err := s.showtimeService.UpdateShowtime(showtime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedShowtime)
}

func (s *ShowtimeController) DeleteShowtime(c *gin.Context) {
	showtimeID := c.Param("showtimeID")
	err := s.showtimeService.DeleteShowtime(showtimeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Showtime deleted successfully"})
}
