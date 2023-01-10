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

### 匿名ユーザ登録
POST  
/api/v1/user/register/
```
curl -X POST http://localhost:8080/api/v1/user/register/
```
Response  
201
```
{
  "timestamp": "2023-01-09T17:48:29+09:00",
  "status_code": 201,
  "message": "Created",
  "errors": [],
  "data": {
    "access_token": "eyJhx (略) lf-Tyo"
  }
}
```
### 誕生日を配列で取得
GET  
/api/v1/birth-days/  
Authorization: Bearer {AccessToken}
```
curl -X GET -H "Authorization: Bearer eyJhbGciO (略) BHrAhDx" http://localhost:8080/api/v1/birth-days/
```
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
      "name": "お父さん",
      "age": 71,
      "date": "1951-12-01T00:00:00Z"
    },
    {
      "id": 2,
      "user_id": 1,
      "name": "お母さん",
      "age": 62,
      "date": "1960-09-23T00:00:00Z"
    }
  ]
}
```
  
401
```
{
  "timestamp": "2023-01-10T01:55:32+09:00",
  "status_code": 401,
  "message": "Unauthorized",
  "errors": [],
  "data": []
}
```
### 誕生日を配列で登録
POST  
/api/v1/birth-days/  
Authorization: Bearer {AccessToken}
```
curl -X POST -H 'Authorization: Bearer eyJhbGciO (略) BHrAhDx' -H "Content-Type: application/json" -d '[{"user_id":1,"name":"お父さん","date":"1951-12-01"}]' http://localhost:8080/api/v1/birth-days/
```
  
Request Body
```json:
[
  {
    "user_id": 1,
    "name": "お父さん",
    "date": "1951-12-01"
  },
  {
    "user_id": 1,
    "name": "お母さん",
    "date": "1960-09-23"
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
      "message": "Validation Error. Property:CreateBirthDayRequestValidates.Elements[1].UserId Value:0",
      "property": "UserId"
    },
    {
      "message": "Validation Error. Property:CreateBirthDayRequestValidates.Elements[1].Date Value:BadDate2025-10Xxx-07",
      "property": "Date"
    }
  ],
  "data": []
}
```
  
401
```
{
  "timestamp": "2023-01-10T01:55:32+09:00",
  "status_code": 401,
  "message": "Unauthorized",
  "errors": [],
  "data": []
}
```
  
