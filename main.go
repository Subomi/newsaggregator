package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/LordRahl90/newsAPIAggregator/newcontent"
	"github.com/LordRahl90/newsAPIAggregator/utility"
)

func main() {
	f, err := os.OpenFile("logs/newsAggregator.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	utility.CheckError(err)
	defer f.Close()
	log.SetOutput(f)
	rand.Seed(time.Now().Unix())

	newcontent.GetAllContents()

	var input string
	fmt.Scanln(&input)

}
