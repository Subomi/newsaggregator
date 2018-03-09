package newcontent

import (
	"log"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestCreateContent(t *testing.T) {
	newContent := NewsContent{
		ID:          bson.NewObjectId(),
		Source:      "USA Today",
		Author:      "Andrea Mandell",
		Title:       "'Black Panther' first reactions: It's 'astonishing,' 'iconic' and 'will save blockbusters'",
		Description: "USA TODAY's Brian Truitt hails the first superhero movie to star a black character \"simply awesome\" and a \"top 5 all-time Marvel movie, easy.\"",
		NewsLink:    "https://www.usatoday.com/story/life/entertainthis/2018/01/30/black-panther-first-reactions-declare-film-astonishing-iconic-and-save-blockbusters/1077768001/",
		Thumbnail:   "https://www.gannett-cdn.com/-mm-/b9e421ed3d9567364df962a5de79c6fc5a35c8b4/c=0-169-2988-1857&r=x1683&c=3200x1680/local/-/media/2018/01/27/USATODAY/USATODAY/636526458976202193-BlackPanther596d2f0946755.jpg",
		PublishedAt: "2018-01-30T08:23:00Z",
	}

	response := newContent.KeepInDatabase()
	log.Fatal(response)
}
