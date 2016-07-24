# jiminy_app
JiMiNy

### GET
```bash
curl -H "Content-Type: application/json" -X GET http://localhost:8888/jiminy/v1/history/list/1

# => 
{
    "header": {
        "response_id": "responseId00001", 
        "result_code": "0"
    },
    "body": {
        "history_list": [
            {
                "number": 1, 
                "send_user": "kamono", 
                "total_like": 5, 
                "url": "3exsRhw3xt8"
            }, 
            {
                "number": 2, 
                "send_user": "nakamura", 
                "total_like": 15, 
                "url": "STg4Ya8bEFo"
            }
        ]
    }
}
```