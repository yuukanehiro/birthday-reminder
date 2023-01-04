# birthday-reminder
 In order not to forget the birthdays of important people.

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
  
GET  
http://localhost:8080/api/v1/birth-days
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

POST  
http://localhost:8080/api/v1/birth-day  
```json:
{
  "user_id": 1,
  "Date": "2022-02-04"
}
```
or
```json:
{
  "id": 101,
  "user_id": 1,
  "Date": "2022-02-04"
}
```
