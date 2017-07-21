package main

import (
	"fmt"
)

const (
	// USER 是你在leetcode.com的用户名
	// 登录后，在页面右上角可以看到
	USER = "aQuaYi"
)

func main() {
	categories := []string{
		"algorithms",
		"database",
		"draft",
		"operating-system",
		"shell",
		"system-design",
	}

	solve := Solveds{}
	cs := Categorys{}
	total := &Category{
		Name: "Total",
	}
	for i := range categories {
		j := getJSON(url(categories[i]))
		c := getCategory(j)

		solve = append(solve, c.run()...)

		total.update(c)
		cs = append(cs, c)
	}
	cs = append(cs, total)
	fmt.Println(solve)
	fmt.Println(cs)
}
