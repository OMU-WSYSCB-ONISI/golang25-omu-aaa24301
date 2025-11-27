package main
import (
	"fmt"
	"net/http"
	"runtime"
	 "net/http/httputil"
)
func main() {
	 const url = "https://omu-aaa24301.github.io/"

  req, _ := http.NewRequest("GET", url, nil)

  dump, _ := httputil.DumpRequestOut(req, true)
  fmt.Printf("%s", dump)

  client := new(http.Client)
  resp, err := client.Do(req)
    if err != nil {
		fmt.Println(err)
		return
	}
  defer resp.Body.Close()


  dumpResp, _ := httputil.DumpResponse(resp, true)
  fmt.Printf("%s", dumpResp)
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.Handle("/", http.FileServer(http.Dir("public/")))

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}
