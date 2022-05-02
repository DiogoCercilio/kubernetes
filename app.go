package main
import "net/http"
import "fmt"
import "time"

var startedAt = time.Now()

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/healthz", Healthz)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World again!</h1>"))
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt)
   
	if duration.Seconds() < 10 {
	  w.WriteHeader(500)
	  w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
	  w.WriteHeader(200)
	  w.Write([]byte("ok"))
	}
  }