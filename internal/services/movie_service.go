package services

import (
	"go-movie-reservation/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MovieService struct {
	DB *gorm.DB
}

func NewMovieService(db *gorm.DB) *MovieService {
	return &MovieService{
		DB: db,
	}
}

func (m *MovieService) CreateMovie(movie model.Movie) (*model.Movie, error) {
	movie.ID = uuid.New()
	if err := m.DB.Create(&movie).Error; err != nil {
		return nil, err
	}

	return &movie, nil
}

func (m *MovieService) GetMovies() ([]*model.Movie, error) {
	var movies []*model.Movie
	if err := m.DB.Find(&movies).Error; err != nil {
		return nil, err
	}

	return movies, nil
}

func (m *MovieService) UpdateMovie(movie model.Movie) (*model.Movie, error) {
	if err := m.DB.Save(&movie).Error; err != nil {
		return nil, err
	}

	return &movie, nil
}

func (m *MovieService) DeleteMovie(movieID string) error {
	if err := m.DB.Where("id = ?", movieID).Delete(&model.Movie{}).Error; err != nil {
		return err
	}

	return nil
}
