## pbファイルの出力

`.proto` ファイルに定義したメッセージやサービスをprotocコマンドでコンパイルして、
Go用のサーバインタフェースと、RPCを実行するためのクライアントプログラムであるスタブ（クライアントスタブ）のコードを生成する。

`protoc --proto_path ./echo/proto --go_out=plugins=grpc:./echo/proto/ ./echo/proto/echo.proto`

protocコマンドで指定するもの

- `.proto` ファイルのパス
- オプションとして `.proto` ファイルを検索するディレクトリ
- オプションとして出力する言語と出力先
  - `--xx_out` のように `xx` で言語を指定するっぽい？
  - `--go_out=plugins=grpc` のようにしてgRPCを **プラグイン** として指定する
    - コードの出力先をコロン以下で指定する

※サーバとクライアントの実装についてはコードを参照すること。

serverの実行: `go run ./echo/server`

clientの実行: `go run ./echo/client testest`

