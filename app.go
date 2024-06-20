package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	//"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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
	Result       *Result
	Bucket       string
}

type Workflow struct {
	XMLName xml.Name `xml:"workflow"`
	Tasks   []Task   `xml:"task"`
	Results []Result `xml:"result"`
}

type Bucket struct {
	Tasks     []Task
	NTasks    int
	NComplete int
	NFailed   int
	NPending  int
	State     string
}

type WorkflowData struct {
	Workflow  Workflow
	SCM       Bucket
	SRPM      Bucket
	JPB       Bucket
	Bootstrap Bucket
	Rebuild   Bucket
	Validate  Bucket
}

var Template = template.Must(template.ParseGlob("*.html"))

func count(b *Bucket) {
	for _, task := range b.Tasks {
		b.NTasks += 1
		if task.Result == nil {
			b.NPending += 1
		} else if task.Result.Outcome == "SUCCESS" {
			b.NComplete += 1
		} else {
			b.NFailed += 1
		}
	}
	if b.NFailed > 0 {
		b.State = "Failed"
	} else if b.NPending > 0 {
		b.State = "Pending"
	} else {
		b.State = "Succeeded"
	}
}

func get_data() WorkflowData {
	// url := os.Getenv("WORKFLOW_XML")
	// if url == "" {
	// 	url = "https://mbi-artifacts.s3.eu-central-1.amazonaws.com/1653bbcc-1ae3-4eaa-949a-239a24cf8de9/workflow.xml"
	// }
	// resp, err := http.Get(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if resp.StatusCode != http.StatusOK {
	// 	resp.Body.Close()
	// 	log.Fatal(fmt.Errorf("HTTP GET failed: %s", resp.Status))
	// }
	// bytes, err := ioutil.ReadAll(resp.Body)

	//bytes, err := os.ReadFile("/home/kojan/git/mbici-web/workflow.xml")
	bytes, err := os.ReadFile("/home/kojan/git/mbici-local/test/workflow.xml")
	//bytes, err := os.ReadFile("/home/kojan/git/mbici-fedora-pr/test/workflow.xml")
	if err != nil {
		log.Fatal(err)
	}
	var workflow Workflow
	if err := xml.Unmarshal(bytes, &workflow); err != nil {
		//resp.Body.Close()
		log.Fatal(err)
	}

	for j, result := range workflow.Results {
		for i, task := range workflow.Tasks {
			if result.TaskId == task.Id {
				workflow.Tasks[i].Result = &workflow.Results[j]
			}
		}
	}

	var data WorkflowData
	data.Workflow = workflow

	for _, task := range workflow.Tasks {
		if strings.HasSuffix(task.Id, "-checkout") {
			data.SCM.Tasks = append(data.SCM.Tasks, task)
		} else if strings.HasSuffix(task.Id, "-srpm") || task.Id == "platform" {
			data.SRPM.Tasks = append(data.SRPM.Tasks, task)
		} else if strings.HasSuffix(task.Id, "-rpm") || strings.HasSuffix(task.Id, "-repo") {
			if strings.HasSuffix(task.Id, "-b2-rpm") || task.Id == "b2-repo" {
				data.JPB.Tasks = append(data.JPB.Tasks, task)
			} else if strings.HasSuffix(task.Id, "-b3-rpm") || task.Id == "b3-repo" {
				data.Bootstrap.Tasks = append(data.Bootstrap.Tasks, task)
			} else {
				data.Rebuild.Tasks = append(data.Rebuild.Tasks, task)
			}
		} else if strings.HasSuffix(task.Id, "-validate") {
			data.Validate.Tasks = append(data.Validate.Tasks, task)
		} else {
			log.Fatal(fmt.Sprintf("Unmatched task ID: %s", task.Id))
		}
	}

	count(&data.SCM)
	count(&data.SRPM)
	count(&data.JPB)
	count(&data.Bootstrap)
	count(&data.Rebuild)
	count(&data.Validate)

	return data
}

func workflow_handler(w http.ResponseWriter, r *http.Request) {
	data := get_data()
	w.Header().Add("Content-Type", "text/html")
	err := Template.ExecuteTemplate(w, "workflow.html", data)
	if err != nil {
		fmt.Println(err)
	}
}

func task_handler(w http.ResponseWriter, r *http.Request) {
	task_id := strings.TrimPrefix(r.URL.Path, "/task/")
	wf_data := get_data()
	var task Task
	found := false
	for _, t := range wf_data.Workflow.Tasks {
		if t.Id == task_id {
			task = t
			found = true
		}
	}
	if !found {
		http.NotFound(w, r)
		return
	}
	w.Header().Add("Content-Type", "text/html")
	err := Template.ExecuteTemplate(w, "task.html", task)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/task/", task_handler)
	http.HandleFunc("/", workflow_handler)
	http.ListenAndServe(":8080", nil)
}
