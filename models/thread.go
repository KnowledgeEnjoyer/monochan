package models

import "time"

type Thread struct {
	Board       string
	threadTitle string
	ThreadBody  string
	threadImg   string
	createdAt   time.Time
	bumpedAt    time.Time
}

func NewThread(board, title, body, img string, createAt time.Time) *Thread {
	return &Thread{
		Board:       board,
		threadTitle: title,
		ThreadBody:  body,
		threadImg:   img,
		createdAt:   time.Now(),
		bumpedAt:    time.Now(),
	}
}

func GetAllThreads() []Thread {
	return []Thread{
		{Board: "b", threadTitle: "Animefag", ThreadBody: "Sopinha, bê! CHANZINHO NOVO SAINDO DO FORNINHO BÊ :3", threadImg: "/", createdAt: time.Now(), bumpedAt: time.Now()},
		{Board: "b", threadTitle: "SEXTA BUCETAL!", ThreadBody: "Que dia é hoje?", threadImg: "/", createdAt: time.Now(), bumpedAt: time.Now()},
	}
}
