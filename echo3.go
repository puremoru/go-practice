package main 

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now();
	fmt.Println(strings.Join(os.Args[1:], " "))
	end := time.Now();
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
}