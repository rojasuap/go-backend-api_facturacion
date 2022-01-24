package promotion

import (
	"time"
)

// Promotion created by a user.
type Promotion struct {
	ID          uint      `json:"id,omitempty"`
	Description string    `json:"description,omitempty"`
	Percentage  float32   `json:"percentage,string,omitempty"`
	Start_date  string    `json:"start_date,omitempty"`
	End_date    string    `json:"end_date,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
