package main

import (
	"github.com/golang/glog"
	"self/go-grpc/pkg/protobuf"
)

func main() {

	defer glog.Flush()

	checkErr(protobuf.Start())
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

