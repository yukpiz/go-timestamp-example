package main

import (
	"log"
	"reflect"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func main() {
	var t1 time.Time
	var t2 timestamp.Timestamp
	t3 := ptypes.TimestampNow()

	log.Println(t1)
	log.Println(reflect.TypeOf(t1).Kind())
	log.Println(t2)
	log.Println(reflect.TypeOf(t2).Kind())
	log.Println(t2.String())
	log.Println(t3)
	log.Println(t3.GetSeconds())

	t := time.Unix(t3.GetSeconds(), int64(t3.GetNanos()))
	log.Printf("NOW: %+v\n", t)
}
