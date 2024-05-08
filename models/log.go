package models

import (
	"time"
)

type Log struct {
	ID uint
	// TimeStamp shows the time after the server returns a response.
	TimeStamp time.Time
	// StatusCode is HTTP response code.
	StatusCode int
	// Latency is how much time the server cost to process a certain request.
	Latency time.Duration
	// ClientIP equals Context's ClientIP method.
	ClientIP string
	// Method is the HTTP method given to the request.
	Method string
	// Path is a path the client requests.
	Path string
	// ErrorMessage is set if error has occurred in processing the request.
	ErrorMessage string
	// IsTerm shows whether gin's output descriptor refers to a terminal.
	IsTerm bool
	// BodySize is the size of the Response Body
	BodySize int
	// App name
	Application string
}
