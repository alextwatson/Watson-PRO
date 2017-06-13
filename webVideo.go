package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var vid []byte
var inputFile string

func main() {
	http.HandleFunc("/motion/notification", notification)
	http.HandleFunc("/motion/video", video)
	http.HandleFunc("/", owner)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func notification(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s", req.Proto, req.URL)
	println(req)
}

func video(w http.ResponseWriter, req *http.Request) {

	//log.Printf("%s %s\n", req.Proto, req.URL)
	vidAVI, _ := ioutil.ReadAll(req.Body)
	inputFile = "/tmp/camra.avi"
	err := ioutil.WriteFile(inputFile, vidAVI, 0644)

	if err != nil {
		println("ERROR=", err)
	}
	cmd := exec.Command("/usr/local/bin/ffmpeg", "-i", inputFile, inputFile+".webm", "-y")

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)

	print(string(vid))

}
func owner(w http.ResponseWriter, req *http.Request) {
	//	log.Printf("%s %s\n", req.Proto, req.URL)
	w.Header().Set("Content-Type", "video/webm")
	content, _ := ioutil.ReadFile(inputFile + ".webm")
	io.WriteString(w, string(content))
}
