package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// TO-DO: Notes[post], Notes/{id}[put,delete],
// import
// import "github.com/gorilla/mux"

type Note struct {
	Key     int64  `json:"key"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// This houses all the notes hardwired
var noteList = []Note{
	{
		Key:     1,
		Title:   "Delegation",
		Content: "Q. How many programmers does it take to change a light bulb? A. None – It’s a hardware problem",
	},
	{
		Key:     2,
		Title:   "Loops",
		Content: "How to keep a programmer in the shower forever. Show him the shampoo bottle instructions: Lather. Rinse. Repeat.",
	},
	{
		Key:     3,
		Title:   "Arrays",
		Content: "Q. Why did the programmer quit his job? A. Because he didn't get arrays.",
	},
	{
		Key:     4,
		Title:   "Hardware vs. Software",
		Content: "What's the difference between hardware and software? You can hit your hardware with a hammer, but you can only curse at your software.",
	},
}

func jsonWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		// fmt.Fprintf(w, "hello there")
		next.ServeHTTP(w, r)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You're on the homepage")

}
func fetchAllTasks(w http.ResponseWriter, r *http.Request) {
	marshalledTask, err := json.Marshal(noteList)
	if err != nil {
		fmt.Fprint(w, "There was an error")
	} else {
		w.Write(marshalledTask)
	}

}
func pageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Error 404, Page not found")

}
func fetchTaskById(w http.ResponseWriter, r *http.Request) {
	id,_ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	// var Stat bool = false

	for _, noteItem := range noteList {
		if noteItem.Key == id {


			marshalledTask, err := json.Marshal(noteItem)
			if err != nil {
				fmt.Fprint(w, "There was an error")
			} else {
				w.Write(marshalledTask)
			}
			// w.Header().Add("")

			return
		}

	}
	pageNotFoundHandler(w,r)
	


}

func main() {
	homeRouter := mux.NewRouter()
	homeRouter.HandleFunc("/", homeHandler)

	noteRouter := homeRouter.PathPrefix("/api/").Subrouter()
	noteRouter.Use(jsonWrapper)

	noteRouter.HandleFunc("/task/", fetchAllTasks).Methods("GET")
	noteRouter.HandleFunc("/task/{id}", fetchTaskById).Methods("GET")
	http.Handle("/", homeRouter)
	// http.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))

	// for _, item := range noteList {
	// 	fmt.Println(item)
	// }
}