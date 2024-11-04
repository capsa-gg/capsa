package main

import (
	"fmt"
	"time"

	"github.com/lucianonooijen/capsa/server/internal/cmd"
)

func main() {
	start := time.Now()

	cmd.Execute()

	end := time.Now()
	took := end.Sub(start).String()
	fmt.Printf("program took %s to run\n", took)
}
