package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MemInfo struct {
	CommitLimit float64
	CommitAS    float64
}

func (m *MemInfo) ParseMemInfo(filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		metrics := strings.Split(t, " ")
		switch metrics[0] {
		case "CommitLimit:":
			num, err := strconv.ParseFloat(metrics[4], 64)
			if err != nil {
				return err
			}
			m.CommitLimit = num
		case "Committed_AS:":
			num, err := strconv.ParseFloat(metrics[5], 64)
			if err != nil {
				return err
			}
			m.CommitAS = num
		}
	}
	return nil
}

func main() {
	m := &MemInfo{}
	m.ParseMemInfo("/proc/meminfo")
	fmt.Println(m.CommitAS)
}
