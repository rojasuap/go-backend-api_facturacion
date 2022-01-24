package bill

import "time"

// Promotion created by a user.
type Bill struct {
	ID           uint      `json:"id,omitempty"`
	CreatedAt    string    `json:"created_at,omitempty"`
	Full_payment float32   `json:"full_payment,string,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
