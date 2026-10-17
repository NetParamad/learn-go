[English](README.md) | [ไทย](README.th.md)

# Learn Go

รีโปทบทวนภาษา Go แบบลงมือทำจริง แยกโฟลเดอร์ละเรื่อง ตั้งแต่พื้นฐานสุดไปจนถึง CLI เล็ก ๆ ที่รันได้และ HTTP server — ใช้ standard library ล้วน ไม่มี dependency ภายนอก

## ภาพรวม

รีโปนี้คือเส้นทางเรียน Go ด้วยตัวเอง แบ่งเป็น 20 โฟลเดอร์ แต่ละโฟลเดอร์เป็น `main` package ที่สมบูรณ์ในตัวเอง มีตัวอย่างสั้น ๆ ที่รันได้พร้อม `README.md` ของตัวเอง เรียนไล่ตามลำดับได้เลย: พื้นฐานภาษา (บทที่ 01–13) ต่อด้วยการใช้งาน standard library จริงกับไฟล์ JSON และ HTTP (14–15) แล้วปิดท้ายด้วยโปรเจกต์เล็กแบบใช้งานจริง (16–20)

ตัวอย่างตั้งใจเทียบให้เห็นกับสิ่งที่คุ้นเคยจาก JavaScript/TypeScript (object → map, localStorage → ไฟล์ JSON, `fetch` → `net/http`) จึงมีโน้ตชี้ความต่างไว้ในโค้ด ทุกตัวอย่างใช้ **standard library ของ Go เท่านั้น** และพิมพ์ผลออกทางเทอร์มินัล

## สิ่งที่จะได้เรียน

### พื้นฐาน (01–13)

- **01-hello-go** — จุดเริ่มอยู่ที่ `func main` พิมพ์ด้วย `fmt.Println`
- **02-variables** — `:=` ประกาศสั้น, `var` ระบุ type ชัดเจน, และ `const` ที่เปลี่ยนค่าไม่ได้
- **03-data-types** — static type กับ zero value (`0`, `""`, `false`) ต่างจาก dynamic type และ `undefined` ของ JS
- **04-operators** — ตัวดำเนินการคำนวณ เปรียบเทียบ และตรรกะ สังเกตว่า `/` กับ int จะปัดเศษทิ้ง และ `++` เป็น statement ไม่ใช่ expression
- **05-conditions** — `if` ต้องได้ boolean (ไม่มีการตีความเป็นจริง/เท็จ), `switch` ไม่ต้อง `break` และมี `if n := 7; ...` แบบประกาศสั้นในเงื่อนไข
- **06-loops** — มีแค่ `for` เท่านั้น: แบบสามส่วนดั้งเดิม, `for cond` ทำหน้าที่เหมือน while และ `range` วนบน slice
- **07-functions** — ระบุ type ของพารามิเตอร์และค่าที่คืน คืนได้หลายค่า และมักใช้คู่กับ `error`
- **08-arrays-slices** — array ขนาดตายตัว เทียบกับ slice ที่ยืดได้ด้วย `append`, slice expression และ `make`
- **09-maps-structs** — `map` เทียบเท่า object ของ JS และ `struct` คือ record ที่มี type รวมถึง slice ของ struct
- **10-methods-interfaces** — method ผูกกับ type และ interface คือชุดของ method (duck typing สไตล์ Go)
- **11-pointers** — `*T` คือ pointer, `&` เอาที่อยู่, `*` ดึงค่า ใช้แก้ค่าของตัวแปรต้นทางได้
- **12-errors** — `error` เป็นค่าที่คืนตามปกติ เช็คด้วย `if err != nil` และสร้างด้วย `errors.New`
- **13-goroutines-channels** — `go f()` รันแบบ concurrent, channel ใช้ส่งค่าละหว่าง goroutine พร้อม `sync.WaitGroup`

### ลงมือใช้ standard library (14–15)

- **14-files-json** — อ่าน/เขียนไฟล์ด้วย `os` และ encode/decode ด้วย `encoding/json` (แทน `localStorage`) โดยเขียนไฟล์ `todos.json`
- **15-http-client** — ยิง HTTP request ด้วย `net/http` แล้ว decode JSON (แทน `fetch`) ดึงข้อมูลสินค้าจาก `fakestoreapi.com` — **ต้องต่อเน็ต**

### โปรเจกต์ใช้งานจริง (16–20)

