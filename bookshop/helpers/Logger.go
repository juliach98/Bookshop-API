package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Logger struct {
	Timestamp string `json:"timestamp"`
	App       string `json:"app"`
	Name      string `json:"name"`
	Warn      string `json:"warn"`
	Url       string `json:"url"`
	ClientIP  string `json:"client_ip"`
	UserID    string `json:"user_id"`
	Message   string `json:"go_message"`
}

func (d *Logger) Print(name string, warn string, url string, ip string, user string, message string) {
	const timeLayout = "2006/01/02 15:04:05 -07:00"
	const logDate = "2006-01-02"

	d.Timestamp = time.Now().Format(timeLayout)
	d.App = "Bookshop"
	d.Name = name
	d.Warn = warn
	d.Url = url
	d.ClientIP = ip
	d.UserID = user
	d.Message = message
	data, _ := json.Marshal(d)

	file, err := os.OpenFile("logs/temp_"+time.Now().Format(logDate)+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString(string(data) + "\n"); err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
