package core

import (
	"log"
)

func Initialize(startupLog string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println(startupLog)
}