- **16-todo-cli** — CLI จัดการ todo เพิ่มและแสดงรายการ เก็บลงไฟล์ `todos.json`
- **17-product-search** — ดึงรายการสินค้าจาก API แล้วกรองฝั่ง client ด้วยคำค้น — **ต้องต่อเน็ต**
- **18-weather-cli** — เช็คอากาศปัจจุบันของเมืองหนึ่ง ผ่าน geocoding + forecast API ของ open-meteo (ไม่ต้องใช้ API key) — **ต้องต่อเน็ต**
- **19-github-profile-viewer** — ดึงและแสดงโปรไฟล์ผู้ใช้ GitHub จาก `api.github.com` — **ต้องต่อเน็ต**
- **20-http-server** — HTTP server พื้นฐานด้วย `net/http` (แทน DOM/events) ให้บริการ route หลักเป็น JSON และ route `/health`

## เทคโนโลยีที่ใช้

- **Go 1.23 ขึ้นไป** (ใน `go.mod` ประกาศ `go 1.23`)
- **ใช้แค่ standard library** — `fmt`, `flag`, `os`, `encoding/json`, `net/http`, `sync` และอื่น ๆ
- ไม่มีโมดูลภายนอก จึงไม่มี `go.sum` และไม่มีขั้นตอนติดตั้ง dependency

## โครงสร้างโปรเจกต์

```
learn-go/
├── go.mod                    # module learn-go, go 1.23
├── README.md                 # ไฟล์นี้ (อังกฤษ)
├── README.th.md              # เวอร์ชันภาษาไทย
├── .gitignore                # ignore binary, todos.json ที่ถูกสร้าง, .scratch/
├── .gitattributes            # normalize เป็น LF; *.go ใช้ eol=lf
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

ทุกโฟลเดอร์ `NN-topic/` มีสองไฟล์พอดี: `main.go` (ตัวอย่างที่รันได้) และ `README.md` โน้ตสองภาษา ส่วน `.scratch/` เป็นโฟลเดอร์ทดลองในเครื่องเท่านั้นและถูก ignore ใน git

## เริ่มต้นใช้งาน

### สิ่งที่ต้องมีก่อน

- Go 1.23 หรือใหม่กว่า ตรวจสอบด้วย:

  ```
  go version
  ```

ไม่ต้องติดตั้งอะไรเพิ่ม เพราะโมดูลไม่มี dependency ภายนอก

### การรันตัวอย่าง

ทุกโฟลเดอร์เป็น `package main` ของตัวเอง จึงรันทีละอันจากโฟลเดอร์หลักของรีโป:

```
go run ./01-hello-go
go run ./02-variables
```

หรือรันจากในโฟลเดอร์นั้น:

```
go run .
```

บางบทต้องต่อเน็ต (15, 17, 18, 19)

## วิธีใช้งาน

### บทเรียนพื้นฐาน (01–15)

```
go run ./03-data-types
go run ./13-goroutines-channels
```

### 16-todo-cli

```
go run ./16-todo-cli -add="อ่าน Go"
go run ./16-todo-cli -list
```

เพิ่ม todo แล้วแสดงรายการ ข้อมูลเก็บใน `todos.json` ที่โฟลเดอร์ปัจจุบัน

### 17-product-search

```
go run ./17-product-search -q="shirt"
```

ถ้าไม่ใส่ `-q` จะแสดงสินค้าทั้งหมด

### 18-weather-cli

```
go run ./18-weather-cli -city="Bangkok"
```

ถ้าไม่ใส่ `-city` ค่าเริ่มต้นคือ `Bangkok`

### 19-github-profile-viewer

```
go run ./19-github-profile-viewer -user="golang"
```

ถ้าไม่ใส่ `-user` ค่าเริ่มต้นคือ `golang`

### 20-http-server

```
go run ./20-http-server
curl localhost:8080/
curl localhost:8080/health
```

เปลี่ยนพอร์ตได้ด้วย flag `-port` หรือ environment variable `PORT` (ถ้าใส่ทั้งคู่ flag จะชนะ):

```
go run ./20-http-server -port=8081
PORT=8082 go run ./20-http-server
```

route หลักคืนค่า JSON (`{"msg":"..."}`) ส่วน `/health` คืนค่า `ok`

## หมายเหตุ

- ตัวอย่างรับ argument ผ่าน `flag` ถ้าอยากดู flag ทั้งหมดให้รันด้วย `-h`
- บทที่ 14 และ 16 เขียนไฟล์ `todos.json` ตอนรัน ทั้งคู่ถูก ignore ใน git
- ตัวอย่าง path ใน README นี้ถือว่ารันจากโฟลเดอร์หลักของรีโป
- โค้ดตัวอย่างจงใจมีคอมเมนต์ภาษาไทยปนอยู่ คู่กับโน้ตภาษาไทย
- บางตัวอย่างเรียก API ภายนอก (`fakestoreapi.com`, `open-meteo.com`, `api.github.com`) ผลลัพธ์จึงขึ้นอยู่กับว่าบริการเหล่านั้นเข้าถึงได้และไม่เปลี่ยนแปลง
