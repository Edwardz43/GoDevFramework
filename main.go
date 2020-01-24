package main

import "Edwardz43/godevframework/routes"

func main() {
	r := routes.Serve()
	r.Run(":8765") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
