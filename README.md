# PostgreSQL API in Go

## Introduction
This short program sprang out of a TfL (Transport for London) Hack Day that happened on the 6th of April 2016. TfL has a lot data held in AWS RedShift which is a PostgreSQL(ish) database. The aim of this program was to provide a HTTP API to the database for use in other projects. The program is general enough to apply to any Postgres database.

## Requirements

  * Go ([here](https://golang.org/dl/))
  * The Postgres driver for Go ([here](https://github.com/lib/pq))
  * Mux ([here](https://github.com/gorilla/mux))

**Dependency Installs**:

        go get github.com/gorilla/mux
        go get github.com/lib/pq

## Building
Whilst in the src directory:

    go build pgapi.go
    go run pgapi.go

## Usage

The endpoint defaults to localhost:8080/data/{table} where {table} is the table you want to query and the parameters are as follows:

  * where (optional) - a simple SQL where clause delimited by a colon i.e. ?where=customer:bob
  * limit (optional) - limit the number of rows returned

Using this syntax you can pull back data as JSON in the following schema:

```javascript
{
  "0": {
    "Fields": {
      "id": 1,
      "name": "James"
    }
  },
  "Elapsed": {
    "milliseconds": 12
  }
}
```

Each object is a row with the row number as a string.

## License

The MIT License (MIT)
Copyright (c) 2016 James Milner

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
