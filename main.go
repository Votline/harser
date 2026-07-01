package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"harser/internal/parser"

	"github.com/Votline/Gurlf"
)

// Config struct contains fields for cfg file
type Config struct {
	FindTitle   []byte `gurlf:"FindTitle"`
	IgnoreTitle []byte `gurlf:"IgnoreTitle,omitempty"`
	FindDesc    []byte `gurlf:"FindDesc,omitempty"`
	IgnoreDesc  []byte `gurlf:"IgnoreDesc,omitempty"`
	Salary      []byte `gurlf:"Salary,omitempty"`
	ExpRange    []byte `gurlf:"ExpRange,omitempty"`
	Schedule    []byte `gurlf:"Schedule,omitempty"`
}

// Employer struct contains fields for 'Employer' response field
type Employer struct {
	Name string `json:"name"`
}

// Salary struct contains fields for 'Salary' response field
type Salary struct {
	From     int    `json:"from"`
	To       int    `json:"to"`
	Currency string `json:"currency"`
}

// Vacancy struct contains fields for vanacy from response
type Vacancy struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Employer     Employer `json:"employer"`
	Salary       *Salary  `json:"salary"`
	URL          string   `json:"url"`
	AlternateURL string   `json:"alternate_url"`
}

// Response struct contains fields for response
type Response struct {
	Items []Vacancy `json:"items"`
	Found int       `json:"found"`
}

func main() {
	const op = "main.main"

	cfgPath := os.Args[1]
	cfg, err := parseCfg(cfgPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: parse config: %s\n", op, err.Error())
		return
	}

	req, err := http.NewRequest("GET", "https://api.hh.ru/vacancies", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: http new request: %s\n", op, err.Error())
		return
	}
	q := req.URL.Query()
	queryApply(&q, cfg)
	req.URL.RawQuery = q.Encode()
	req.Header.Set("User-Agent", "HH-Harser-App/1.0 (votline@gmail.com)")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: http do: %s\n", op, err.Error())
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "%s: http status: %s\n", op, res.Status)
		return
	}

	var apiResp Response
	if err := json.NewDecoder(res.Body).Decode(&apiResp); err != nil {
		fmt.Fprintf(os.Stderr, "%s: json decode: %s\n", op, err.Error())
		return
	}

	for _, v := range apiResp.Items {
		fmt.Printf("%s\n", v.Name)
	}
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

// queryApply set values from cfg to query
func queryApply(q *url.Values, cfg *Config) {
	str := ""
	parser.AddTitle(&str, cfg.FindTitle)
	q.Set("text", str)

	str = ""
	parser.AddSalary(&str, cfg.Salary)
	q.Set("salary", str)
	if str != "" {
		q.Set("only_with_salary", "true")
	}

	str = ""
	parser.AddSchedule(&str, cfg.Schedule)
	q.Set("schedule", str)
}
