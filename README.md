# Golang xlsx tools

## csv2xlsx

csv ファイルを xlsx ファイルに変換するツール

```bash
# 同じディレクトリに xlsx ファイルが作成される
$ ./bin/csv2xlsx ./testdata/test1.csv
```

```bash
# オプションで出力ファイルを指定できる
$ ./bin/csv2xlsx -o ./testdata/test1.xlsx ./testdata/test1.csv
```

## xlsx2csv

xlsx ファイルの各シートを csv ファイルに変換するツール

```bash
# 同じディレクトリに test2 ディレクトリが作成され、csv ファイルが作成される
$ ./bin/xlsx2csv ./testdata/test2.xlsx
```

以下は出力例、各シートは「連番-シート名.csv」の形式で出力される

```text
testdata
 |_ test2
   |_ 01-Sheet1.csv
   |_ 02-Sheet2.csv
```

```bash
# オプションで出力先のディレクトリを指定できる
$ ./bin/xlsx2csv -o ./testdata ./testdata/test2.xlsx
```

## xlsxcmp

2 つの xlsx ファイルを比較する

```bash
$ ./bin/xlsxcmp ./testdata/test1.xlsx ./testdata/test1.xlsx
same
$ ./bin/xlsxcmp ./testdata/test1.xlsx ./testdata/test2.xlsx
different
```

## ビルド方法

go 言語の開発環境下で以下のコマンドを実行

```bash
$ make
```
