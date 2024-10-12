package models

import (
	"log"
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
	query := `
		SELECT Threads.ThreadId, Board, ThreadTitle, ThreadBody, ThreadImg, Threads.CreatedAt, PostId, PostTitle, PostBody, PostImg, Posts.CreatedAt
			FROM Threads
			INNER JOIN 	Posts
			ON Threads.ThreadId = Posts.ThreadId;
	`
	var threads []Thread

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var t Thread
		var p Post
		err := rows.Scan(&t.ThreadId, &t.Board, &t.ThreadTitle, &t.ThreadBody, &t.ThreadImg, &t.CreatedAt, &p.PostId, &p.PostTitle, &p.PostBody, &p.PostImg, &p.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		t.RecentPosts = append(t.RecentPosts, p)
		threads = append(threads, t)
	}

	log.Println(threads)

	return threads

	// return []Thread{
	// 	{Board: "b", ThreadTitle: "Animefag", ThreadBody: "Sopinha, bê! CHANZINHO NOVO SAINDO DO FORNINHO BÊ :3", ThreadImg: "/", RecentPosts: posts, CreatedAt: time.Now(), BumpedAt: time.Now()},
	// 	{Board: "b", ThreadTitle: "SEXTA BUCETAL!", ThreadBody: "Que dia é hoje?", ThreadImg: "/", RecentPosts: posts, CreatedAt: time.Now(), BumpedAt: time.Now()},
	// }
}
