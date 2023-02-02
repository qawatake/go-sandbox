# template-sandbox

`html/template`で遊んでみる。

## template.go

`tmpls/components.html.gotempl`に再利用可能なtemplateを用意しておき、それをいろいろなところで使いまわしている。

## template_test.go

golden file testを書いてみた。`go test -update`でgolden fileの更新を行うことができる。
