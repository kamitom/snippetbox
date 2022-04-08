package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	log.Println("new uuid: ", uuid.NewString())

	// Get the IP address from an incoming HTTP request
	ip, err := getIP(r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "No valid IP", 400)
		return
	}

	log.Println("incoming IP: ", ip)

	// Initialize a slice containing the paths to the two files. Note that the
	// home.page.tmpl file must be the *first* file in the slice.

	// Include the footer partial in the template files.
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// Use the template.ParseFiles() function to read the files and store the
	// templates in a template set. Notice that we can pass the slice of file paths
	// as a variadic parameter?
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// We then use the Execute() method on the template set to write the template
	// content as the response body. The last parameter to Execute() represents any
	// dynamic data that we want to pass in, which for now we'll leave as nil.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	// fmt.Fprintf(w, `
	// <h1>Hello from Helenfit...</h1>
	// <h2>www.helenfit.com</h2>
	// <h3>uuid: %s</h3>
	// `, uuid.NewString())
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, `
	<h1>Display a specific snippet with ID %d...</h1>
	<h2>www.helenfit.com</h2>
	<h3>uuid: %s </h3>
	`, id, uuid.NewString())
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func getIP(r *http.Request) (string, error) {
	//Get IP from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}
	return "", fmt.Errorf("No valid ip found")
}
