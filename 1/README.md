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




## References
* [Googleが作ったRPCフレームワークgRPCを使ってみた](https://www.sambaiz.net/article/12/)
* [gRPC(Go) で API を実装する](https://blog.fenrir-inc.com/jp/2016/10/grpc-go.html)
* [Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)