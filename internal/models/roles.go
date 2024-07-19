package models

type Role struct {
	Name       string   `json:"name"`
	Operations []string `json:"operations"`
}
