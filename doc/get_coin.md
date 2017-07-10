# GET Status coin

## GET /coin
ดึงข้อมูลจำนวนเหรียญในตู้ ทั้งหมด

===

### Response

* Status: 200

```json
"result": "success"
"data":
[
  {
    "id": 1,
    "name_coin": "เหรียญ 1",
    "count": 1000,
    "image": "img/c1.png"
  },
  {
    "id": 2,
    "name_coin": "เหรียญ 2",
    "count": 1000,
    "image": "img/c2.png"
  },
  {
    "id": 1,
    "name_coin": "เหรียญ 5",
    "count": 500,
    "image": "img/c5.png"
  },{
    "id": 1,
    "name_coin": "เหรียญ 10",
    "count": 500,
    "image": "img/c10.png"
  }
]
```