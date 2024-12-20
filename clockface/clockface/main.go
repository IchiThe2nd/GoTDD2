package main

import (
	"os"
	"time"

	"github.com/IchiThe2nd/GoTDD2/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
