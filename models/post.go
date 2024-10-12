package models

import "time"

type Post struct {
	PostId    int
	PostTitle string
	PostBody  string
	PostImg   string
	CreatedAt time.Time
}
