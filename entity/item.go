package entity

import "github.com/google/uuid"

// Item represents a Item for all subdomains
type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}
