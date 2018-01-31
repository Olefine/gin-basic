package main

import (
	"fmt"
)

// Article is model for displaying articles
type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []Article{
	Article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	Article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func getAllArticles() []Article {
	return articleList
}

func getArticleByID(id int) (*Article, error) {
	for _, v := range articleList {
		if v.ID == id {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("Could not find article with ID = %d", id)
}
