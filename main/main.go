package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/suniln7/urlshortner"
)

func main() {
	mux := defaultMux()

	yamlFilePath := flag.String("yaml", "default.yaml", "Path to the YAML file for URL mappings")
	flag.Parse()

	// Check if the YAML file exists
	yamlData, err := readYAMLFile(*yamlFilePath)
	if err != nil {
		fmt.Printf("Error reading YAML file: %v\n", err)
		fmt.Println("Falling back to the hardcoded map.")
	}

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshortner.MapHandler(pathsToUrls, mux)

	// If the YAML file is valid, build the YAMLHandler using the mapHandler as the fallback
	var yamlHandler http.Handler
	if len(yamlData) > 0 {
		yamlHandler, err = urlshortner.YAMLHandler(yamlData, mapHandler)
		if err != nil {
			fmt.Printf("Error parsing YAML file: %v\n", err)
			fmt.Println("Using mapHandler as fallback.")
			yamlHandler = mapHandler
		}
	} else {
		yamlHandler = mapHandler
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

// Function to read the YAML file and return its contents
func readYAMLFile(filepath string) ([]byte, error) {
	if _, err := os.Stat(filepath); err == nil {
		return os.ReadFile(filepath)
	}
	return nil, fmt.Errorf("YAML file does not exist: %s", filepath)
}
