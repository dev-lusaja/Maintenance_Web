package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
)

// Esta variable apunta a la direccion en memoria de Template
var Templates *template.Template

var Config struct {
	Host          string
	Retry_time    string
	Port          string
	Template_dir  string
	Template_name string
	Static_route  string
	Static_dir    string
}

func init() {
	// Seteamos las variables de configuracion
	Config.Host = os.Getenv("HOST")
	Config.Retry_time = os.Getenv("RETRY_TIME")
	Config.Port = os.Getenv("MAINTENANCE_PORT")
	Config.Template_dir = os.Getenv("TEMPLATE_DIR") + "*"
	Config.Template_name = os.Getenv("TEMPLATE_NAME")
	Config.Static_route = os.Getenv("STATIC_ROUTE")
	Config.Static_dir = os.Getenv("STATIC_DIR")

	// Asignamos la ruta donde se almacenan los Templates
	Templates = template.Must(template.ParseGlob(Config.Template_dir))
}

func Maintenance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Retry-After", Config.Retry_time)
	w.WriteHeader(http.StatusServiceUnavailable)
	Templates.ExecuteTemplate(w, Config.Template_name, nil)
}

func MaintenanceJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Retry-After", Config.Retry_time)
	w.WriteHeader(http.StatusServiceUnavailable)
	mensaje := make(map[string]string)
	mensaje["mensaje"] = "Estamos en mantenimiento"
	json, _ := json.Marshal(mensaje)
	w.Write(json)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Maintenance).Host(Config.Host)
	r.HandleFunc("/", Maintenance).Host("{subdomain}." + Config.Host)
	r.HandleFunc("/json", MaintenanceJson).Host(Config.Host)
	r.HandleFunc("/json", MaintenanceJson).Host("{subdomain}." + Config.Host)

	// Directorios estaticos
	s := http.StripPrefix(Config.Static_route, http.FileServer(http.Dir(Config.Static_dir)))
	r.PathPrefix(Config.Static_route).Handler(s)
	http.Handle("/", r)

	// Inicializando Servidor Web
	fmt.Println("Escuchando en el puerto: " + Config.Port)
	http.ListenAndServe(":"+Config.Port, nil)
}
