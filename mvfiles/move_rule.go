package mvfiles

import (
	"encoding/csv"
	"os"
	"strings"
)

type MoveRule struct {
	Pattern string
	Moveto  string
}

func ReadCsv(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func MoveRules(config string) ([]MoveRule, error) {
	lines, err := ReadCsv(config)
	if err != nil {
		return nil, err
	}
	rules := make([]MoveRule, len(lines))
	for i, line := range lines {
		rules[i] = MoveRule{line[0], strings.TrimLeft(line[1], " ")}
	}
	return rules, nil
}
