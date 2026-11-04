[EN](README.md) | [TH](README.th.md)

<h1 align="center">Learn Go</h1>

<p align="center">
  A hands-on Go review repo: 20 runnable lessons, standard library only.
</p>

<p align="center">
  <a href="#overview">Overview</a> ·
  <a href="#features">Features</a> ·
  <a href="#tech-stack">Tech stack</a> ·
  <a href="#project-structure">Structure</a> ·
  <a href="#getting-started">Getting started</a> ·
  <a href="#environment-variables">Environment</a> ·
  <a href="#usage">Usage</a> ·
  <a href="#notes">Notes</a>
</p>

---

## Overview

This repository is a self-study path through the Go language. Each of the 20 numbered folders is a self-contained `main` package with a short, runnable example and its own bilingual `README.md`. You progress in order: language fundamentals (lessons 01–13), then practical standard-library work with files, JSON, and HTTP (14–15), and finally small "real-world" projects (16–20). The examples deliberately mirror concepts you already know from JavaScript/TypeScript (objects → maps, localStorage → JSON files, `fetch` → `net/http`), so the notes call out the differences inline. All examples use the Go **standard library only** and print to the terminal.

## Features

### Fundamentals (01–13)

- **01-hello-go** — Entry point is `func main`. Print with `fmt.Println`.
- **02-variables** — `:=` short declare, `var` with an explicit type, and `const` that cannot be reassigned.
- **03-data-types** — Static types and zero values (`0`, `""`, `false`) versus JS's dynamic typing and `undefined`.
- **04-operators** — Arithmetic, comparison, and logical operators. Integer `/` truncates, and `++` is a statement, not an expression.
- **05-conditions** — `if` requires a boolean (no truthy values), `switch` needs no `break`, plus short-statement `if n := 7; ...`.
- **06-loops** — Only `for` exists: classic three-clause, `for cond` as a while-loop, and `range` over slices.
- **07-functions** — Typed parameters and returns, multiple return values, commonly paired with `error`.
- **08-arrays-slices** — Fixed-size arrays versus growable slices with `append`, slice expressions, and `make`.
- **09-maps-structs** — `map` as the equivalent of JS objects, and `struct` as a typed record; slices of structs.
- **10-methods-interfaces** — Methods bound to a type, and interfaces as method sets (Go-style duck typing).
- **11-pointers** — `*T` for pointers, `&` to take an address, `*` to dereference; mutate a caller's value.
- **12-errors** — `error` as an ordinary return value, checked with `if err != nil`, using `errors.New`.
- **13-goroutines-channels** — `go f()` for concurrency, buffered channels to pass values between goroutines, with `sync.WaitGroup`.

### Standard library in practice (14–15)

- **14-files-json** — Read/write files with `os`, encode/decode with `encoding/json` (replacing `localStorage`); writes `todos.json`.
- **15-http-client** — Make an HTTP request with `net/http` and decode JSON (replacing `fetch`); fetches a product from `fakestoreapi.com`. **Needs internet.**

### Real-world projects (16–20)

- **16-todo-cli** — Todo CLI with add and list, persisted to a `todos.json` file.
- **17-product-search** — Fetch a product list from an API and filter it client-side by query string. **Needs internet.**
- **18-weather-cli** — Look up current weather for a city using the open-meteo geocoding + forecast APIs (no API key required). **Needs internet.**
- **19-github-profile-viewer** — Fetch and display a GitHub user profile from `api.github.com`. **Needs internet.**
- **20-http-server** — A basic HTTP server with `net/http` (replacing DOM/events), serving a JSON root route and a `/health` route.

## Tech stack

| Area | Choice |
| --- | --- |
| Language | Go 1.23+ (`go.mod` declares `go 1.23`) |
| Dependencies | None — standard library only, so no `go.sum` and no install step |
| CLI arguments | `flag` |
| JSON | `encoding/json` |
| HTTP | `net/http` — client (15, 17–19) and server (20) |
| Concurrency | `goroutines`, `channels`, `sync.WaitGroup` (13) |
| Local data | `todos.json` written at runtime (14, 16); git-ignored |

## Project structure

```
learn-go/
├── go.mod
├── README.md
├── README.th.md
├── .gitignore
├── .gitattributes
├── 01-hello-go/
│   ├── README.md
│   └── main.go
├── 02-variables/
├── 03-data-types/
├── 04-operators/
├── 05-conditions/
├── 06-loops/
├── 07-functions/
├── 08-arrays-slices/
├── 09-maps-structs/
├── 10-methods-interfaces/
├── 11-pointers/
├── 12-errors/
├── 13-goroutines-channels/
├── 14-files-json/
├── 15-http-client/
├── 16-todo-cli/
├── 17-product-search/
├── 18-weather-cli/
├── 19-github-profile-viewer/
└── 20-http-server/
    ├── README.md
    └── main.go
```

Every `NN-topic/` folder contains exactly two files: a `main.go` (the runnable example) and a bilingual `README.md`. `.scratch/` is a local-only scratch folder and is git-ignored.

## Getting started

Requirements: Go 1.23 or newer (check with `go version`). There is nothing to install — the module has no external dependencies. Some lessons (15, 17, 18, 19) require internet access.

Every folder has its own `package main`, so you run one at a time from the repository root:

```bash
git clone https://github.com/NetParamad/learn-go.git
cd learn-go
go run ./01-hello-go
go run ./02-variables
```

Or from inside a folder:

```bash
go run .
```

## Environment variables

| Variable | Required | Notes |
| --- | --- | --- |
| `PORT` | No | Port for `20-http-server`. Defaults to `8080`; the `-port` flag overrides it. |

No other environment variables are read.

## Usage

1. **Run a concept lesson** (01–15) from the repository root:

   ```bash
   go run ./03-data-types
   go run ./13-goroutines-channels
   ```

2. **Add and list todos** with `16-todo-cli`:

   ```bash
   go run ./16-todo-cli -add="read Go"
   go run ./16-todo-cli -list
   ```

   Data is stored in `todos.json` in the working directory.

3. **Search products** with `17-product-search`:

   ```bash
   go run ./17-product-search -q="shirt"
   ```

   Omit `-q` to list all products.

4. **Check the weather** with `18-weather-cli`:

   ```bash
   go run ./18-weather-cli -city="Bangkok"
   ```

   Defaults to `Bangkok` if `-city` is omitted.

5. **View a GitHub profile** with `19-github-profile-viewer`:

   ```bash
   go run ./19-github-profile-viewer -user="golang"
   ```

   Defaults to `golang` if `-user` is omitted.

6. **Start the HTTP server** with `20-http-server`:

   ```bash
   go run ./20-http-server
   curl localhost:8080/
   curl localhost:8080/health
   ```

   Change the port with the `-port` flag or the `PORT` environment variable (the flag wins):

   ```bash
   go run ./20-http-server -port=8081
   PORT=8082 go run ./20-http-server
   ```

   The root route returns JSON (`{"msg":"..."}`); `/health` returns `ok`.

## Notes

- Examples use `flag` for CLI arguments; run any CLI example with `-h` to see its flags.
- Lessons 14 and 16 write a `todos.json` file at runtime; both paths are git-ignored.
- Absolute paths from the README examples assume you run from the repository root.
- The lessons intentionally mix Thai comments into the source to accompany the Thai notes.
- Some examples reference an external API (`fakestoreapi.com`, `open-meteo.com`, `api.github.com`); behavior depends on those services being reachable and unchanged.
- There is no test suite in this repository.
- **License** — MIT (see [LICENSE](LICENSE)).
