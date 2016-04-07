package main

import (
    "fmt"
    "os"
    "encoding/json"
    "io"
    "log"
    "net/http"
    "database/sql"
    "strings"
    "time"
    "strconv"
    "github.com/lib/pq"
)

func handler(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/json")

    // Timing
    start := time.Now()

    // Postgres Credentials
    // Get them from environment variables
    var (
        DB_HOST     = os.Getenv("PG_HOST")
        DB_NAME     = os.Getenv("PG_NAME")
        DB_USER     = os.Getenv("PG_USER")
        DB_PASSWORD = os.Getenv("PG_PASSWORD")
        DB_PORT     = os.Getenv("PG_PORT")
    )


    // Postgres Connect
    dbinfo := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%s sslmode=disable",
                           DB_HOST, DB_NAME, DB_USER, DB_PASSWORD, DB_PORT)
    db, err := sql.Open("postgres", dbinfo)
    if err != nil {
        handleError(w, err.Error())
    }
    defer db.Close()

    table := r.FormValue("table")
    limit := r.FormValue("limit")
    where := r.FormValue("where")

    // If table defined
    if table != "" {
        var q string

        if limit == "" {
          limit = "100"
        }

        table := pq.QuoteIdentifier(table)
        if where != "" {
          clause := strings.Split(where, ":") // We delimit where clause using :
          ident := clause[0]
          equals := clause[1]
          q = fmt.Sprintf("SELECT * FROM %s WHERE %s = %s LIMIT %s", table, ident, equals, limit)
        } else {
          q = fmt.Sprintf("SELECT * FROM %s LIMIT %s", table, limit)
        }

        log.Print(q) // Log the query to console

        // Get rows or error
        rows, err := db.Query(q) // Pass the query to the database
        if err != nil {
            handleError(w, err.Error())
            return
        }
        defer rows.Close()

        // Setup variables for
        columns, _ := rows.Columns()
        colLen := len(columns)
        values := make([]interface{}, colLen)
        valuePtrs := make([]interface{}, colLen)
        returnjson := make(map[string]interface{})
        rowNum := 0

        // Struct for our rows
        type Row struct {
           Fields map[string]interface{}
        }

        for rows.Next() {

            for i, _ := range columns {
                valuePtrs[i] = &values[i]
            }

            rows.Scan(valuePtrs...)

            rowStruct := Row{ Fields : make(map[string]interface{}) }

            for i, col := range columns {

                var v interface{}

                val := values[i]

                b, ok := val.([]byte)

                if (ok) {
                    v = string(b) // Turn bytes into string
                } else {
                    v = val // All other values
                }

                rowStruct.Fields[col] = v

                //fmt.Println(col, v)
            }
            returnjson[strconv.Itoa(rowNum)] = rowStruct
            rowNum += 1
        }

        // Add elapsed time row
        elapsed := make(map[string]interface{})
        elapsed["milliseconds"] = time.Since(start)/1000000
        returnjson["Elapsed"] = elapsed

        // Return the data
        json.NewEncoder(w).Encode(returnjson)

    }
}

func handleError(w http.ResponseWriter, err string) {
    type APIError struct {
        Error string
    }
    re, _ := json.Marshal(APIError{Error: err})
    io.WriteString(w, string(re))
}

func main() {
    http.HandleFunc("/data", handler)
    http.ListenAndServe(":8080", nil)
}
