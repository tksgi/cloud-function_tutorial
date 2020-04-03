package helloworld

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"time"
)

type FirestoreValue struct {
	CreateTime time.Time
	Name       string
	Msg        string
	UpdateTime time.Time
}

func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Hello, World!")
		return
	}
	var msg = "Hello, "
	if d.Name == "" {
		msg += "World!"
	} else {
		msg += html.EscapeString(d.Name) + "!"
	}
	data := &FirestoreValue{
		CreateTime: time.Now(),
		Name:       d.Name,
		Msg:        msg,
		UpdateTime: time.Now(),
	}

	_, _, err := client.Collection("results").Add(ctx, data)
	if err != nil {
		fmt.Fprintf(w, "Failed to create record: %v", err)
	}
	fmt.Fprintf(w, msg)
}
