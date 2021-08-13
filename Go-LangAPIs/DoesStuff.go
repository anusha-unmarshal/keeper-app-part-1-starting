package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Note struct {
	ID      int64  `json:"key"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// This houses all the notes hardwired


var dsn = "host=localhost user=jawad_notesapp database=mydb password= dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
var db *gorm.DB
var DbEr error

func MakeCon() (*gorm.DB, error) {
	db, DbErrObj := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, DbErrObj
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

func getNotes(w http.ResponseWriter, r *http.Request) {

	var AllNotes = make([]Note, 0)
	res := db.Find(&AllNotes)
	if res.Error != nil {
		log.Fatal("There was an error", res)
	} else {
		marshalledTask, err := json.Marshal(AllNotes)
		if err != nil {
			fmt.Fprint(w, "There was an error", err)
		} else {
			w.Write(marshalledTask)
		}
	}
}

func pageNotFoundHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotFound)
	t := struct {
		Msg string `json:"message"`
	}{Msg: "Error 404, Page not found"}
	msg, _ := json.Marshal(t)
	w.Write(msg)
}

func getNoteByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	var noteItem Note
	result := db.First(&noteItem, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		pageNotFoundHandler(w, r)
	} else {
		marshalledTask, err := json.Marshal(noteItem)
		if err != nil {
			fmt.Fprint(w, "There was an error")
		} else {
			w.Write(marshalledTask)
		}

	}
}

func errHandler(w http.ResponseWriter, err error) {

	ms := fmt.Sprintln(err)
	t := struct {
		Msg string `json:"message"`
	}{Msg: ms}
	w.WriteHeader(http.StatusBadRequest)
	msg, _ := json.Marshal(t)
	w.Write(msg)

}

func addNoteToDB(w http.ResponseWriter, r *http.Request) {
	item, err := decodeAndRetBody(r)
	if err != nil {
		errHandler(w, err)
		return
	}

	db.Select("title", "content").Create(&item)
	w.WriteHeader(http.StatusCreated)
}
func decodeAndRetBody(r *http.Request) (Note, error) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	var item Note
	err := d.Decode(&item)
	return item, err

}

func updateNote(w http.ResponseWriter, r *http.Request) {
	// DB call find by id
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	var noteItem Note
	result := db.First(&noteItem, id)
	statResp := http.StatusOK
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		statResp = http.StatusCreated
	}
	updateItem, err := decodeAndRetBody(r)
	updateItem.ID = id
	if err != nil {
		errHandler(w, err)
		return
	}
	result = db.Save(&updateItem)
	if result.Error != nil {
		log.Println("Error when saving:", result.Error)
		errHandler(w, result.Error)
		return
	}
	w.WriteHeader(statResp)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	var noteItem Note
	result := db.Delete(&noteItem, id)
	if result.RowsAffected  == 0 {
		pageNotFoundHandler(w, r)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)

}

func main() {
	db, DbEr = MakeCon()
	if DbEr != nil {
		log.Fatal("Database failed to connect", DbEr)
	}

	homeRouter := mux.NewRouter()
	homeRouter.HandleFunc("/", homeHandler)

	noteRouter := homeRouter.PathPrefix("/api/").Subrouter()
	noteRouter.Use(jsonWrapper)

	noteRouter.HandleFunc("/task/", getNotes).Methods("GET")     //changed
	noteRouter.HandleFunc("/task/", addNoteToDB).Methods("POST") //changed

	noteRouter.HandleFunc("/task/{id}", getNoteByID).Methods("GET") //changed
	noteRouter.HandleFunc("/task/{id}", updateNote).Methods("PUT")  //changed
	noteRouter.HandleFunc("/task/{id}", deleteNote).Methods("DELETE") //changed

	http.Handle("/", homeRouter)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
