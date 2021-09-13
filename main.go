package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type ViewData struct {
	Title string
	Users []string
}

func ReadCount(dt *ViewData) {
	var lines []string

	dt.Title = "111"

	file, err := os.Open("log.txt")
	if err != nil {
		//fmt.Println(err)
		//os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err != nil || len(line) == 0 {
			break
		}
		lines = append(lines, string(line))
	}

	//layoutISO = "2006-01-02"
	//layoutUS = "January 2, 2006"

	//new_time := time.Date(lines[1])

	t, _ := time.Parse("Jan _2 15:04:05", lines[1])

	//fmt_time := new_time.Format("15:03:04")
	fmt.Println(t)
	fmt.Println(time.Now().UTC().Format(time.Stamp))

	td := time.Now().Sub(t)
	t_now := time.Now()
	tm := t_now.Add(-td * time.Minute)
	fmt.Println(tm.Format("15:04:03"))
}

func main() {

	Data := ViewData{
		Title: "Статистика печати АЗС",
		Users: []string{"АЗС 1", "АЗС 2", "АЗС 3"},
	}

	fs := http.FileServer(http.Dir("./html/resources/"))
	http.Handle("/resources/", http.StripPrefix("/resources/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Printf(http.Request.p r.URL)
		keys := r.URL.Query()
		if len(keys) > 1 {
			log.Println(keys["time"][0]) //time    login
			//usr := make([]string)
			Data.Users = append(Data.Users, keys["time"][0])

		} else {
			log.Println("Url Param is missing")
		}

		ReadCount(&Data)
		tmpl, _ := template.ParseFiles("html/stat.html")
		tmpl.Execute(w, Data)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":80", nil)
}

//639399
