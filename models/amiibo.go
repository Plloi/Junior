package models

// Amiibo represents an Amiibo figure
type Amiibo struct {
	ID          int64
	Serial      string //Maybe change to int, check figures
	Description string
}
