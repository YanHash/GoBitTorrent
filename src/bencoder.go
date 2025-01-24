package src

import (
	"os"
	"strconv"
)

type Type string

const (
	Integer     Type = "int"
	String      Type = "str"
	Dictionnary      = "dict"
	List        Type = "list"
)

type bcode struct {
	typeData Type
	data     string
	length   uint
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read(path string) string {
	dat, err := os.ReadFile(path)
	check(err)
	return string(dat)
}

func getIntListDict(s string) (string, int) {
	i := 1
	output := ""
	for string(s[i]) != "e" {
		if i >= len(s) {
			return "", -1
		}
		output += string(s[i])
		i++
	}
	return output, i
}

func getStr(s string) (string, int) {
	i := 1
	slength := ""
	for string(s[i]) != ":" {
		if i >= len(s) {
			return "", -1
		}
		slength += string(s[i])
		i++
	}

	length, err := strconv.Atoi(slength)
	if err != nil {
		return "", -1
	}

	output := string(s[i : i+length])
	return output, length
}

func parse(tstring string) {
	switch string(tstring[0]) {
	case "i":

	}
}

func Decod(path string) {
	tfile := read(path)
	parse(tfile)
}
