package controllers

import (
	"go-movie-reservation/internal/model"
	"go-movie-reservation/internal/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MovieController struct {
	MovieService *services.MovieService
}

func NewMovieController(movieService *services.MovieService) *MovieController {
	return &MovieController{
		MovieService: movieService,
	}
}

type CreateMovieRequest struct {
	Title       string `json:"title" binding:"required"`
	Director    string `json:"director" binding:"required"`
	ReleaseDate string `json:"releaseDate" binding:"required"`
	Duration    int    `json:"duration" binding:"required"`
	Description string `json:"description" binding:"required"`
	Genre       string `json:"genre" binding:"required"`
	PosterURL   string `json:"posterURL" binding:"required"`
}

func (m *MovieController) CreateMovie(c *gin.Context) {
	var req CreateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	releaseDate, err := time.Parse("2006-01-02", req.ReleaseDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid release date format"})
		return
	}

	movie := model.Movie{
		ID:          uuid.New(),
		Title:       req.Title,
		Director:    req.Director,
		ReleaseDate: releaseDate,
		Duration:    req.Duration,
		Description: req.Description,
		Genre:       model.Genre(req.Genre),
		PosterImage: req.PosterURL,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createdMovie, err := m.MovieService.CreateMovie(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdMovie)
}

func (m *MovieController) GetMovies(c *gin.Context) {
	movies, err := m.MovieService.GetMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (m *MovieController) UpdateMovie(c *gin.Context) {
	var movie model.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedMovie, err := m.MovieService.UpdateMovie(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedMovie)
}

func (m *MovieController) DeleteMovie(c *gin.Context) {
	movieID := c.Param("movieId")
	if err := m.MovieService.DeleteMovie(movieID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
