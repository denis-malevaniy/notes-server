package main

//заметки хранятся в ОП(потом заменить на БД)

import (
	"fmt"
	"net/http"
	"encoding/json"
	"sync"
)

type Note struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
}

var (
	notes []Note
	mu    sync.Mutex
	nextID = 1
)

func addNote(title, content string){
	mu.Lock()
	defer mu.Unlock()
	note := Note{ID: nextID, Title: title, Content: content}
	nextID++ 
	notes = append(notes, note)
}

func notesHandler (w http.ResponseWriter, r * http.Request){
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json,json.NewEncoder(w).Encode(notes)
}

func main() {
	addNote("FIRST", "UOUOUOUOUOUUOOU GO GO")

	http.HandleFunc("/notes", notesHandler)
	fmt.Println("started at http://localhost:9000")
	http.ListenAndServe(":9000", nil)
}
