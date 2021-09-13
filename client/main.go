package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"gopkg.in/ini.v1"
)

var Server string
var Port string
var SendTime bool
var Response string
var Login string

//Сэм Ньюмен "Создание микросервисов"

func main() {

	t, _ := time.Parse(time.RFC822Z, "08 Sep 21 13:21 +0600")
	fmt.Println("Now: ", time.Now().Format(time.RFC822Z))
	fmt.Println("Log time: ", t.Format(time.RFC3339Nano))
	td := time.Now().Sub(t)
	fmt.Println("Time duration: ", td.Minutes())
	//----------------------------------------------------------------------------------

	cfg, err := ini.Load("settings.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	Server = cfg.Section("server").Key("ip").String()
	Port = cfg.Section("server").Key("http_port").String()
	SendTime, _ = cfg.Section("send_data").Key("time").Bool()
	Response = "?login=" + Login + "&time=" + strings.Replace(time.Now().Format(time.RFC822Z), " ", "%20", -1)

	resp, err := http.Get("http://" + Server + ":" + Port + Response) //  jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalln(err)
		fmt.Println(resp)
	}

	//fmt.Printf("Data send key \n%s = %s \n", cfg.Section("send_data").Keys()[0].Name(), cfg.Section("send_data").Keys()[0].Value())
	//for a, y := range cfg.Section("server").Keys() {
	//	fmt.Println(a, "    ", y.Name(), " = ", y.Value())
	//}

	// Classic read of values, default section can be represented as empty string
	//fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
	//fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())

	// Let's do some candidate value limitation
	//fmt.Println("Server Protocol:",
	//	cfg.Section("server").Key("protocol").In("http", []string{"http", "https"}))
	// Value read that is not in candidates will be discarded and fall back to given default value
	//fmt.Println("Email Protocol:",
	//	cfg.Section("server").Key("protocol").In("smtp", []string{"imap", "smtp"}))

	// Try out auto-type conversion
	//fmt.Printf("Port Number: (%[1]T) %[1]d\n", cfg.Section("server").Key("http_port").MustInt(9999))
	//fmt.Printf("Enforce Domain: (%[1]T) %[1]v\n", cfg.Section("server").Key("enforce_domain").MustBool(false))

	// Now, make some changes and save it
	//	cfg.Section("").Key("app_mode").SetValue("production")
	//	cfg.SaveTo("my.ini.local")
}
