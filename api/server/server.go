package server

import (
	"github.com/ASeegull/edriver-space/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
	*echo.Echo
	Users []model.User
	Cars  []model.Car
	Data  model.Data
}

// NewServer starts new Server
func NewServer() *Server {
	s := &Server{
		Echo:  echo.New(),
		Users: []model.User{},
		Cars:  []model.Car{},
		Data:  model.Data{},
	}
	// All possible routes
	s.routes()
	return s
}

// routes stores all possible routes
func (s *Server) routes() {
	s.GET("/", s.hello())             // Home
	s.GET("/Version", s.getVersion()) // Get project version

	// Fines group
	fines := s.Group("/fines")
	fines.POST("/upload", s.uploadXMLData())    // Upload parking fines data from xml file
	fines.GET("", s.getParkingFines())          // Get all parking fines
	fines.GET("/:id", s.getParkingFine())       // Get parking fine by id
	fines.DELETE("/:id", s.deleteParkingFine()) // Remove parking fine by id

	// Users CRUD
	s.GET("/Users", s.getUsers())         // Get all users
	s.GET("/User/:id", s.getUser())       // Get user by id
	s.POST("/User", s.createUser())       // Create new user
	s.PATCH("/User/:id", s.updateUser())  // Update user data
	s.DELETE("/User/:id", s.deleteUser()) // Delete user

	// Cars CRUD
	s.GET("/Cars", s.getCars())         // Get all cars
	s.GET("/Car/:id", s.getCar())       // Get car by id
	s.POST("/Car", s.createCar())       // Create new car
	s.PATCH("/Car/:id", s.updateCar())  // Update car data
	s.DELETE("/Car/:id", s.deleteCar()) // Delete car
}

// getVersion returns current version of the app
func (s *Server) getVersion() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		version := "0.0" // Use environment variable instead
		return ctx.JSON(http.StatusOK, version)
	}
}

// hello greets developers group
func (s *Server) hello() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		hi := "Hello Lv-644.Go!"
		return ctx.JSON(http.StatusOK, hi)
	}
}
