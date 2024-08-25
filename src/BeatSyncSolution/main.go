package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func runCommand(command []string) {
	cmd := exec.Command(command[0], command[1:]...)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(stdout))
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>BeatSyncSolution</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting web server on :3000")
	os.Setenv("PATH", os.Getenv("PATH")+":/root/.local/bin")
	runCommand([]string{"sh", "-c", "zotify --help"})

	http.ListenAndServe(":3000", nil)
}
