package server

import "log"

func Init() {
	r := NewRouter()

	log.Fatal(r.Run(":8080"))
}
