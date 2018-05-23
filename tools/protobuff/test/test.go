package main

import (
	"github.com/rockamring/goproject/tools/protobuff/proto"
	"github.com/golang/protobuf/proto"
	"log"
)

func main(){
	test:=&example.Test{
		Label: proto.String("hello"),
		Type: proto.Int32(30),
		Reps: []int64{1, 2, 3},
		Optionalgroup: &example.Test_OptionalGroup{
			RequiredField: proto.String("world"),
		},
	}
	data, err := proto.Marshal(test)

	if err != nil{
		log.Fatal("marshal error: ", err)
	}

	newtest := &example.Test{}
	err = proto.Unmarshal(data, newtest)
	if err != nil{
		log.Fatal("unmarshal error: ", err)
	}

	log.Println(newtest)
}
