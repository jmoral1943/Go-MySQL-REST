package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoral1943/musicapi/app"
)

func main() {

	// request.HandleRequests()

	a := app.App{}
	a.Initialize(os.Getenv("MusicAPIDB"))

	a.Run(":8080")
}

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the HomePage!")
// 	fmt.Println("Endpoint Hit: homePage")
// }

// func returnAllArticles(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: returnAllArticles")

// 	json.NewEncoder(w).Encode(Songs)
// }

// func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key := vars["id"]

// 	for _, article := range Songs {
// 		if article.ID == key {
// 			json.NewEncoder(w).Encode(article)
// 		}
// 	}
// }

// func createNewArticle(w http.ResponseWriter, r *http.Request) {
// 	// get the body of our POST request
// 	// return the string response containing the request body
// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	var song Song
// 	json.Unmarshal(reqBody, &song)
// 	// update our global Articles array to include
// 	// our new Article
// 	Songs = append(Song, article)

// 	json.NewEncoder(w).Encode(article)
// }

// func deleteArticle(w http.ResponseWriter, r *http.Request) {
// 	// once again, we will need to parse the path parameters
// 	vars := mux.Vars(r)
// 	// we will need to extract the `id` of the article we
// 	// wish to delete
// 	id := vars["id"]

// 	// we then need to loop through all our articles
// 	for index, article := range Articles {
// 		// if our id path parameter matches one of our
// 		// articles
// 		if article.ID == id {
// 			// updates our Articles array to remove the
// 			// article
// 			Articles = append(Articles[:index], Articles[index+1:]...)
// 		}
// 	}

// }

// func updateArticle(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	id := vars["id"]

// 	// return the string response containing the request body
// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	var article Article
// 	json.Unmarshal(reqBody, &article)

// 	for i := 0; i < len(Articles); i++ {
// 		a := &Articles[i]

// 		if Articles[i].ID == id {
// 			fmt.Println(&Articles[i].ID)
// 			a.Title = article.Title
// 			a.Desc = article.Desc
// 			a.Content = article.Content
// 			json.NewEncoder(w).Encode(Articles[i])
// 		}
// 	}
// }

// func main() {
// 	Articles = []Article{
// 		Article{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
// 		Article{ID: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
// 	}
// 	request.HandleRequests()

// }
