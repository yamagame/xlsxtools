all: \
	csv2xlsx \
	xlsx2csv \
	xlsxcmp

csv2xlsx: cmd/csv2xlsx/main.go
	go build -o bin/csv2xlsx cmd/csv2xlsx/main.go

xlsx2csv: cmd/xlsx2csv/main.go
	go build -o bin/xlsx2csv cmd/xlsx2csv/main.go

xlsxcmp: cmd/xlsxcmp/main.go
	go build -o bin/xlsxcmp cmd/xlsxcmp/main.go
