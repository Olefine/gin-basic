package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOk := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)

		pageOk := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOk && pageOk
	})
}

func TestShowArticlePageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	saveLists()
	articleList = append(articleList, Article{ID: 3, Title: "Hello world", Content: "Hello test article"})
	defer restoreLists()

	r.GET("/articles/:article_id", getArticle)

	req, _ := http.NewRequest("GET", "/articles/3", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOk := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)

		pageOk := err == nil && strings.Index(string(p), "<p>Hello test article</p>") > 0

		return statusOk && pageOk
	})
}
