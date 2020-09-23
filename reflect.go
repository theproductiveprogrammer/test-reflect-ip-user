package main

import (
  "fmt"
  "net/http"
  "html/template"
)

func main() {
	const PORT = "80"

	http.HandleFunc("/", reflect)

	fmt.Printf("\nServer started on port %s...\n", PORT)
	http.ListenAndServe(":"+PORT, nil)
}


func reflect(w http.ResponseWriter, r *http.Request) {
  const tpl = `
<!doctype html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>Welcome {{.User}}</title>
  </head>
  <body>
    <h1>Welcome {{.User}}</h1>
    <p>You have originated from IP: {{.IP}}</p>
  </body>
</html>
`
t, err := template.New("index").Parse(tpl)
if err != nil {
  w.WriteHeader(500)
  w.Write([]byte(err.Error()))
  return
}

data := Req{ "Test", "127.0.0.1" }

err = t.Execute(w, data)
if err != nil {
  w.WriteHeader(500)
  w.Write([]byte(err.Error()))
}

}

type Req struct {
  User string
  IP   string
}
