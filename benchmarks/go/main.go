// SPDX-License-Identifier: BSD-3-Clause
package main

import (
	"fmt"
	"strings"

	"github.com/go-ruby-csv/csv"
)

func buildCSV() string {
	var b strings.Builder
	for i := 0; i < 200; i++ {
		fmt.Fprintf(&b, "%d,name-%d,%d,%d,%d,%d\n", i, i, i*1, i*2, i*3, i*4)
	}
	return b.String()
}

func main() {
	opts := csv.DefaultOptions(csv.Options{})
	data := buildCSV()
	rows, _ := csv.ParseRows(data, opts)
	bench("parse-200x6", 500, func() { v, _ := csv.ParseRows(data, opts); sink = v })
	bench("generate-200x6", 500, func() { v, _ := csv.Generate(rows, opts); sink = v })
}
