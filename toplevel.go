package icecubeapp

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", welcome)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, testingForm)
}

const testingForm = `
<html>
  <head>
    <title>IceCube App</title>
  </head>
  <body>
    <p>Hello world!</p>
  </body>
</html>
`
