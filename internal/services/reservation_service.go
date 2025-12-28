package services

import (
	"go-movie-reservation/internal/model"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReservationService struct {
	DB *gorm.DB
}

func NewReservationService(db *gorm.DB) *ReservationService {
	return &ReservationService{
		DB: db,
	}
}

func (r *ReservationService) GetAllReservations() ([]*model.Reservation, error) {
	var reservations []*model.Reservation
	if err := r.DB.Find(&reservations).Error; err != nil {
		return nil, err
	}

	return reservations, nil
}

func (r *ReservationService) GetAvailableSeats(showtimeID string) ([]int, error) {
	var showtime model.Showtime
	if err := r.DB.Where("id = ?", showtimeID).First(&showtime).Error; err != nil {
		return nil, err
	}

	reservations := make([]model.Reservation, 0)
	if err := r.DB.Where("showtime_id = ?", showtimeID).Find(&reservations).Error; err != nil {
		return nil, err
	}

	var bookedSeats []int
	for _, reservation := range reservations {
		seatNumbers := strings.Split(reservation.SeatNumbers, ",")
		for _, seatStr := range seatNumbers {
			seat, err := strconv.Atoi(strings.TrimSpace(seatStr))
			if err != nil {
				return nil, err
			}
			bookedSeats = append(bookedSeats, seat)
		}
	}

	availableSeats := make([]int, showtime.AvailableSeats)
	for i := 0; i < showtime.AvailableSeats; i++ {
		availableSeats[i] = i + 1
	}

	for _, bookedSeat := range bookedSeats {
		for i, seat := range availableSeats {
			if seat == bookedSeat {
				availableSeats = append(availableSeats[:i], availableSeats[i+1:]...)
				break
			}
		}
	}

	return availableSeats, nil
}

func (r *ReservationService) CreateReservation(reservation model.Reservation) (*model.Reservation, error) {
	reservation.ID = uuid.New()
	if err := r.DB.Create(&reservation).Error; err != nil {
		return nil, err
	}

	return &reservation, nil
}

func (r *ReservationService) GetUserReservations(userID string) ([]*model.Reservation, error) {
	var reservations []*model.Reservation
	if err := r.DB.Where("user_id=?", userID).Find(&reservations).Error; err != nil {
		return nil, err
	}

	return reservations, nil
}

func (r *ReservationService) CancelReservation(reservationID string) error {
	if err := r.DB.Where("id = ?", reservationID).Delete(&model.Reservation{}).Error; err != nil {
		return err
	}

	return nil
}
