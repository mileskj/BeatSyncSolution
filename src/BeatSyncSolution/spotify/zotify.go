package spotify

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var zotifyPath = "/root/.local/bin"

func runCommand(command []string) {
	cmd := exec.Command(command[0], command[1:]...)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(stdout))
	//Don't need to worry about saving the stdout until I find that I do need to worry about it
	//Print the Stdout of this, but should I also be tracking the stderr?
}

func setupZotify() {
	//Adds Zotify to the PATH so we can call it normally
	os.Setenv("PATH", (os.Getenv("PATH") + ":" + zotifyPath))

	//Need to figure out authentication
	//Authentication has changed because of the Spotify API, now need to run a script to get a proper credentials.json here: https://github.com/dspearson/librespot-auth.git
	//Rust toolchain had to be installed to then run this script, to then run a build on the script from the repo above
	//Maybe need to include this stuff in the Dockerfile on setup
	//If this is a librespot problem, then this will happen regardless on if I use Zotify or not
	//It seems to specifically get a complicated auth token for your account, maybe theres a way to do this automatically?
}
