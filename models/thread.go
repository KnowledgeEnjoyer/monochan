package models

import (
	"log"
	"strconv"
	"time"

	"github.com/KnowledgeEnjoyer/monochan/db"
)

type Thread struct {
	ThreadId    int
	Board       string
	ThreadTitle string
	ThreadBody  string
	ThreadImg   string
	RecentPosts []Post
	CreatedAt   time.Time
	BumpedAt    time.Time
}

func NewThread(board, title, body, img string, createAt time.Time) *Thread {
	return &Thread{
		Board:       board,
		ThreadTitle: title,
		ThreadBody:  body,
		ThreadImg:   img,
		CreatedAt:   time.Now(),
		BumpedAt:    time.Now(),
	}
}

func GetAllThreads() []Thread {
	db := db.GetDBInstance()
	queryT := `
		SELECT ThreadId, Board, ThreadTitle, ThreadBody, ThreadImg, Threads.CreatedAt
			FROM Threads;
	`
	var threads []Thread

	rowsT, err := db.Query(queryT)
	if err != nil {
		log.Fatal(err)
	}

	defer rowsT.Close()

	for rowsT.Next() {
		var t Thread
		err := rowsT.Scan(&t.ThreadId, &t.Board, &t.ThreadTitle, &t.ThreadBody, &t.ThreadImg, &t.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		threads = append(threads, t)
	}

	for i, v := range threads {
		queryP := "SELECT PostId, PostTitle, PostBody, PostImg, CreatedAt " +
			"FROM Posts WHERE ThreadId = " + strconv.Itoa(v.ThreadId) + ";"

		rowsP, err := db.Query(queryP)
		if err != nil {
			log.Fatal(err)
		}

		defer rowsP.Close()

		for rowsP.Next() {
			var p Post
			rowsP.Scan(&p.PostId, &p.PostTitle, &p.PostBody, &p.PostImg, &p.CreatedAt)
			v.RecentPosts = append(v.RecentPosts, p)
		}

		threads[i].RecentPosts = append(threads[i].RecentPosts, v.RecentPosts...)
	}

	return threads

}
