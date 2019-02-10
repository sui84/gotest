package main

import (
    //"encoding/json"
    "fmt"
    "io/ioutil"
	"net/http"
	"io"
)
/*type Myobj struct {
    message  string
}
*/

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//调用rest api
		url := "http://localhost:8080/ping"
		//resp, err := http.Get(url, "application/json", bytes.NewBuffer(jsonData))
		resp, err := http.Get(url)
		body, err := ioutil.ReadAll(resp.Body)
	    if err != nil {
	        fmt.Printf("read body err, %v\n", err)
	        return
	    }
	    /*
	    var myobj Myobj
	    err := json.Unmarshal(body,&myobj)
	    if err != nil {
	        fmt.Println("error:", err)
	    }
	    */
	    fmt.Println("r.body:", r.Body)
		io.WriteString(w, string(body))
	})
	http.ListenAndServe(":8001", nil)
}

