package tool

import (
	"log"
	"time"
)

func TimeCost(start time.Time, mark string) {
	terminal := time.Since(start)
	log.Println(mark, terminal)
}
