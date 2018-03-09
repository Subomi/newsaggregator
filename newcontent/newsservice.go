package newcontent

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/LordRahl90/newsAPIAggregator/utility"
	"gopkg.in/mgo.v2/bson"
)

//NewsSource struct
type NewsSource struct {
	ID   string
	Name string
}

//NewsContent -- struct of the newscontent
type NewsContent struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Source      NewsSource    `json:"source" bson:"source"`
	Author      string        `json:"author" bson:"author"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	NewsLink    string        `json:"url" bson:"news_link"`
	Thumbnail   string        `json:"urlToImage" bson:"thumbnail"`
	PublishedAt string        `json:"publishedAt" bson:"publishedAt"`
}

//NewsResponse struct from the API
type NewsResponse struct {
	Status       string
	TotalResults int
	Articles     []NewsContent
}

//NewsContents Array of NewsConrent
type NewsContents []NewsContent

//GetAllContents - function to retrieve all the available contents
func GetAllContents() {
	categories := []string{"business", "technology", "entertainment", "general", "health", "sports", "science"}

	for _, v := range categories {
		go GetCategoryContent(v)
	}
}

//GetCategoryContent - function to retrieve each content based on the category
func GetCategoryContent(category string) {
	client := utility.GetHTTPConnection()
	url := "https://newsapi.org/v2/top-headlines?country=us&category=" + category + "&apiKey=9d206d9be4174155bf59edb914ce4101"

	response, err := client.Get(url)
	utility.CheckError(err)

	if response.StatusCode != 200 {
		fmt.Println("We cant download for ", category)
		log.Println(response.Body)
		return
	}
	defer response.Body.Close()
	responseBytes, err := ioutil.ReadAll(response.Body)
	utility.CheckError(err)

	// bodyString := string(responseBytes)

	//lets convert the bodystring to an array of news details
	var newsResponse NewsResponse
	err = json.Unmarshal(responseBytes, &newsResponse)
	utility.CheckError(err)

	log.Println(newsResponse.Status)
	log.Println(newsResponse.TotalResults)
	log.Println(newsResponse.Articles)
	log.Println("downloaded for ", category)

	//lets save the record in the database
	for _, article := range newsResponse.Articles {
		article.ID = bson.NewObjectId()
		status := make(chan bool)
		go article.KeepInDatabase(status)
		if !<-status {
			log.Println("An error occurred while keeping ", article.NewsLink)
		}
	}

	fmt.Println("downloaded for ", category)
}

//KeepInDatabase - function to keep the the newscontent inside the database
func (content NewsContent) KeepInDatabase(c chan bool) {
	//lets just check against the newsurl
	response := false
	db := utility.GetConnection()
	collection := db.C("news_headlines")
	count, err := collection.Find(bson.M{"$and": []bson.M{
		bson.M{"news_link": content.NewsLink},
		bson.M{"author": content.Author},
	}}).Count()
	utility.CheckError(err)

	if count <= 0 {
		err = collection.Insert(&content)
		utility.CheckError(err)
		response = true
	}

	c <- response
}
