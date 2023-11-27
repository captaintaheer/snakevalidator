// game/game.go
package game

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"snakevalidator/messages"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

// State represents the game state.
type State struct {
	GameID string `json:"gameId"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Score  int    `json:"score"`
	Fruit  Fruit  `json:"fruit"`
	Snake  Snake  `json:"snake"`
	Ticks  []Tick `json:"ticks"`
}

// Fruit represents the fruit position.
type Fruit struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Snake represents the snake position and velocity.
type Snake struct {
	X    int `json:"x"`
	Y    int `json:"y"`
	VelX int `json:"velX"`
	VelY int `json:"velY"`
}

// Tick represents a single movement of the snake.
type Tick struct {
	VelX int `json:"velX"`
	VelY int `json:"velY"`
}

func generateGameID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d", rand.Intn(100000))
}

func generateRandomFruitPosition(width, height int) Fruit {
	rand.Seed(time.Now().UnixNano())
	return Fruit{
		X: rand.Intn(width),
		Y: rand.Intn(height),
	}
}

func NewGame(c *gin.Context) {
	widthStr, exists := c.GetQuery("w")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": messages.ErrorMessageInvalidRequest, "details": messages.ErrorMessageMissingParameterW})
		return
	}

	heightStr, exists := c.GetQuery("h")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": messages.ErrorMessageInvalidRequest, "details": messages.ErrorMessageMissingParameterH})
		return
	}

	width, err := strconv.Atoi(widthStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": messages.ErrorMessageInvalidRequest, "details": messages.ErrorMessageInvalidW})
		return
	}

	height, err := strconv.Atoi(heightStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": messages.ErrorMessageInvalidRequest, "details": messages.ErrorMessageInvalidH})
		return
	}

	if c.Request.Method != "GET" {
		c.AbortWithStatus(http.StatusMethodNotAllowed)
		return
	}

	gameID := generateGameID()
	fruit := generateRandomFruitPosition(width, height)

	state := State{
		GameID: gameID,
		Width:  width,
		Height: height,
		Score:  0,
		Fruit:  fruit,
		Snake:  Snake{X: 0, Y: 0, VelX: 1, VelY: 0},
	}

	c.JSON(http.StatusOK, state)
}

// ValidateMoves validates the moves and updates the game state.
func ValidateMoves(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": messages.ErrorMessageMethodNotAllowed})
		return
	}

	var requestBody State
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		var details string
		if errors.Is(err, binding.ErrMultiFileHeaderLenInvalid) {
			details = err.Error()
		} else {
			details = messages.ErrorMessageInvalidRequest
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": messages.ErrorMessageInvalidRequest, "details": details})
		return
	}

	if err := validateRequiredFields(requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": messages.ErrorMessageInvalidRequest, "details": err.Error()})
		return
	}

	for _, tick := range requestBody.Ticks {
		requestBody.Snake.X += tick.VelX
		requestBody.Snake.Y += tick.VelY

		if isOutOfBounds(requestBody.Snake, requestBody.Width, requestBody.Height) {
			c.JSON(http.StatusBadRequest, gin.H{"error": messages.ErrorMessageGameOver})
			return
		}

		if isFruitReached(requestBody.Snake, requestBody.Fruit) {
			requestBody.Fruit = generateRandomFruitPosition(requestBody.Width, requestBody.Height)
			requestBody.Score++
		}
	}

	c.JSON(http.StatusOK, requestBody)
}

func validateRequiredFields(s State) error {
	if s.GameID == "" || s.Width <= 0 || s.Height <= 0 {
		return errors.New(messages.ErrorMessageMissingFields)
	}
	return nil
}

func isOutOfBounds(snake Snake, width, height int) bool {
	return snake.X < 0 || snake.X >= width || snake.Y < 0 || snake.Y >= height
}

func isFruitReached(snake Snake, fruit Fruit) bool {
	return snake.X == fruit.X && snake.Y == fruit.Y
}
