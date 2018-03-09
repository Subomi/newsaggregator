package utility

import (
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
)

//Config struct
type Config struct {
	URL string
}

//GetConnection - function to retrieve the mongodb connection
func GetConnection() *mgo.Database {
	session, err := mgo.Dial("127.0.0.1")
	CheckError(err)

	session.SetMode(mgo.Monotonic, true)
	database := session.DB("newsapi")
	return database
}

//GetHTTPConnection - function to give an instance of our own special http connection
func GetHTTPConnection() *http.Client {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	return client
}

//CheckError function to check and curtail the error
func CheckError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// func GetLogger() *log {
// 	customLog = new(log)
// 	f, err := os.OpenFile("newsAggregator.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
// 	CheckError(err)

// 	defer f.Close()
// 	customLog.SetOutput(f)
// 	return customLog
// }
