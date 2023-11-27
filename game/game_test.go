package game

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewGameAndValidateMoves(t *testing.T) {
	r := gin.Default()

	r.GET("/new", NewGame)
	r.POST("/validate", ValidateMoves)

	t.Run("NewGame", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/new?w=10&h=10", nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var newState State
		err = json.Unmarshal(w.Body.Bytes(), &newState)
		assert.NoError(t, err)

		assert.NotEmpty(t, newState.GameID)
		assert.Equal(t, 10, newState.Width)
		assert.Equal(t, 10, newState.Height)
		assert.Equal(t, 0, newState.Score)
		assert.NotNil(t, newState.Fruit)
		assert.NotNil(t, newState.Snake)
	})

	// Test case for /validate endpoint
	t.Run("ValidateMoves", func(t *testing.T) {
		gameState := State{
			GameID: "123",
			Width:  10,
			Height: 10,
			Score:  0,
			Fruit:  Fruit{X: 5, Y: 5},
			Snake:  Snake{X: 0, Y: 0, VelX: 1, VelY: 0},
			Ticks: []Tick{
				{VelX: 1, VelY: 0},
				{VelX: 1, VelY: 0},
			},
		}
		reqBody, err := json.Marshal(gameState)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/validate", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var updatedState State
		err = json.Unmarshal(w.Body.Bytes(), &updatedState)
		assert.NoError(t, err)

		fmt.Printf("Updated State: %+v\n", updatedState)

		assert.NotEmpty(t, updatedState.GameID)
		assert.Equal(t, 1, updatedState.Score)
		assert.NotNil(t, updatedState.Fruit)
		assert.NotNil(t, updatedState.Snake)

		assert.Equal(t, 2, updatedState.Snake.X)
		assert.Equal(t, 0, updatedState.Snake.Y)
	})
}
