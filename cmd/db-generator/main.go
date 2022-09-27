package main

import (
	"context"
	"log"
	"online-teaching/internal/ent"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "test_rw:54rltyi5BCdcm06wu22A0brvvzU5uDgB@tcp(bjfk-d13.yz02:15101)/online-teaching?parseTime=true&loc=Asia%2FShanghai&charset=utf8mb4")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
