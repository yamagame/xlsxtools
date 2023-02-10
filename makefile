all: \
	csv2xlsx \
	xlsx2csv \
	xlsxcmp \
	csv2sql

csv2xlsx: cmd/csv2xlsx/main.go
	go build -o bin/csv2xlsx cmd/csv2xlsx/main.go

xlsx2csv: cmd/xlsx2csv/main.go
	go build -o bin/xlsx2csv cmd/xlsx2csv/main.go

xlsxcmp: cmd/xlsxcmp/main.go
	go build -o bin/xlsxcmp cmd/xlsxcmp/main.go

csv2sql: cmd/csv2sql/main.go
	go build -o bin/csv2sql cmd/csv2sql/main.go
