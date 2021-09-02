package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"main/src/internal/data"
	"main/src/internal/repository"
	"main/src/service"
	"net/http"
)

type Server struct {
	e    *echo.Echo
	port string
}

func NewServer(lisPort int) *Server {
	return &Server{
		e:    echo.New(),
		port: fmt.Sprintf(":%d", lisPort),
	}
}

func (s *Server) Start() {
	s.InitRoutes()
	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000/", "http://localhost:3000"},
		AllowHeaders: []string{},
	}))
	s.e.Logger.Fatal(s.e.Start(s.port))
}

func (s *Server) InitRoutes() {
	s.e.POST("/signup", signUp)
	s.e.PUT("/login", login)
	s.e.PUT("/logout", logout)
	s.e.GET("/get-last-prices", getPriceUpdates)
	s.e.GET("/get-user-positions", getUserPositions)
	s.e.POST("/open-position", openUserPosition)
	s.e.PUT("/close-position", closeUserPosition)
	s.e.GET("/user-balance", getUserBalance)
}

func signUp(c echo.Context) error {
	return nil
}

func login(c echo.Context) error {
	var logInReq data.LogInRequest
	if c.Bind(&logInReq) != nil {
		return echo.ErrBadRequest
	}

	resp, err := service.Login(logInReq.Username, logInReq.Password)

	if err != nil {
		log.Error(err)
	}

	response := make(map[string]string)
	response["token"] = resp
	response["username"] = logInReq.Username
	response["password"] = logInReq.Password

	return c.JSON(http.StatusOK, response)
}

func logout(c echo.Context) error {
	return nil
}

func getPriceUpdates(c echo.Context) error {
	prices := make([]data.SymbolPrice, 0, len(repository.Repository.Data))
	for _, v := range repository.Repository.Data {
		prices = append(prices, v)
	}
	return c.JSON(http.StatusOK, prices)
}

func getUserPositions(c echo.Context) error {
	token := c.QueryParam("token")
	if token == "" {
		return echo.ErrUnauthorized
	}

	result, err := service.GetUserData(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, result)
}

func openUserPosition(c echo.Context) error {
	request := data.OpenPosRequest{}

	if err := c.Bind(&request); err != nil {
		return echo.ErrBadRequest
	}

	if request.Token == "" {
		return echo.ErrUnauthorized
	}

	res, err := service.OpenPosition(repository.Repository.Data[request.Symbol], request.IsBay, request.Token)
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, res)
}

func closeUserPosition(c echo.Context) error {
	pos := data.Position{}
	if err := c.Bind(&pos); err != nil {
		return echo.ErrBadRequest
	}

	token := c.QueryParam("token")

	if token == "" {
		return echo.ErrUnauthorized
	}

	if _, err := service.ClosePosition(pos, token); err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, "OK")
}

func getUserBalance(c echo.Context) error {
	token := c.QueryParam("token")
	balanceResp, err := service.GetUserBalance(token)
	if err != nil {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, balanceResp)
}
