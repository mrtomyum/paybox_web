# POST add coin 

## POST /coin
ส่งข้อมูลการเพิ่มเหรียญใน ตู้

===

### Request
```json
    {
        "id": 1,
        "count": 1000
      },
      {
        "id": 2,
        "count": 1000
      },
      {
        "id": 1,
        "count": 500,
      },{
        "id": 1,
        "count": 500,
      }
```
* Content-Type: "application/json"

===

### Response

* Status: 200
* Content-Type: "application/json; charset=utf-8"

```json
    {
        "result": "success",
        "massage": "add coin completed"
    } 
```