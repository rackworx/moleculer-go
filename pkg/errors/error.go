package errors

// should be returned by transporters when a connection error occurs
type ConnectionError struct {
	WasReconnect bool
	Err          error
}

// Moelculer can define errors that are allowed to be retried when thrown
type Error struct {
	Retryable bool
	Code      int
	Err       error
}
