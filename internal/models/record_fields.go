package models

import "time"

type RecordFields struct {
	Role          Role      `json:"role"`
	FullName      string    `json:"full_name"`
	Passport      int       `json:"passport"`
	IsEmployee    bool      `json:"is_employee"`
	ActualCompany string    `json:"actual_company"`
	Reason        string    `json:"reason"`
	AddedAt       time.Time `json:"added_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Worker struct {
	ID            int       `json:"id"`
	FullName      string    `json:"full_name"`
	Passport      int       `json:"passport"`
	IsEmployee    bool      `json:"is_employee"`
	ActualCompany string    `json:"actual_company"`
	Reason        string    `json:"reason"`
	AddedAt       time.Time `json:"added_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
