package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const scriptname = "temp.sh"

func main() {
	// http.HandleFunc("/pull/", pullHandler)
	http.HandleFunc("/", cmdHandler)
	if err := http.ListenAndServe(":7002", nil); err != nil {
		log.Println("fail to listen 7002:", err)
	}
}

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	var output string
	cmd, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(cmd))
	if len(cmd) > 0 {
		output = exe_cmd(cmd)
	} else {
		output = "param cmd is required"
	}
	fmt.Fprintf(w, output)
}

// func pullHandler(w http.ResponseWriter, r *http.Request) {
// 	partOfURL := r.URL.Path[len("/pull/"):]
// 	switch partOfURL {
// 	case "nginx":
// 		out := exe_bash("./pull_nginx.sh")
// 		fmt.Fprintf(w, out)
// 	case "weixin":
// 		out := exe_bash("./pull_weixin.sh")
// 		fmt.Fprintf(w, out)
// 	default:
// 		fmt.Fprintf(w, "no script")
// 	}
// }

// func exe_bash(scriptPath string) string {
// 	return exe_cmd("bash " + scriptPath)
// }

func exe_cmd(cmd []byte) string {
	writeCmdToTempShell(cmd)
	out, err := exec.Command("bash", scriptname).CombinedOutput()
	outstr := string(out)
	if err != nil {
		outstr += err.Error()
	}
	fmt.Println(outstr)
	return outstr
}

func writeCmdToTempShell(cmd []byte) {
	f, _ := os.OpenFile(scriptname, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	defer f.Close()
	f.Write(cmd)
}
