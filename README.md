# go-tmdb

[GoDoc](https://godoc.org/github.com/ryanbradynd05/go-tmdb)  --  [Code Coverage Report](http://rawgit.com/ryanbradynd05/go-tmdb/master/coverage/coverage_report.html)
=================================

A Go Wrapper for the API of [The Movie DB](http://www.themoviedb.org/). Complete [documentation](https://godoc.org/github.com/ryanbradynd05/go-tmdb) and [test suite](http://rawgit.com/ryanbradynd05/go-tmdb/master/coverage/coverage_report.html) are included.

An **api_key** is needed to use the API. Register for one at [themoviedb.org](https://www.themoviedb.org/documentation/api).

Note: This product uses the TMDb API but is not endorsed or certified by TMDb.

<img src="https://assets.tmdb.org/images/logos/var_7_0_tmdb-logo-2_Antitled.svg" alt="The Movie DB" width="200" height="200" />

## How to install

```shell
go get github.com/ryanbradynd05/go-tmdb
```

## How to use

Import the library

```go
import "github.com/ryanbradynd05/go-tmdb"
```
    
Create a container struct for global properties
```go
type TMDb struct {
  apiKey string
}
```

Initialize the library using your api_key
```go
func Init(apiKey string) *TMDb {
  return &TMDb{apiKey: apiKey}
}
```

Use the api methods as you want, for example:

```go
fightClubInfo, err := TMDb.GetMovieInfo(550, nil)
```

To use optional parameters, pass in a map[string]string of options and values:

```go
var options = make(map[string]string)
options["language"] = "es"
spanishFightClub, err := TMDb.GetMovieInfo(550, options)
```

All functions return Go structs. To return JSON, use the ToJSON function:

```go
fightClubJson, err := TMDb.ToJSON(fightClubInfo)
```

## How to test

Create a local.yml file in the root directory that mirrors the local.yml.example file. Then, either run go test to simply run the tests or run the coverage.sh file to run the tests with coverage info.

## Available methods

All themoviedb.org API v3 GET methods are included. The POST and DELETE APIs are not included yet. For examples on how to call each function, refer to that function's tests. For documentation of the TheMovieDB's API, see their [documentation](http://docs.themoviedb.apiary.io/).

## License 

The MIT License (MIT)

Copyright (c) 2015 Ryan Brady

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
