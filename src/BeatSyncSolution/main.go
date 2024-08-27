package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	spotify
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>BeatSyncSolution</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting web server on :3000")

	runCommand([]string{"sh", "-c", "zotify --help"})

	http.ListenAndServe(":3000", nil)
}
