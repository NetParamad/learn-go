[English](README.md) | [ไทย](README.th.md)

# Learn Go

A hands-on Go review repo: one folder per topic, from the absolute basics to small runnable CLI tools and an HTTP server — standard library only, no third-party dependencies.

## Overview

This repository is a self-study path through the Go language. Each of the 20 numbered folders is a self-contained `main` package with a short, runnable example and its own `README.md`. You progress in order: language fundamentals (lessons 01–13), then practical standard-library work with files, JSON, and HTTP (14–15), and finally small "real-world" projects (16–20).

The examples deliberately mirror concepts you already know from JavaScript/TypeScript (objects → maps, localStorage → JSON files, `fetch` → `net/http`), so the notes call out the differences inline. All examples use the Go **standard library only** and print to the terminal.

## What you'll learn

### Fundamentals (01–13)

- **01-hello-go** — Entry point is `func main`. Print with `fmt.Println`.
- **02-variables** — `:=` short declare, `var` with an explicit type, and `const` that cannot be reassigned.
- **03-data-types** — Static types and zero values (`0`, `""`, `false`) versus JS's dynamic typing and `undefined`.
- **04-operators** — Arithmetic, comparison, and logical operators. Note that integer `/` truncates and `++` is a statement, not an expression.
- **05-conditions** — `if` requires a boolean (no truthy values), `switch` needs no `break`, plus short-statement `if n := 7; ...`.
- **06-loops** — Only `for` exists: classic three-clause, `for cond` as a while-loop, and `range` over slices.
- **07-functions** — Typed parameters and returns, multiple return values, commonly paired with `error`.
- **08-arrays-slices** — Fixed-size arrays versus growable slices with `append`, slice expressions, and `make`.
- **09-maps-structs** — `map` as the equivalent of JS objects, and `struct` as a typed record; slices of structs.
- **10-methods-interfaces** — Methods bound to a type, and interfaces as method sets (Go-style duck typing).
- **11-pointers** — `*T` for pointers, `&` to take an address, `*` to dereference; mutate a caller's value.
- **12-errors** — `error` as an ordinary return value, checked with `if err != nil`, using `errors.New`.
- **13-goroutines-channels** — `go f()` for concurrency, channels to pass values between goroutines, with `sync.WaitGroup`.

### Stdlib in practice (14–15)

- **14-files-json** — Read/write files with `os`, encode/decode with `encoding/json` (replacing `localStorage`); writes `todos.json`.
- **15-http-client** — Make HTTP requests with `net/http` and decode JSON (replacing `fetch`); fetches a product from `fakestoreapi.com`. **Needs internet.**

### Real-world projects (16–20)

- **16-todo-cli** — Todo CLI with create and list, persisted to a `todos.json` file.
- **17-product-search** — Fetch a product list from an API and filter it client-side by query string. **Needs internet.**
- **18-weather-cli** — Look up current weather for a city using the open-meteo geocoding + forecast APIs (no API key required). **Needs internet.**
- **19-github-profile-viewer** — Fetch and display a GitHub user profile from `api.github.com`. **Needs internet.**
- **20-http-server** — A basic HTTP server with `net/http` (replacing DOM/events), serving a JSON root route and a `/health` route.

## Tech stack

- **Go 1.23+** (module declared `go 1.23` in `go.mod`)
- **Standard library only** — `fmt`, `flag`, `os`, `encoding/json`, `net/http`, `sync`, and others
- No third-party modules, so there is no `go.sum` and no dependency install step

## Project structure

```
learn-go/
├── go.mod                    # module learn-go, go 1.23
├── README.md                 # this file (English)
├── README.th.md              # Thai version
├── .gitignore                # ignores binaries, generated todos.json, .scratch/
├── .gitattributes            # LF normalization; *.go uses eol=lf
├── 01-hello-go/
│   ├── README.md
│   └── main.go
├── 02-variables/
│   └── ...
├── ...
└── 20-http-server/
    ├── README.md
    └── main.go
```

Each `NN-topic/` folder contains exactly two files: a `main.go` (the runnable example) and a bilingual `README.md` note. `.scratch/` is a local-only scratch folder and is git-ignored.

## Getting started

### Prerequisites

- Go 1.23 or newer. Check with:

  ```
  go version
  ```

There is nothing to install: the module has no external dependencies.

### Running an example

Every folder has its own `package main`, so you run one at a time from the repository root:

```
go run ./01-hello-go
go run ./02-variables
```

Or from inside a folder:

```
go run .
```

Some lessons require internet access (15, 17, 18, 19).

## Usage

### Concept lessons (01–15)

```
go run ./03-data-types
go run ./13-goroutines-channels
```

### 16-todo-cli

```
go run ./16-todo-cli -add="read Go"
go run ./16-todo-cli -list
```

Adds todos and lists them; data is stored in `todos.json` in the working directory.

### 17-product-search

```
go run ./17-product-search -q="shirt"
```

Omit `-q` to list all products.

### 18-weather-cli

```
go run ./18-weather-cli -city="Bangkok"
```

Defaults to `Bangkok` if `-city` is omitted.

### 19-github-profile-viewer

```
go run ./19-github-profile-viewer -user="golang"
```

Defaults to `golang` if `-user` is omitted.

### 20-http-server

```
go run ./20-http-server
curl localhost:8080/
curl localhost:8080/health
```

Change the port with the `-port` flag or the `PORT` environment variable (the flag wins):

```
go run ./20-http-server -port=8081
PORT=8082 go run ./20-http-server
```

The root route returns JSON (`{"msg":"..."}`); `/health` returns `ok`.

## Notes

- Examples use `flag` for CLI arguments; run any CLI example with `-h` to see its flags.
- Lessons 14 and 16 write a `todos.json` file at runtime; both are git-ignored.
- Absolute paths from the README examples assume you run from the repository root.
- The lessons intentionally mix Thai comments into the source to accompany the Thai notes.
- Some examples reference an external API (`fakestoreapi.com`, `open-meteo.com`, `api.github.com`); behavior depends on those services being reachable and unchanged.
