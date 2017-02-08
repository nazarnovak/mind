package main

import (
	"log"
	"net/http"

	"github.com/desertbit/glue"
	"github.com/nazarnovak/mind/data"
)

func main() {
	err := LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	err = data.OpenDB(
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.User,
		conf.DB.Password,
		conf.DB.Name,
	)

	if err != nil {
		log.Fatalln(err)
	}
	defer data.CloseDB()

	if err != nil {
		log.Fatalln(err)
	}

	err = data.InitHub(conf.Redis.URL)
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

	// NotLikeThis
	data.Greet = conf.App.Greet

	http.Handle("/", UIRouter)
	http.Handle("/css/",http.StripPrefix("/css", cssFS))
	http.Handle("/js/", http.StripPrefix("/js", jsFS))
	http.Handle("/glue/", glueSrv)

	log.Printf("Listening on :%s", conf.App.Port)
	err = http.ListenAndServe(":" + conf.App.Port, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
