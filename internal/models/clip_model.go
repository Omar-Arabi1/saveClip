package models

type Clip struct{
	Body string `json:"body"`
	Label string `json:"label"`
	Priority int `json:"priority"`
	CreationDate string `json:"creationDate"`
}