package main

import (
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("./*.html"))
}

var fm = template.FuncMap{
	"doubleNum": doubleNum,
	"sqrt":      sqrt,
	"sqr":       sqr,
}

func doubleNum(n int) int {
	return n + n
}

func sqrt(n float64) float64 {
	return math.Sqrt(n)
}

func sqr(n int) float64 {
	return math.Pow(float64(n), 2)
}

func main() {

	tpl.ExecuteTemplate(os.Stdout, "index.html", 3)

}
