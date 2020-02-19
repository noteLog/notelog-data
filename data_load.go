package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func CreateUpdateBlogTable(dbConn string) {
	db, sqlOpenErr := sql.Open("mysql", dbConn)
	if sqlOpenErr != nil {
		log.Fatalf("sqlOpenErr = %v", sqlOpenErr)
	}
	defer db.Close()

	createStmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS posts(post_id INT NOT NULL AUTO_INCREMENT, title VARCHAR(100) NOT NULL UNIQUE, text TEXT, url VARCHAR(100),date DATETIME NOT NULL, PRIMARY KEY (post_id))")
	createStmt.Exec()

	insertStmt, err := db.Prepare("INSERT IGNORE INTO `posts` SET `title`=?, `text`=?, `url`=?,`date`=?;")
	if err != nil {
		panic(err)
	}

	posts := GetBlogPosts()
	var postInserted int64
	for _, post := range posts {
		res, err := insertStmt.Exec(&post.Title, &post.Content, &post.URL, &post.Date)
		if err != nil {
			log.Printf("err = %+v\n", err)

		} else {
			count, err := res.RowsAffected()
			if err != nil {
				log.Printf("err = %+v\n", err)
			}
			postInserted += count
		}

	}
	log.Printf("%v blog posts inserted.", postInserted)
}
