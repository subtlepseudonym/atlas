/* Map Generation Service
 * subtlepseudonym (subtlepseudonym@gmail.com)
 */

package main

import (
 	"atlas/atlas"

	"bytes"
	"errors"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var (
	lasPort string = ":10010"
	routerIP string = "192.168.1.1"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Index (currently) pulls boostrap from maxcdn rather than local storage
	http.ServeFile(w, r, "index.json")
}

func getIP() (string, error) {
	// Queries router for intranetwork ip address then pipes it to awk
	ip := exec.Command("ip", "route", "get", routerIP)
	awk := exec.Command("awk", "{ print $NF; exit }")

	// Connecting the pipe
	reader, writer := io.Pipe()
	var buf bytes.Buffer
	ip.Stdout = writer
	awk.Stdin = reader
	awk.Stdout = &buf

	ip.Start()
	awk.Start()

	ip.Wait()
	writer.Close()

	awk.Wait()
	reader.Close()

	ret := strings.TrimSpace(buf.String())
	if ret == "" {
		return "", errors.New("Error retrieving intranetwork ip from " + routerIP)
	}

	return ret, nil
}

func init() {
	rand.Seed(6)

	// Enable logging to local file
	log.SetOutput(os.Stdout)

	// Print public ip address and ListenAndServe port
	pub_ip, err := getIP()
	if err != nil {	log.Fatal(err) }
	log.Println("ListenAndServe at", pub_ip+lasPort)
}

func main() {
	atlas.AtlasTest()

	http.HandleFunc("/", indexHandler)

	log.Fatal(http.ListenAndServe(lasPort, nil))
}
