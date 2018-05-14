# [Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)の写経

## 概要

* メッセージフォーマとは`.proto`ファイルに定義される

## なぜprotobufなのか

* gobsはGoだけならいいけど、他のプラットフォームで動かない
* テキストベースだとパースでrun-time costが掛かる
* XMLは人が読めるし、ライブラリも豊富だけど、でかい


## protobufのインストール

※[Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)には載っていない

```sh
❯ brew install protobuf
```

## 独自のprotocol formatを定義する

```proto
syntax = "proto3";
package tutorial;

message Person {
  string name = 1;
  int32 id = 2;
  string email = 3;
  
  enum PhoneType {
    MOBILE = 0;
    HOME = 1;
    WORK = 2;
  }
  
  message PhoneNumber {
    string number = 1;
    PhoneType type = 2;
  }
  
  repeated PhoneNumber phones = 4;
}

message AddressBook {
  repeated Person people = 1;
}
```

## protocのインストール

```sh
❯ go get -u github.com/golang/protobuf/protoc-gen-go
```

## コンパイル

```sh
❯ protoc --go_out=plugins=grpc:tutorial addressbook.proto
```

`addressbook.pb.go`が生成される。

## ファイルに書き出す

```go
package main

import (
	"io/ioutil"
	"log"

	pb "./tutorial"
	"github.com/golang/protobuf/proto"
)

func main() {
	p := &pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	out, err := proto.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("./person.bin", out, 0644); err != nil {
		log.Fatalln("Failed to write person:", err)
	}
}
```

## References
* [Googleが作ったRPCフレームワークgRPCを使ってみた](https://www.sambaiz.net/article/12/)
* [gRPC(Go) で API を実装する](https://blog.fenrir-inc.com/jp/2016/10/grpc-go.html)
* [Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)