package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"projects/api-products/product"
)

// Variables
const port string = ":3000"

var tpl *template.Template

func init() {
	//tpl = template.Must(template.ParseGlob("/home/dev/projects/api-products/web/templates/*.html"))
	tpl = template.Must(template.ParseGlob("D:/Gdrive-MSE/projects/api-products/web/templates/*.html"))
}

func main() {

	// unc http.HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	http.HandleFunc("/", ProductHtml)
	http.HandleFunc("/products", ListProduct)

	// func http.ListenAndServe(addr string, handler http.Handler) error
	// func log.Panic(v ...any)
	log.Panic(http.ListenAndServe(port, nil))
}

func ProductHtml(w http.ResponseWriter, r *http.Request) {

	Data := product.ListProducts()

	// func (*template.Template).Execute(wr io.Writer, data any) error
	log.Panic(tpl.ExecuteTemplate(w, "index.html", Data))
}

func ListProduct(w http.ResponseWriter, r *http.Request) {
	// func (http.ResponseWriter).Header() http.Header
	// func (http.Header).Set(key string, value string)
	w.Header().Set("Content-Type", "application/json")

	// unc json.NewEncoder(w io.Writer) *json.Encoder
	// func (*json.Encoder).Encode(v any) error
	json.NewEncoder(w).Encode(product.ListProducts())

	// not working ??????
	// func json.Marshal(v any) ([]byte, error)
	// json.Marshal(product.ListProducts())
}
