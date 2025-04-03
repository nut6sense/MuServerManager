# Channel Manager (Local Backend)

ระบบสำหรับควบคุมการรัน Channel Server แบบ Local

## โครงสร้าง
- config/channels.json: รายชื่อ Channel และคำสั่งรัน
- controllers/: API handler
- process/: จัดการ process จริง (Start / Stop)
- main.go: จุดเริ่ม API Server

## วิธีใช้งาน
1. วางโฟลเดอร์นี้ไว้ข้างๆ `mu-server/`
2. แก้ไข `config/channels.json` ถ้าต้องการเพิ่ม Channel
3. รัน Backend ด้วย:
    go run main.go

4. ใช้ API ผ่าน Postman / curl เช่น:
    POST http://localhost:8080/start/channel1
    POST http://localhost:8080/stop/channel1
    GET  http://localhost:8080/status
    POST http://localhost:8080/start-all
    POST http://localhost:8080/stop-all

## หมายเหตุ
- ใช้ได้กับ Windows (`cmd /C ...`) เป็นหลัก
- รองรับการเพิ่ม Channel ได้ไม่จำกัด