package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func commandExecute(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("path", r.URL.Path)
	GETCommand := strings.Replace(r.URL.Path, "/", "", 1)
	fmt.Printf("%s\n", GETCommand)
	len := len(GETCommand)
	var args []string
	var command string
	if len > 1 {
		args = strings.Split(GETCommand, ",")
	} else {
		args = append(args, "ls")
		args = append(args, "-al")
	}
	command = fmt.Sprintf("/usr/bin/which")
	out, err := exec.Command(command, args[0]).Output()
	if err != nil {
		fmt.Printf("Error in which command: %s", err)
		fmt.Fprintf(w, "Error in which command: %s", err)
	}
	realCommand := strings.Split(string(out), "\n")[0]
	args[0] = realCommand
	toExecute := fmt.Sprintf("%s", strings.Join(args, " "))
	fmt.Printf("RealCommand:\t%s\nBuildCommand:\t%s\nStarArgs:\t%s\nRealArgs:\t%s\n", realCommand, toExecute, args, GETCommand)
	fmt.Fprintf(w, "RealCommand:\t%s\nBuildCommand:\t%s\nStarArgs:\t%s\nRealArgs:\t%s\n", realCommand, toExecute, args, GETCommand)
	outCommand, err := exec.Command("bash", "-c", toExecute).Output()
	if err != nil {
		fmt.Printf("Error in exec: %s", err)
		fmt.Fprintf(w, "Error in exec: %s", err)
	}
	fmt.Fprintf(w, "\nResult:\n\n%s\n", outCommand)
}

func main() {
	http.HandleFunc("/", commandExecute)
	port := 9090
	listenPort := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(listenPort, nil)
	fmt.Printf("Listen port: %d", port)
	if err != nil {
		fmt.Printf("ListenAndServer Fatal: %s", err)
	}

}
