[EN](README.md) | [TH](README.th.md)

<h1 align="center">Learn Go</h1>

<p align="center">
  รีโปทบทวนภาษา Go แบบลงมือทำ: 20 บทเรียนที่รันได้ ใช้ standard library ล้วน
</p>

<p align="center">
  <a href="#ภาพรวม">ภาพรวม</a> ·
  <a href="#ฟีเจอร์">ฟีเจอร์</a> ·
  <a href="#เทคโนโลยีที่ใช้">เทคโนโลยี</a> ·
  <a href="#โครงสร้างโปรเจกต์">โครงสร้าง</a> ·
  <a href="#เริ่มต้นใช้งาน">เริ่มต้นใช้งาน</a> ·
  <a href="#ตัวแปรสภาพแวดล้อม">ตัวแปรสภาพแวดล้อม</a> ·
  <a href="#วิธีใช้งาน">วิธีใช้งาน</a> ·
  <a href="#หมายเหตุ">หมายเหตุ</a>
</p>

---

## ภาพรวม

รีโปนี้คือเส้นทางเรียน Go ด้วยตัวเอง แบ่งเป็น 20 โฟลเดอร์ แต่ละโฟลเดอร์เป็น `main` package ที่สมบูรณ์ในตัวเอง มีตัวอย่างสั้น ๆ ที่รันได้พร้อม `README.md` สองภาษา เรียนไล่ตามลำดับได้เลย: พื้นฐานภาษา (บทที่ 01–13) ต่อด้วยการใช้งาน standard library จริงกับไฟล์ JSON และ HTTP (14–15) แล้วปิดท้ายด้วยโปรเจกต์เล็กแบบใช้งานจริง (16–20) ตัวอย่างตั้งใจเทียบให้เห็นกับสิ่งที่คุ้นเคยจาก JavaScript/TypeScript (object → map, localStorage → ไฟล์ JSON, `fetch` → `net/http`) จึงมีโน้ตชี้ความต่างไว้ในโค้ด ทุกตัวอย่างใช้ **standard library ของ Go เท่านั้น** และพิมพ์ผลออกทางเทอร์มินัล

## ฟีเจอร์

### พื้นฐาน (01–13)

- **01-hello-go** — จุดเริ่มอยู่ที่ `func main` พิมพ์ด้วย `fmt.Println`
- **02-variables** — `:=` ประกาศสั้น, `var` ระบุ type ชัดเจน, และ `const` ที่เปลี่ยนค่าไม่ได้
- **03-data-types** — static type กับ zero value (`0`, `""`, `false`) ต่างจาก dynamic type และ `undefined` ของ JS
- **04-operators** — ตัวดำเนินการคำนวณ เปรียบเทียบ และตรรกะ `/` กับ int จะปัดเศษทิ้ง และ `++` เป็น statement ไม่ใช่ expression
- **05-conditions** — `if` ต้องได้ boolean (ไม่มีการตีความเป็นจริง/เท็จ), `switch` ไม่ต้อง `break` และมี `if n := 7; ...` แบบประกาศสั้นในเงื่อนไข
- **06-loops** — มีแค่ `for` เท่านั้น: แบบสามส่วนดั้งเดิม, `for cond` ทำหน้าที่เหมือน while และ `range` วนบน slice
- **07-functions** — ระบุ type ของพารามิเตอร์และค่าที่คืน คืนได้หลายค่า และมักใช้คู่กับ `error`
- **08-arrays-slices** — array ขนาดตายตัว เทียบกับ slice ที่ยืดได้ด้วย `append`, slice expression และ `make`
- **09-maps-structs** — `map` เทียบเท่า object ของ JS และ `struct` คือ record ที่มี type รวมถึง slice ของ struct
- **10-methods-interfaces** — method ผูกกับ type และ interface คือชุดของ method (duck typing สไตล์ Go)
- **11-pointers** — `*T` คือ pointer, `&` เอาที่อยู่, `*` ดึงค่า ใช้แก้ค่าของตัวแปรต้นทางได้
- **12-errors** — `error` เป็นค่าที่คืนตามปกติ เช็คด้วย `if err != nil` และสร้างด้วย `errors.New`
- **13-goroutines-channels** — `go f()` รันแบบ concurrent, buffered channel ใช้ส่งค่าละหว่าง goroutine พร้อม `sync.WaitGroup`

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

| ด้าน | ที่เลือกใช้ |
| --- | --- |
| ภาษา | Go 1.23+ (`go.mod` ประกาศ `go 1.23`) |
| Dependency | ไม่มี — ใช้ standard library ล้วน จึงไม่มี `go.sum` และไม่มีขั้นตอนติดตั้ง |
| Argument ของ CLI | `flag` |
| JSON | `encoding/json` |
| HTTP | `net/http` — client (15, 17–19) และ server (20) |
| Concurrency | `goroutines`, `channels`, `sync.WaitGroup` (13) |
| ข้อมูลในเครื่อง | `todos.json` ที่เขียนตอนรัน (14, 16) และถูก ignore ใน git |

