package main

import (
	"fmt"
	"os"

	"github.com/zziyaa/goset"
)

func main() {
	set := goset.NewSet(true, 1, 2, 3)
	fmt.Fprintf(os.Stdout, "contains 1: %v\n", set.Contains(1))
	fmt.Fprintf(os.Stdout, "contains 5: %v\n", set.Contains(5))
}
