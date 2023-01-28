package main

import (
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

type Handler struct {

}

func (*Handler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	body, _ := ioutil.ReadAll(request.Body)
	data := WebHookData{}
	json.Unmarshal(body, &data)
	logger := Logger{
		Path: getLogPath(data.Repository.HTMLURL),
	}
	cmd := exec.Command("/bin/sh", getCmdPath(data.Repository.HTMLURL))
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil || io.EOF == err {
			break
		}
		logger.Write(line)
	}
	cmd.Wait()
}

func getCmdPath(repoUrl string) string {
	return "/root/Document/" + strings.TrimLeft(repoUrl, "https://") + "/ci.sh"
}

func getLogPath(repoUrl string) string {
	return "/root/Document/" + strings.TrimLeft(repoUrl, "https://") + ".ci.log"
}


func main() {
	http.ListenAndServe("0.0.0.0:8899", &Handler{})
}