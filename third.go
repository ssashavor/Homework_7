package main

import (
	"fmt"
	"log"
	"mux"
	"net/http"
)

var FormTmpl = []byte(`
<html>
<head>
    <meta charset="UTF-8" />
</head>
<body>
<div>Known token: <span id="known-token"></span></div>
<div>
    <form method="POST" action="/">
        <label>Name</label><input name="name" type="text" value="" placeholder ="name" />
        <label>Address</label><input name="address" type="text" value="" placeholder ="address"/>
        <input type="submit" value="submit" />
    </form>
</div>
</body>
<script type="text/javascript">
    const elem = document.querySelector("#known-token");
    let token = document.cookie.split(';').filter((item) => item.trim().startsWith('token='))[0];
    elem.innerHTML = token.replace('token=', '');
</script>
</html>
`)

func HandleGet(w http.ResponseWriter, r *http.Request) {
	w.Write(FormTmpl)
	return
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	address := r.FormValue("address")

	cookie := http.Cookie{
		Name:  "token",
		Value: fmt.Sprintf("%s:%s", name, address),
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HandlePost).Methods(http.MethodPost)
	router.HandleFunc("/", HandleGet).Methods(http.MethodGet)
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
