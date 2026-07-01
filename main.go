package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"unsafe"

	"github.com/Votline/Gurlf"
)

type Config struct {
	FindTitle   []byte `gurlf:"FindTitle"`
	IgnoreTitle []byte `gurlf:"IgnoreTitle"`
	FindDesc    []byte `gurlf:"FindDesc"`
	IgnoreDesc  []byte `gurlf:"IgnoreDesc"`
	SalaryRange []byte `gurlf:"SalaryRange"`
}

func main() {
	const op = "main.main"

	cfgPath := os.Args[1]
	cfg, err := parseCfg(cfgPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: parse config: %s\n", op, err.Error())
		return
	}
	fmt.Println(cfg)

	req, err := http.NewRequest("GET", "https://api.hh.ru/vacancies", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: http new request: %s\n", op, err.Error())
	}
	queryStr := ""
	addFindTitle(&queryStr, cfg.FindTitle)
	q := req.URL.Query()
	q.Add("text", queryStr)
}

// parseCfg parses config file
// found 'ignore' and 'find' sections
// returns Config struct
func parseCfg(path string) (*Config, error) {
	const op = "main.parseCfg"

	sdata, err := gurlf.ScanFile(path)
	if err != nil {
		return nil, fmt.Errorf("%s: gurlf scanfile: %w", op, err)
	}
	cfgdata := sdata[0]

	var cfg Config
	if err := gurlf.Unmarshal(cfgdata, &cfg); err != nil {
		return nil, fmt.Errorf("%s: gurlf unmarshal :%w", op, err)
	}

	return &cfg, nil
}

// rangeByByte call yield for each matching row
func rangeByByte(src []byte, sep byte, yield func(start, end int)) {
	start := 0
	for start < len(src) {
		end := bytes.IndexByte(src[start:], sep)
		if end == -1 {
			end = len(src)
		} else {
			end += start
		}
		yield(start, end)
		start = end + 1
	}
}

// addFindTitle add values from 'FindTitle' field
// to query string
func addFindTitle(query *string, find []byte) {
	var buf bytes.Buffer
	buf.Write([]byte("NAME:("))
	rangeByByte(find, byte(','), func(start, end int) {
		if start == end {
			return
		}
		src := find[start:end]
		buf.Write(src)
		buf.Write([]byte(" OR "))
	})
	findBytes := buf.Bytes()

	lastOrIdx := bytes.LastIndex(findBytes, []byte(" OR "))
	if lastOrIdx != -1 {
		findBytes = findBytes[:lastOrIdx]
	}
	findBytes = append(findBytes, ')')

	findStr := unsafe.String(unsafe.SliceData(findBytes), len(findBytes))
	*query = findStr
}
