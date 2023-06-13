# goreplaceコマンド

文字列を置換しながらファイルをコピーするコマンド。golangのパッケージをコピーする用途で作成。

```bash
# 使用例 srcディレクトリをdstディレクトリにコピーする
$ goreplace -c config.yaml ./src ./dst

# 使用例 srcディレクトリのmain.goファイルをdstディレクトリにコピーする
$ goreplace -c config.yaml ./src/main.go ./dst
```

以下のYAMLファイルを「-c」オプションで渡すと、ファイルの中に記述された path_filepath 文字列を path_to_hoge 文字列に置換してコピーする。

```yaml
pkgs:
  - src: "path_filepath"
    dst: "path_to_hoge"
```
