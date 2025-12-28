package routes

import (
	"go-movie-reservation/internal/controllers"
	"go-movie-reservation/internal/services"
	"go-movie-reservation/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Init service
	authService := services.NewAuthService(db)
	movieService := services.NewMovieService(db)
	reservationService := services.NewReservationService(db)
	showtimeService := services.NewShowtimeService(db)

	// Init controller
	authController := controllers.NewAuthController(authService)
	movieController := controllers.NewMovieController(movieService)
	reservationController := controllers.NewReservationController(reservationService, showtimeService)
	showtimeController := controllers.NewShowtimeController(showtimeService)

	// Public routes
	public := router.Group("/api")
	{
		public.POST("/signup", authController.SignUp)
		public.POST("/login", authController.Login)
		public.GET("/movies", movieController.GetMovies)
	}

	// protected route
	protected := router.Group("/api")
	protected.Use(utils.AuthMiddleware())
	{
		protected.GET("/user/reservations", reservationController.GetUserReservations)
		protected.POST("/reservations", reservationController.CreateReservation)
		protected.DELETE("/reservations/:reservationId", reservationController.CancelReservation)
		protected.GET("/showtimes/:showtimeId/seats", reservationController.GetAvailableSeats)
		protected.GET("/movies/:movieID/showtimes", showtimeController.GetShowtimes)
	}

	// Admin route
	admin := router.Group("/api/admin")
	admin.Use(utils.AuthMiddleware(), utils.AdminMiddleware())
	{
		admin.POST("/movies", movieController.CreateMovie)
		admin.PUT("/movies/:movieId", movieController.UpdateMovie)
		admin.DELETE("/movies/:movieId", movieController.DeleteMovie)
		admin.GET("/reservations", reservationController.GetAllReservations)
		admin.POST("/users/:userId/promote", authController.PromoteToAdmin)
		admin.POST("/showtimes", showtimeController.CreateShowtime)
		admin.PUT("/showtimes/:showtimeId", showtimeController.UpdateShowtime)
		admin.DELETE("/showtimes/:showtimeId", showtimeController.DeleteShowtime)
	}

	return router
}
