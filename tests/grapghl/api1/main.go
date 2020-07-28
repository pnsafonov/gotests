package main

import (
    "encoding/json"
    "github.com/graphql-go/graphql"
    "gotests/tests/grapghl/api1/schema"
    "log"
    "net/http"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query()
    query := q.Get("query")

    result := graphql.Do(graphql.Params{
        Schema:        schema.Schema,
        RequestString: query,
    })
    json.NewEncoder(w).Encode(result)
}

func main() {
    http.HandleFunc("/", httpHandler)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("ListenAndServe, err = %v\n", err)
    }
}
