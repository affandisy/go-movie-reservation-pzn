package services

import (
	"go-movie-reservation/internal/model"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShowtimeService struct {
	DB *gorm.DB
}

func NewShowtimeService(db *gorm.DB) *ShowtimeService {
	return &ShowtimeService{
		DB: db,
	}
}

func (s *ShowtimeService) CreateShowtime(showtime model.Showtime) (*model.Showtime, error) {
	showtime.ID = uuid.New()
	if err := s.DB.Create(&showtime).Error; err != nil {
		return nil, err
	}

	return &showtime, nil
}

func (s *ShowtimeService) GetShowtimes(movieID string) ([]*model.Showtime, error) {
	var showtimes []*model.Showtime
	if err := s.DB.Where("movie_id=?", movieID).Find(&showtimes).Error; err != nil {
		return nil, err
	}

	return showtimes, nil
}

func (s *ShowtimeService) UpdateShowtime(showtime model.Showtime) (*model.Showtime, error) {
	if err := s.DB.Save(&showtime).Error; err != nil {
		return nil, err
	}

	return &showtime, nil
}

func (s *ShowtimeService) DeleteShowtime(showtimeID string) error {
	if err := s.DB.Where("id = ?", showtimeID).Delete(&model.Showtime{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *ShowtimeService) GetAvailableSeats(showtimeID string) ([]string, error) {
	var seats string
	err := s.DB.Table("showtimes").Select("available_seats").Where("id = ?", showtimeID).Scan(&seats).Error
	if err != nil {
		return nil, err
	}

	seatArray := strings.Split(seats, ",")
	for i, seat := range seatArray {
		seatArray[i] = strings.TrimSpace(seat)
	}

	return seatArray, nil
}
