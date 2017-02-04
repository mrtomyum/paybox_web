# Post Sale to Cloud.

This can be run with `silk -silk.url="http://api.nava.work/sell"`

## POST /sell

Post sale transaction to cloud api.

===

* `Status: 200`
* `Content-Type: "application/json;charset=utf-8"`
* Accept: "application/json"

```
{
    "terminal_id": "123",
    "sale_id": 1,
  	"total": 120,
  	"type": "take_home",
  	"sale_pay":{"th20b":0, "th50b":0, "th100b":0,"th500b":0,"th1000b":0, "th1c":0, "th2c":0, "th5c":0, "th10c":2}
  	"sale_subs": [{
  		"line": 1,
  		"item_id": 1,
  		"item_name": "Capuchino Hot",
  		"price_id": 1,
  		"price": 60,
  		"qty": 1,
  		"unit": "แก้ว"
  	}, {
  		"line": 2,
  		"item_id": 2,
  		"item_name": "Capuchino Ice",
  		"price_id": 2,
  		"price": 80,
  		"qty ": 2,
  		"unit": "แก้ว"
  	}]
  }
```

===
* Status:201
* Content-Type: "application/json"
```
{
    "status":200,
    "message":"success",
    "result":{
        ...
    }
}

```