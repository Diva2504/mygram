package main

import (
	"github.com/takadev15/mygram-api/databases"
	"github.com/takadev15/mygram-api/routes"
)

func main() {
  databases.DBInit()
  r := routes.RouteList()
  r.Run(":3030")
}
