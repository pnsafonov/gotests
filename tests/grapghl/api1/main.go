package main

import (
    "encoding/json"
    "github.com/graphql-go/graphql"
    "gotests/tests/grapghl/api1/schema"
    "log"
    "net/http"
)

type requestBody struct {
    Query string
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
    var query string

    //q := r.URL.Query()
    //query := q.Get("query")

    if r.Method == "GET" {
        q := r.URL.Query()
        query = q.Get("query")
    } else {
        //bytes, err := ioutil.ReadAll(r.Body)
        //if err != nil {
        //    http.Error(w, "Invalid Request",400)
        //    return
        //}

        decoder := json.NewDecoder(r.Body)

        var req requestBody
        err := decoder.Decode(&req)
        if err != nil {
            http.Error(w, "Invalid Request",400)
            return
        }

        query = req.Query
    }

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
