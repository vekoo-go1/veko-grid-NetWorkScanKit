// Veko Grid - Tool CLI berbasis Go untuk eksplorasi jaringan anonim
// Tool ini dibuat untuk penelitian akademik dan audit keamanan legal
package main

import (
	"fmt"
	"os"

	"veko-grid/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
