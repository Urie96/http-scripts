package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/pull/", pullHandler)
	http.HandleFunc("/cmd", cmdHandler)
	http.ListenAndServe(":7002", nil)
}

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	cmd := r.URL.Query()["cmd"][0]
	output := exe_cmd(cmd)
	fmt.Fprintf(w, output)
}

func pullHandler(w http.ResponseWriter, r *http.Request) {
	partOfURL := r.URL.Path[len("/pull/"):]
	switch partOfURL {
	case "nginx":
		out := exe_bash("./pull_nginx.sh")
		fmt.Fprintf(w, out)
	case "weixin":
		out := exe_bash("./pull_weixin.sh")
		fmt.Fprintf(w, out)
	default:
		fmt.Fprintf(w, "no script")
	}
}

func exe_bash(scriptPath string) string {
	return exe_cmd("bash " + scriptPath)
}

func exe_cmd(cmd string) string {
	log.Println(cmd)
	if strings.HasPrefix(cmd, "./") {
		return exe_bash(cmd)
	}
	parts := strings.Fields(cmd)
	fmt.Printf("%q\n", parts)
	out, err := exec.Command(parts[0], parts[1:]...).CombinedOutput()
	outstr := string(out)
	if err != nil {
		outstr += err.Error()
	}
	fmt.Println(outstr)
	return outstr
}
