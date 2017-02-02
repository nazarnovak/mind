package main

import (
	"log"
	"net/http"
	"os"

	"github.com/desertbit/glue"
	"github.com/nazarnovak/mind/data"
)

//TODO
//User Cookie to authenticate

func main() {
	err := LoadEnvFile()
	if err != nil {
		log.Fatalln(err)
	}



	err = data.OpenDB()
	if err != nil {
		log.Fatalln(err)
	}
	defer data.CloseDB()

	if err != nil {
		log.Fatalln(err)
	}

	err = data.InitHub()
	if err != nil {
		log.Fatalln(err)
	}

	glueSrv := glue.NewServer(glue.Options{
		HTTPSocketType: glue.HTTPSocketTypeNone,
	})
	glueSrv.OnNewSocket(data.HandleSocket)
	defer glueSrv.Release()

	cssFS := http.FileServer(http.Dir("public/css"))
	jsFS := http.FileServer(http.Dir("public/js"))

	http.Handle("/", UIRouter)
	http.Handle("/css/",http.StripPrefix("/css", cssFS))
	http.Handle("/js/", http.StripPrefix("/js", jsFS))
	http.Handle("/glue/", glueSrv)

	port := os.Getenv("PORT")
	log.Printf("Listening on :%s", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
