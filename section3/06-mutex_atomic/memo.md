排他的制御

以下コマンドでデータ競合があるか確認できる。
`$ go run -race section3/06-mutex_atomic/main.go`

データ競合はmutexで解決する。