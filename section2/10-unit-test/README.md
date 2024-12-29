# 各コマンド
## gotestsインストール（これやらないと　GO: Generate Unit Test For Functionが使えなかった）
`go get -u github.com/cweill/gotests/...`
## テスト実行
`$ go test -v ./section2/10-unit-test`
## カバレッジ計測
`$ go test -v -cover -coverprofile coverage.out ./section2/10-unit-test`
## カバレッジのログ出力
`$ go tool cover -html=coverrage.out ./section2/10-unit-test`