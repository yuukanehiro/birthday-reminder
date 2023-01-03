package main

import (
  "birthday-reminder/routes"
)

func main() {
  e := routes.Init()
  // todo. 環境変数直接ではなく、configファイルを経由して指定する
  // rdb_interface = Rdb.NewRdb(os.Getenv("DB_RDBMS"))
  // db := rdb_interface.ConnectDB()
  // defer db.Close()

  e.Logger.Fatal(e.Start(":8080"))
}
