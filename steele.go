package main

import (
	"fmt"
	"image/png"
	"net/http"
	"text/template"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type pageInfo struct {
	Title string
}

func main() {
	fmt.Print("Setting up the webserver...")
	http.HandleFunc("/", homepage)
	http.HandleFunc("/generator", qrCode)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Print("There was an error with setting up the webserver: ", err, ". Check you network config. If that doesn't work, then please create a GitHub Issue (https://git.io/JYMMA)")
	}
	fmt.Print("Successful!\nHead over to http://localhost to get to the generator.")
}
func homepage(w http.ResponseWriter, r *http.Request) {
	pageTitle := pageInfo{Title: "Home"}
	templateFile, _ := template.ParseFiles("index.html")
	templateFile.Execute(w, pageTitle)
}
func qrCode(w http.ResponseWriter, r *http.Request) {
	dataString := r.FormValue("dataString")

	qrCode, _ := qr.Encode(dataString, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)

	png.Encode(w, qrCode)
}
