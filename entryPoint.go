package main

import (
	"fmt"
	"NumMeth/model"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"NumMeth/service"
	"math"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile("NumMeth/model/"+filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "NumMeth/model/" + title + ".txt"
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
	p := Page{Title: "Title", Body: []byte("")}
	renderTemplate(w, "home", &p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles("NumMeth/html/"+tmpl+".html", "NumMeth/html/header.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		//p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]

	methodChoices := []string{"plain", "pole", "modifPole"}
	methodChoice := "pole"

	for _, v := range methodChoices {
		if v == r.FormValue("methodChoices") {
			methodChoice = v
		}
	}

	var choice = 0

	if title == "file1" {
		choice = 1
	} else if title == "file2" {
		choice = 2
	} else if title == "file3" {
		choice = 3
	} else if title == "file4" {
		choice = 4
	} else if title == "file5" {
		choice = 5
	} else if title == "file6" {
		choice = 6
	}

	var res, iter, value = 0.0, 0, 0.0

	var parsingText = parse(r.FormValue("body"))
	if choice == 0 {
		n, errn := strconv.Atoi(parsingText["n"])
		a, erra := strconv.Atoi(parsingText["a"])
		if erra != nil {
			http.Error(w, erra.Error(), 500)
			return
		}
		if errn != nil {
			http.Error(w, errn.Error(), 500)
			return
		}
		res, iter, value = service.CalcSQRT(n, a)
	} else {

		x0, err := strconv.ParseFloat(parsingText["x0"], 64)
		if methodChoice == "plain" {
			res, iter, value = service.CalcNewtone(choice, x0)
		} else if methodChoice == "pole" {
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
			res, iter, value = service.CalcNewtonePol(choice, x0, c, d)
		} else {
			res, iter, value = service.CalcFlyNewtonePol(choice, x0)
		}
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	fmt.Fprintf(w, "<a style=' background: blueviolet;"+"padding:5px;	border-radius:5px;	"+"font-size:20px;color:white;font-size:20px' href='/'>Home</a>")
	fmt.Fprintf(w, "<div style=' background: blueviolet;"+"padding:5px;	border-radius:5px;	"+"font-size:20px;color:white;font-size:20px; position: absolute; "+"top:0; right:0; width:150px; text-align:center;font-size:20px; color:white'>Romanenko Olena </div>")

	if math.IsNaN(value) || iter == -1 || iter == model.IterationCount {
		fmt.Fprintf(w, "<h1>%s </h1> ", r.FormValue("body"))
		//fmt.Fprintf(w, "<h1>%f </h1> ", x0)
		fmt.Fprintf(w, "<h2 style = 'color:red'>%s </h2> ", "Оберіть інші початкові значення")
	} else {
		if choice == 6 {
			fmt.Fprintf(w, "<p/> <p/><div style='font-size:20px'> %s %4.2f %s </div>", "Для збільшення капіталу у 8 разів за рік необхідно знайти депозит під", float64(1200)*(res-1), " відсотків річних")
		} else {
			fmt.Fprintf(w, "<h1>%s </h1> ", r.FormValue("body"))
			fmt.Fprintf(w, "<label style='font-size:20px'>%12s</label>  <input type = text style='font-size:20px' value = '%9.12f' />", "Корінь", res)
			fmt.Fprintf(w, "<p><br><label style='font-size:20px'>%12s</label>  <input type = text style='font-size:20px' value = '%9.12f' /> ", "Нев'язка", value)
			fmt.Fprintf(w, "<p> <br><label style='font-size:20px'>%12s</label>  <input type = text style='font-size:20px' value = '%d' />", "Ітерацій", iter)
		}
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.Handle("static", http.StripPrefix("/static/", http.FileServer(http.Dir("NumMeth/static"))))
	http.ListenAndServe(":9090", nil)
}
