package main

import (
  "fmt"
  "net/http"
  "html/template"
)

func main() {
	const PORT = "80"

	http.HandleFunc("/", reflect)
  http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))

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
    <style>
    .ip { color: #FF5722; }
    img { width: 100%; }
    </style>
  </head>
  <body>
    <h1>Welcome {{.User}}</h1>
    <p class="ip">You have originated from IP: {{.IP}}</p>
    <p><img src="img/linkedin-sample.png"></p>
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