## โครงสร้างโปรเจกต์

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

ทุกโฟลเดอร์ `NN-topic/` มีสองไฟล์พอดี: `main.go` (ตัวอย่างที่รันได้) และ `README.md` สองภาษา ส่วน `.scratch/` เป็นโฟลเดอร์ทดลองในเครื่องเท่านั้นและถูก ignore ใน git

## เริ่มต้นใช้งาน

ความต้องการ: Go 1.23 หรือใหม่กว่า (ตรวจด้วย `go version`) ไม่ต้องติดตั้งอะไรเพิ่ม เพราะโมดูลไม่มี dependency ภายนอก บางบท (15, 17, 18, 19) ต้องต่อเน็ต

ทุกโฟลเดอร์เป็น `package main` ของตัวเอง จึงรันทีละอันจากโฟลเดอร์หลักของรีโป:

```bash
git clone https://github.com/NetParamad/learn-go.git
cd learn-go
go run ./01-hello-go
go run ./02-variables
```

หรือรันจากในโฟลเดอร์นั้น:

```bash
go run .
```

## ตัวแปรสภาพแวดล้อม

| ตัวแปร | จำเป็น | หมายเหตุ |
| --- | --- | --- |
| `PORT` | ไม่ | พอร์ตของ `20-http-server` ค่าเริ่มต้นคือ `8080` และ flag `-port` จะเขียนทับค่านี้ |

ไม่มีตัวแปรสภาพแวดล้อมอื่นที่โค้ดอ่าน

## วิธีใช้งาน

1. **รันบทเรียนพื้นฐาน** (01–15) จากโฟลเดอร์หลักของรีโป:

   ```bash
   go run ./03-data-types
   go run ./13-goroutines-channels
   ```

2. **เพิ่มและแสดง todo** ด้วย `16-todo-cli`:

   ```bash
   go run ./16-todo-cli -add="read Go"
   go run ./16-todo-cli -list
   ```

   ข้อมูลเก็บใน `todos.json` ที่โฟลเดอร์ปัจจุบัน

3. **ค้นหาสินค้า** ด้วย `17-product-search`:

   ```bash
   go run ./17-product-search -q="shirt"
   ```

   ถ้าไม่ใส่ `-q` จะแสดงสินค้าทั้งหมด

4. **เช็คอากาศ** ด้วย `18-weather-cli`:

   ```bash
   go run ./18-weather-cli -city="Bangkok"
   ```

   ถ้าไม่ใส่ `-city` ค่าเริ่มต้นคือ `Bangkok`

5. **ดูโปรไฟล์ GitHub** ด้วย `19-github-profile-viewer`:

   ```bash
   go run ./19-github-profile-viewer -user="golang"
   ```

   ถ้าไม่ใส่ `-user` ค่าเริ่มต้นคือ `golang`

6. **เปิด HTTP server** ด้วย `20-http-server`:

   ```bash
   go run ./20-http-server
   curl localhost:8080/
   curl localhost:8080/health
   ```

   เปลี่ยนพอร์ตได้ด้วย flag `-port` หรือ environment variable `PORT` (ถ้าใส่ทั้งคู่ flag จะชนะ):

   ```bash
   go run ./20-http-server -port=8081
   PORT=8082 go run ./20-http-server
   ```

   route หลักคืนค่า JSON (`{"msg":"..."}`) ส่วน `/health` คืนค่า `ok`

## หมายเหตุ

- ตัวอย่างรับ argument ผ่าน `flag` ถ้าอยากดู flag ทั้งหมดให้รันด้วย `-h`
- บทที่ 14 และ 16 เขียนไฟล์ `todos.json` ตอนรัน ทั้งสอง path ถูก ignore ใน git
- ตัวอย่าง path ใน README นี้ถือว่ารันจากโฟลเดอร์หลักของรีโป
- โค้ดตัวอย่างจงใจมีคอมเมนต์ภาษาไทยปนอยู่ คู่กับโน้ตภาษาไทย
- บางตัวอย่างเรียก API ภายนอก (`fakestoreapi.com`, `open-meteo.com`, `api.github.com`) ผลลัพธ์จึงขึ้นอยู่กับว่าบริการเหล่านั้นเข้าถึงได้และไม่เปลี่ยนแปลง
- รีโปนี้ไม่มีชุดเทสต์
- **สัญญาอนุญาต** — MIT (ดู [LICENSE](LICENSE))
