# Hardware API
## 1. API Overview
HARDWARE API เป็น API ใช้เพื่อสำหรับติดต่อกับ Hardware ผ่านทาง Websocket โดยจะมีลักษณะเป็น Request-Response เพื่ออ่าน-เขียนสถานะ และสั่งการอุปกรณ์ต่างๆ และมี Event เป็นสัญญาณแจ้งเตือนการเกิดสถานการณ์ต่างๆบนอุปกรณ์ ใช้ Format ข้อมูลเป็น JSON เพื่อให้ง่ายต่อการพัฒนาและทดสอบ
### 1.1 Flow of communication.
### 1.2 Example of State Machine for client side.
### 1.3) JSON Format

| field | detail |
|--------|----------------------------------------------------------------------------------------------------------|
| device   | ชื่อของอุปกรณ์ที่ต้องการติดต่อด้วย ได้แก่ bill_rcpt, coin_rcpt, coin_hopper, printer โดยมีรายละเอียดดังนี้ bill_acc : เครื่องรับธนบัตร coin_acc : เครื่องรับเหรียญ coin_hopper : เครื่องจ่ายเหรียญ printer : เครื่องพิมพ์ main_board : Main Board                                                               |
| type | ชนิดของ Message |
|request  | ระบุว่า Message นี้เป็นคำสั่ง Request ทิศทางจาก Client ไป Server |
|response | ระบุว่า Message นี้เป็นการตอบผลการ Request ที่ Client ส่งเข้ามาทิศทางจาก Server ไป Client |
| event:  | ระบุว่า Message เป็นการแจ้งเตือนเหตุการณ์ที่เกิดขึ้นในอุปกรณ์ ทิศทางจาก Server ไป Client |
| command | ชื่อของคำสั่ง ซึ่งสามารถใช้ได้แตกต่างกันในอุปกรณ์แต่ละอุปกรณ์ |                                                                |
| data   | ข้อมูลที่คำสั่งต่างๆ ซึ่งจะแตกต่างกันในแต่ละคำสั่ง                                                                     |

## 2.1 Coins Hopper
### 2.1.1 คำสั่งร้องขอ ID ของอุปกรณ์
ใช้สำหรับการร้องขอหมายเลข Serial Number ของ อุปกรณ์ Coins Hopper
#### Command:
    machine_id
#### Request Data: null
#### Request
```
{
	"device": "coin_hopper",
	"type": "request",
	"command": "machine_id",
	"result": null,
	"data": null
}
```
#### Response Data:
ชนิด String
หมายเลข Serial Number ของ Coins Hopper |
#### Response
```
{
	"device": "coin_hopper",
	"type": "response",
	"command": "machine_id",
	"result": < true / false > ,
	"data": "<หมายเลข Serial Number>"
}
```

### 2.1.2