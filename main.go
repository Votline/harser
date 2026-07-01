package main

import (
	"fmt"
	"os"

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
	cfgPath := os.Args[1]
	cfg, err := parseCfg(cfgPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse config: %v\n", err)
		return
	}
	fmt.Println(cfg)
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
