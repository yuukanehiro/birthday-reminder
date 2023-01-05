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
```json:
[
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
```

### 誕生日を登録
POST  
http://localhost:8080/api/v1/birth-day/  
  
RequestBody
```json:
{
  "user_id": 1,
  "Date": "2022-01-04T00:00:00Z"
}
```
Response
```json:
{}
```