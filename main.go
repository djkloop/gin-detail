package main

import "gin_demo/server"

func main() {
	r := server.ServerRouter()
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
