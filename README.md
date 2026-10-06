# Learn Go / เรียน Go

## ภาษาไทย
Repo ทบทวน Go แยกโฟลเดอร์ละเรื่อง เน้นตัวอย่างรันได้สั้นๆ (stdlib เท่านั้น)

ต้องการ: Go 1.23+ (เช็คด้วย `go version`)

- `01-hello-go` ถึง `15-http-client`: รันด้วย `go run ./<บท>` เช่น `go run ./02-variables`
- `16-todo-cli` ถึง `19-github-profile-viewer`: CLI + API รันด้วย `go run . -h` ดู flag
- `20-http-server`: รันด้วย `go run ./20-http-server` แล้ว `curl localhost:8080/` (เปลี่ยนพอร์ตด้วย `-port=8081` หรือ env `PORT`)

เป้าหมาย: `net/http` + แสดงผล ดูที่ `15-http-client` และ `20-http-server`

โปรเจคใช้งานจริง:
- `16-todo-cli` — Todo + ไฟล์ json (แทน localStorage)
- `17-product-search` — ค้นหา + กรองสินค้า (API)
- `18-weather-cli` — เช็คอากาศรายเมือง (open-meteo)
- `19-github-profile-viewer` — ดูโปรไฟล์ GitHub (API)
- `20-http-server` — server พื้นฐาน (แทน DOM)

## English
Go review repo, one folder per topic, short runnable examples (stdlib only).

Requires: Go 1.23+ (check with `go version`)

- `01-hello-go` to `15-http-client`: run with `go run ./<topic>` e.g. `go run ./02-variables`
- `16-todo-cli` to `19-github-profile-viewer`: CLI + API, run with `go run . -h` for flags
- `20-http-server`: run with `go run ./20-http-server` then `curl localhost:8080/` (change port with `-port=8081` or env `PORT`)

Goal: `net/http` + render, see `15-http-client` and `20-http-server`

Real-world projects:
- `16-todo-cli` — Todo + json file (instead of localStorage)
- `17-product-search` — product search + filter (API)
- `18-weather-cli` — city weather (open-meteo)
- `19-github-profile-viewer` — GitHub profile (API)
- `20-http-server` — basic server (instead of DOM)
