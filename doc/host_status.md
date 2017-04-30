# Host Status
ตรวจสอบสถานะอุปกรณ์ เป็น Goroutine ทุกๆ 5 วินาที แล้วส่ง msg ให้กับ UI ทันทีที่ตรวจพบความเปลี่ยนแปลง

## WEB SOCKET request
```
{
	"command": "host_status",
	"type": "request"
}
```
---
```
{
	"command": "host_status",
	"type": "response",
	"data": {
		"bill_acc_online": true,
		"bill_acc_status": "stop",
		"coin_acc_online": true,
		"coin_acc_status": "stop",
		"coin_hopper_online": true,
		"coin_hopper_c1": 0,
		"coin_hopper_c2": 100,
		"coin_hopper_c5": 100,
		"coin_hopper_c10": 500,
		"printer_online": true,
		"printer_status": "paper_out",
		"gsm_online": true,
		"server_online": false,
		"door1_status": "close",
		"door2_status": "close"
	}
}
```


ตัวอย่าง
=====
## Event
จะส่งทันทีผ่าน Web Socket ส่งรายชื่ออุปกรณ์ และสถานะปัจจุบัน เพื่อส่งให้ UI ทำการแสดงผลแบบ Real Time เช่น เงินหมดเน็ตหลุด "3g_status": "offline" เป็นต้น
```
{
	"command": "host_status",
	"type": "event",
	"data": {
		"bill_acc_online": true,
		"bill_acc_status": "stop",
		"coin_acc_online": true,
		"coin_acc_status": "stop",
		"coin_hopper_online": true,
		"coin_hopper_c1": 0,
		"coin_hopper_c2": 100,
		"coin_hopper_c5": 100,
		"coin_hopper_c10": 500,
		"printer_online": true,
		"printer_status": "paper_out",
		"gsm_online": true,
		"server_online": false,
		"door1_status": "close",
		"door2_status": "close"
	}
}
```

