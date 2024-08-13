package main

import (
  "fmt"
  "net"
  "net/http"
  "log"
  "encoding/json"
  "io"
  "time"
)

type Alert_Request struct {
	Status string `json:"status"`
	Alerts []struct {
		Status string
		Labels map[string]string
	}
}

func alert_webhook(w http.ResponseWriter, r *http.Request) {
    var req Alert_Request
    j, _ := io.ReadAll(r.Body)
    json.Unmarshal(j, &req)
    fmt.Printf("%v\n", req)
    severity := req.Alerts[0].Labels["severity"]
    switch severity {
      case "critical":
          ChangePatlite(1,0,0,1) // red & beep
      case "warning":
          ChangePatlite(0,1,0,1) // yellow & beep
      default:
          fmt.Println("default")
    }

    go func(){
      time.Sleep(1000 * time.Millisecond)
      ChangePatlite(0,0,0,0)
    }()
}

func main(){
  ChangePatlite(0,0,0,0)

  http.HandleFunc("/alert_webhook", alert_webhook)
  fmt.Println("Listening 0.0.0.0:8085")
  log.Fatal(http.ListenAndServe(":8085", nil))
}

func ChangePatlite(red int, yellow int, green int, beep int){
  data := (beep & 0x1) << 3 | (green & 0x1) << 2 | (yellow & 0x1) << 1 | (red & 0x1)
  conn, _ := net.Dial("udp", "172.16.254.240:10000")
  fmt.Fprintf(conn, "%c%c", 0x57, data)
  //fmt.Printf("%c%x", 0x57, data)
}
