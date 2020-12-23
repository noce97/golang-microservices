package main

/*
Con go build -o <nombre_que_quieras> main.go se genera un archivo compilado que luego puedes ejecutar sin problema donde
	quieras (dentro de $GOPATH claro). Puedes moverlo a cualquier directorio, que funcionará con ./<nombre_que_hayas_puesto>
*/

import (
	"net/http"
)

func main(){
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World!"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}