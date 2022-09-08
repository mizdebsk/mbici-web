package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Artifact struct {
	XMLName xml.Name `xml:"artifact"`
	Type    string   `xml:"type"`
	Name    string   `xml:"name"`
}

type Result struct {
	XMLName       xml.Name   `xml:"result"`
	Id            string     `xml:"id"`
	TaskId        string     `xml:"task"`
	Artifacts     []Artifact `xml:"artifact"`
	Outcome       string     `xml:"outcome"`
	OutcomeReason string     `xml:"outcomeReason"`
	TimeStarted   string     `xml:"timeStarted"`
	TimeFinished  string     `xml:"timeFinished"`
}

type Parameter struct {
	XMLName xml.Name `xml:"parameter"`
	Name    string   `xml:"name"`
	Value   string   `xml:"value"`
}

type Task struct {
	XMLName      xml.Name    `xml:"task"`
	Id           string      `xml:"id"`
	Handler      string      `xml:"handler"`
	Dependencies []string    `xml:"dependency"`
	Parameters   []Parameter `xml:"parameter"`
	Result       Result
}

type Workflow struct {
	XMLName xml.Name `xml:"workflow"`
	Tasks   []Task   `xml:"task"`
	Results []Result `xml:"result"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("WORKFLOW_XML")
	if url == "" {
		url = "https://mbi-artifacts.s3.eu-central-1.amazonaws.com/1653bbcc-1ae3-4eaa-949a-239a24cf8de9/workflow.xml"
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatal(fmt.Errorf("HTTP GET failed: %s", resp.Status))
	}
	resps, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var workflow Workflow
	if err := xml.Unmarshal(resps, &workflow); err != nil {
		resp.Body.Close()
		log.Fatal(err)
	}
	w.Header().Add("Content-Type", "text/html")
	err = Template.Execute(w, workflow)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
