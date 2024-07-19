package models

import "time"

type RecordFields struct {
	Role             Role      `json:"role"`
	FullName         string    `json:"full_name"`
	Passport         Passport  `json:"passport"`
	PreviousPassport Passport  `json:"previous_passport"`
	IsEmployee       bool      `json:"is_employee"`
	ActualCompany    string    `json:"actual_company"`
	Reason           string    `json:"reason"`
	AddedAt          time.Time `json:"added_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Passport struct {
	PassportSeries int `json:"passport_series"`
	PassportNumber int `json:"passport_number"`
}
