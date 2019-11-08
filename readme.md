```go
package main

import (
	"fmt"
	"log"

	"github.com/jakobii/tsql"
)

func main() {

	Conn := tsql.Connection{
		Server:   "localhost",
		Database: "...",
		Username: "...",
		Password: "...",
		Port:     5432,
	}
	cols, err := tsql.GetColumns(Conn, "<schema>", "<table>")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cols)

}
```