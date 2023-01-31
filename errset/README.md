# errset

errorを返すような複数の関数をgoroutineで実行する際に、発生したエラーをまとめて補足する。
[golang.org/x/sync/errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup)は最初に発生したエラーしか補足してくれず困ったので実装した。
