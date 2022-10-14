package models

import "time"

type Client struct {
	Id        int       `json:"id"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}
