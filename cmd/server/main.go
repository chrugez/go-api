package main

import (
	"github.com/chrugez/go-api/internal/routers"
)

func main() {
  r := routers.NewRouter()
  
  r.Run(":8002") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
