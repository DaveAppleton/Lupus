package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// ShowFile - display html file
func ShowFile(w http.ResponseWriter, fileName string) {
	var htmlString string
	htmlBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		htmlString = "File Not Found : " + fileName
	} else {
		htmlString = string(htmlBytes)
	}
	fmt.Fprintf(w, htmlString)
}

func test(w http.ResponseWriter, r *http.Request) {
	ShowFile(w, "index.html")
}

func js(w http.ResponseWriter, r *http.Request) {
	ShowFile(w, "web3.min.js")
}

func main() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./MemoriesOfLupus.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	})
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	go ethFilter()
	ch := make(chan bool)
	http.HandleFunc("/checkTx", checkTx)
	http.HandleFunc("/checkAnswer", checkAnswer)
	http.Handle("/inc/", http.StripPrefix("/inc/", http.FileServer(http.Dir("inc"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.HandleFunc("/", test)

	http.ListenAndServe(":9092", nil)
	<-ch

}
