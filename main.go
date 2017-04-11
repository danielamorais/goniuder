package main

import (
	"github.com/dghubble/oauth1"
	"github.com/dghubble/go-twitter/twitter"
	"fmt"
	"net/http"
	"html/template"
)

func main() {
	http.HandleFunc("/login", routers)
	http.Handle("/static/stylesheets/", http.StripPrefix("/static/stylesheets/", http.FileServer(http.Dir("/static/stylesheets"))))
	http.Handle("/static/js/", http.StripPrefix("/static/js/", http.FileServer(http.Dir("http/js"))))

	http.ListenAndServe(":8080", nil)
}

func routers(w http.ResponseWriter, httpRequest *http.Request){
	//w.Header().Set("Content-Type", "application/json")

	if httpRequest.Method == "GET" {
		fmt.Println("Hello world!")
		template, _ := template.ParseFiles("templates/index.html")
		template.Execute(w, httpRequest)
	}

}

func tweet(){
	config := oauth1.NewConfig("consumerKey", "consumerSecret")
	token := oauth1.NewToken("accessToken", "accessSecret")
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	tweet, resp, err := client.Statuses.Update("just setting up my twttr", nil)

	fmt.Println(tweet)
	fmt.Println(resp)
	fmt.Println(err)
}
