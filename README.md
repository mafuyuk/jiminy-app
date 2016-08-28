# jiminy_app
JiMiNy
前準備
===
###MySQL
```
CREATE DATABASE test default character set utf8mb4;
CREATE TABLE `videos` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `video_id` varchar(255) CHARACTER SET utf8 NOT NULL,
  `comment` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL,
  `play_time` varchar(255) CHARACTER SET utf8 NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4
```

履歴情報
===
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
```
### POST
```bash
curl -H "Content-Type: application/json" -X POST http://localhost:61970/jiminy/history

# =>
{"result":201}
```
Youtube動画情報
===
### GET
```bash
curl -H "Content-Type: application/json" -X GET http://localhost:61970/jiminy/video

# =>
[
  {"Id":1,"VideoId":"Su0p6OS2Zyc","Comment":"スキマスイッチの藍を秦基博が歌っている","PlayTime":"3:49"},
  {"Id":2,"VideoId":"Su0p6OS2Zyc","Comment":"スキマスイッチの藍を秦基博が歌っている","PlayTime":"3:49"},
  {"Id":3,"VideoId":"Su0p6OS2Zyc","Comment":"スキマスイッチの藍を秦基博が歌っている","PlayTime":"3:49"},
  {"Id":4,"VideoId":"Su0p6OS2Zyc","Comment":"スキマスイッチの藍を秦基博が歌っている","PlayTime":"3:49"},
  {"Id":5,"VideoId":"Su0p6OS2Zyc","Comment":"スキマスイッチの藍を秦基博が歌っている","PlayTime":"3:49"}
]
```
```bash
curl -H "Content-Type: application/json" -X GET http://localhost:61970/jiminy/video/:id

# =>
[{"Id":2,"VideoId":"Su0p6OS2Zyc","Comment":"スキマスイッチの藍を秦基博が歌っている","PlayTime":"3:49"}]
```
### POST
```bash
curl -H "Content-Type: application/json" -X POST http://localhost:61970/jiminy/video -d '{"video_id":"Su0p6OS2Zyc","comment":"スキマスイッチの藍を秦基博が歌っている","play_time":"3:49"}'
```