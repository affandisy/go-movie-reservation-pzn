package controllers

import (
	"go-movie-reservation/internal/model"
	"go-movie-reservation/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReservationController struct {
	ReservationService *services.ReservationService
	ShowtimeService    *services.ShowtimeService
}

func NewReservationController(reservationService *services.ReservationService, showtimeService *services.ShowtimeService) *ReservationController {
	return &ReservationController{
		ReservationService: reservationService,
		ShowtimeService:    showtimeService,
	}
}

func (r *ReservationController) GetAllReservations(c *gin.Context) {
	reservations, err := r.ReservationService.GetAllReservations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservations)
}

func (r *ReservationController) GetAvailableSeats(c *gin.Context) {
	showtimeID := c.Param("showtimeId")

	seats, err := r.ShowtimeService.GetAvailableSeats(showtimeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, seats)
}

func (r *ReservationController) CreateReservation(c *gin.Context) {
	var reservation model.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userUUID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	reservation.UserID = userUUID

	newReservation, err := r.ReservationService.CreateReservation(reservation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newReservation)
}

func (r *ReservationController) GetUserReservations(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	reservations, err := r.ReservationService.GetUserReservations(userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservations)
}

func (r *ReservationController) CancelReservation(c *gin.Context) {
	reservationID := c.Param("reservationId")
	if err := r.ReservationService.CancelReservation(reservationID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation cancelled successfully"})
}
