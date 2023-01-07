# birthday-reminder
 In order not to forget the birthdays of important people.

## Features
 Developing based on Clean Architecture
## Usage
 Download Project.
```bash:
$ git clone https://github.com/yuukanehiro/birthday-reminder.git
```
 Run Docker.
```bash:
$ cd birthday-reminder
$ docker-compose build --no-cache
$ docker-compose up -d
```

## API

### 誕生日を配列で取得
GET  
http://localhost:8080/api/v1/birth-days/
  
Response  
200
```json:
{
  "timestamp": "2023-01-06T19:21:44+09:00",
  "status_code": 200,
  "message": "OK",
  "errors": [],
  "data": [
    {
      "id": 1,
      "user_id": 1,
      "date": "2018-10-20T00:00:00Z"
    },
    {
      "id": 2,
      "user_id": 1,
      "date": "2018-10-21T00:00:00Z"
    }
  ]
}
```

### 誕生日を登録
POST  
http://localhost:8080/api/v1/birth-days/  
  
RequestBody
```json:
[
  {
    "user_id": 1,
    "date": "2022-01-04"
  },
  {
    "user_id": 1,
    "date": "2022-01-05"
  }
]
```
Response  
201
```json:
{
  "timestamp": "2023-01-06T22:58:37+09:00",
  "status_code": 201,
  "message": "Created",
  "errors": [],
  "data": []
}
```
400
```
{
  "timestamp": "2023-01-07T21:11:39+09:00",
  "status_code": 400,
  "message": "Bad Request",
  "errors": [
    {
      "Message": "Validation Error. Property:UserId Value:0",
      "Property": "UserId"
    },
    {
      "Message": "Validation Error. Property:Date Value:BadDate2025-10Xxx-07",
      "Property": "Date"
    }
  ],
  "data": []
}
```