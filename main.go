// main.go

package main

import "os"

func main() {
	port := os.Getenv("APP_PORT")
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_URL"))

	a.Run(":" + port)
}
