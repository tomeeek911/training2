package main

import "fmt"

func main() {

	x := map[string][]string{
		`bond_james`: []string{`not stirred`, `martinies`, `women`},
		`mm`:         []string{`james_bond`, `literature`, `computer science`},
		`no_dr`:      []string{`being evil`, `ice cream`, `sunsets`},
	}
	delete(x, `bond_james`)
	x[`tomasz`] = []string{`29`}
	for k, v := range x {
		fmt.Println(k, v)
	}
}
