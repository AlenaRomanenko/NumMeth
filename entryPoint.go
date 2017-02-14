//go build ./gowiki/entryPoint.go
//bash-3.2$ ./entryPoint

package main

import (
	"fmt"
	"gowiki/model"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile("gowiki/model/"+filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "gowiki/model/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func parse(title string) map[string]string {
	var ss = strings.Split(title, "\n")
	var m = make(map[string]string)
	for _, pair := range ss {
		pair = strings.Replace(pair, "\r", "", -1)
		z := strings.Split(pair, "=")
		if len(z) != 2 {
			continue
		} else {
			m[z[0]] = z[1]
		}
	}
	return m
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Choose example </h1>")
	fmt.Fprintf(w, "<p>[<a href='/edit/sqrt'>Root n-order from a</a>]</p>")
	fmt.Fprintf(w, "<p>[<a href='/edit/file1'>x*x - 3.765</a>]</p>")
	fmt.Fprintf(w, "<p>[<a href='/edit/file2'>(1-x*x)*(1-x*x)-x</a>]</p>")
	fmt.Fprintf(w, "<p>[<a href='/edit/file3'>3 - 5*x + x*x*x</a>]</p>")
    fmt.Fprintf(w, "<p>[<a href='/edit/file4'>e^x - 1 - 2*x</a>]</p>")

}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles("gowiki/html/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func save2(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]

	methodChoices := []string{"plain", "pole"}
	methodChoice := "pole"

	for _, v := range methodChoices {
		if v == r.FormValue("methodChoices") {
			methodChoice = v
		}
	}

	fmt.Println(title)
	var choice = 0
	if title == "/file1" {
		choice = 1
	} else if title == "/file2" {
		choice = 2
	} else if title == "/file3" {
		choice = 3
	} else if title == "/file4" {
		choice = 4
	}

	var parsingText = parse(r.FormValue("body"))
	x0, err := strconv.ParseFloat(parsingText["x0"], 64)
	fmt.Println(choice)
	var res, iter = 0.0, 0

	if choice == 0 {
		n, errn := strconv.Atoi(parsingText["n"])
		a, erra := strconv.Atoi(parsingText["a"])
		if erra != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		if errn != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		res, iter = model.Calc(x0, n, a)
	} else {
		if methodChoice == "plain" {
			res, iter = model.CalcNewtone(choice, x0)
		} else {
			c, errc := strconv.ParseFloat(parsingText["c"], 64)
			d, errd := strconv.ParseFloat(parsingText["d"], 64)
			if errc != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			if errd != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			res, iter = model.CalcNewtonePol(choice, x0, c, d)
		}
	}
	

	if iter == model.IterationCount {
		fmt.Fprintf(w, "<h1>%s </h1> ", r.FormValue("body"))
        fmt.Fprintf(w, "<h1>%f </h1> ", x0)
        fmt.Fprintf(w, "<h2 style = 'color:red'>%s </h2> ", "Bad Choice")
	} else {

		fmt.Fprintf(w, "<h1>%s </h1> ", r.FormValue("body"))
		fmt.Fprintf(w, "<label style='font-size:20px'> %s</label>  <input type = text style='font-size:20px' value = '%f' /> ", "Root is", res)
		fmt.Fprintf(w, "<br><label style='font-size:20px'> %s</label>  <input type = text style='font-size:20px' value = '%d' /> ", "Iterations are", iter)
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, "<p style='font-size:20px'>[<a href='/'>Home</a>]</p>")

	http.Redirect(w, r, "/view/"+title, http.StatusFound)

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/save2/", save2)
	http.ListenAndServe(":9090", nil)
}
