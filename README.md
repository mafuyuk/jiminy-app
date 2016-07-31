# jiminy_app
JiMiNy

### GET
```bash
curl -H "Content-Type: application/json" -X GET http://localhost:61970/jiminy/history/1

# => 
[
  {
    "number":1,
    "url":"3exsRhw3xt8",
    "send_user":"kamono",
    "total_like":5
  },
  {
    "number":2,
    "url":"STg4Ya8bEFo",
    "send_user":"nakamura",
    "total_like":15
  }
]

### POST
```bash
curl -H "Content-Type: application/json" -X POST http://localhost:61970/jiminy/history

# =>
{"result":201}
```