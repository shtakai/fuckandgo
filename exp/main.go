package main

import (
	"html/template"
	"os"
)

func main() {
	m := make(map[string]int)
	m["AveKilled"] = 1
	m["TrampKillCount"] = 100
	t, err := template.ParseFiles("hell.gohtml")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name  string
		Moron string
		Map   map[string]int
	}{"Fucked Sucked", "AVE SINxo", m}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
