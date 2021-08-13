package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)


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
	// w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusNotFound)
	t := struct {
		Msg string `json:"message"`
	}{Msg: "Error 404, Page not found"}
	msg, _ := json.Marshal(t)
	w.Write(msg)
}
func fetchTaskById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

	for _, noteItem := range noteList {
		if noteItem.Key == id {

			marshalledTask, err := json.Marshal(noteItem)
			if err != nil {
				fmt.Fprint(w, "There was an error")
			} else {
				w.Write(marshalledTask)
			}

			return
		}

	}
	pageNotFoundHandler(w, r)

}
func errHandler(w http.ResponseWriter, err error) {
	// w.Header().Add("Content-Type", "text/plain")
	ms := fmt.Sprintln(err)
	t := struct {
		Msg string `json:"message"`
	}{Msg: ms}
	w.WriteHeader(http.StatusBadRequest)
	msg, _ := json.Marshal(t)
	w.Write(msg)

}

func addTask(w http.ResponseWriter, r *http.Request) {
	item, err := decodeAndRetBody(r)
	if err != nil {
		errHandler(w, err)
		return
	}
	noteList = append(noteList, item)
	w.WriteHeader(http.StatusCreated)
	// fmt.Println(item)

}
func decodeAndRetBody(r *http.Request) (Note, error) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	var item Note
	err := d.Decode(&item)
	return item, err

}
func modifyTask(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	updateItem, err := decodeAndRetBody(r)
	if err != nil {
		errHandler(w, err)
		return
	}
	for index, noteItem := range noteList {
		if noteItem.Key == id {
			if updateItem.Title != "" {
				noteList[index].Title = updateItem.Title
			}
			if updateItem.Content != "" {
				noteList[index].Content = updateItem.Content
			}
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	noteList = append(noteList, updateItem)
	w.WriteHeader(http.StatusCreated)
}

func removeNoteFromList(i int) {
	noteList[i] = noteList[len(noteList)-1]
	noteList = noteList[:len(noteList)-1]
}

func removeTask(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	for index, noteItem := range noteList {
		if noteItem.Key == id {
			removeNoteFromList(index)
			w.WriteHeader(http.StatusNoContent)
		}
	}
	pageNotFoundHandler(w, r)

}

func main() {
	homeRouter := mux.NewRouter()
	homeRouter.HandleFunc("/", homeHandler)

	noteRouter := homeRouter.PathPrefix("/api/").Subrouter()
	noteRouter.Use(jsonWrapper)

	noteRouter.HandleFunc("/task/", fetchAllTasks).Methods("GET")
	noteRouter.HandleFunc("/task/", addTask).Methods("POST")

	noteRouter.HandleFunc("/task/{id}", fetchTaskById).Methods("GET")
	noteRouter.HandleFunc("/task/{id}", modifyTask).Methods("PUT")
	noteRouter.HandleFunc("/task/{id}", removeTask).Methods("DELETE")

	http.Handle("/", homeRouter)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
