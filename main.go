package main

import (
	. "gin/database"
)

func main() {
	defer SqlDb.Close()
	router := initRouter()
	router.Run(":8000")
}
