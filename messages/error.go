// messages/errors.go
package messages

const (
	ErrorMessageMethodNotAllowed = "Method not allowed"
	ErrorMessageInvalidRequest    = "Invalid request"
	ErrorMessageMissingFields     = "Missing required fields"
	ErrorMessageFruitNotFound     = "Fruit not found, the ticks do not lead the snake to the fruit position"
	ErrorMessageGameOver          = "Game is over, snake went out of bounds or made an invalid move"
	ErrorMessageMissingParameterW    = "Missing 'w' parameter"
	ErrorMessageMissingParameterH     = "Missing 'h' parameter"
	ErrorMessageInvalidW        = "'w' must be a valid integer"
	ErrorMessageInvalidH          = "'h' must be a valid integer"

)
