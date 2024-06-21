package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	//"io/ioutil"
	"log"
	"mime"
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
	State        string
}

type Workflow struct {
	XMLName xml.Name `xml:"workflow"`
	Tasks   []Task   `xml:"task"`
	Results []Result `xml:"result"`
}

type Component struct {
	XMLName   xml.Name `xml:"component"`
	Name      string   `xml:"name"`
	SCM       string   `xml:"scm"`
	Commit    string   `xml:"commit"`
	Lookaside string   `xml:"lookaside"`
}

type Subject struct {
	XMLName    xml.Name    `xml:"subject"`
	Components []Component `xml:"component"`
}

type Macro struct {
	XMLName xml.Name `xml:"macro"`
	Name    string   `xml:"name"`
	Value   string   `xml:"value"`
}

type Phase struct {
	XMLName    xml.Name `xml:"phase"`
	Name       string   `xml:"name"`
	Macros     []Macro  `xml:"macro"`
	Components []string `xml:"component"`
}

type Plan struct {
	XMLName xml.Name `xml:"plan"`
	Macros  []Macro  `xml:"macro"`
	Phases  []Phase  `xml:"phase"`
}

type Bucket struct {
	Id        string
	Tasks     []*Task
	NTasks    int
	NComplete int
	NFailed   int
	NPending  int
	State     string
}

type PipelineData struct {
	Workflow  Workflow
	SCM       Bucket
	SRPM      Bucket
	JPB       Bucket
	Bootstrap Bucket
	Rebuild   Bucket
	Validate  Bucket
}

var Template = template.Must(template.ParseGlob("*.html"))

func count(id string, b *Bucket) {
	b.Id = id
	for _, task := range b.Tasks {
		b.NTasks += 1
		if task.State == "Pending" {
			b.NPending += 1
		} else if task.State == "Succeeded" {
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

func get_workflow() Workflow {
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

	for i, task := range workflow.Tasks {
		for j, result := range workflow.Results {
			if result.TaskId == task.Id {
				workflow.Tasks[i].Result = &workflow.Results[j]
			}
		}
		if workflow.Tasks[i].Result == nil {
			workflow.Tasks[i].State = "Pending"
		} else if workflow.Tasks[i].Result.Outcome == "SUCCESS" {
			workflow.Tasks[i].State = "Succeeded"
		} else {
			workflow.Tasks[i].State = "Failed"
		}
	}

	return workflow
}

func get_subject() Subject {
	bytes, err := os.ReadFile("/home/kojan/git/mbici-local/test/subject.xml")
	if err != nil {
		log.Fatal(err)
	}
	var subject Subject
	if err := xml.Unmarshal(bytes, &subject); err != nil {
		//resp.Body.Close()
		log.Fatal(err)
	}
	return subject
}

func get_plan() Plan {
	bytes, err := os.ReadFile("/home/kojan/git/mbici-config/plan/bootstrap-all-rawhide.xml")
	if err != nil {
		log.Fatal(err)
	}
	var plan Plan
	if err := xml.Unmarshal(bytes, &plan); err != nil {
		//resp.Body.Close()
		log.Fatal(err)
	}
	return plan
}

func pipeline_handler(w http.ResponseWriter, r *http.Request) {
	workflow := get_workflow()

	var data PipelineData
	data.Workflow = workflow

	for i, task := range workflow.Tasks {
		if strings.HasSuffix(task.Id, "-checkout") {
			data.SCM.Tasks = append(data.SCM.Tasks, &workflow.Tasks[i])
		} else if strings.HasSuffix(task.Id, "-srpm") || task.Id == "platform" || task.Id == "platform-repo" {
			data.SRPM.Tasks = append(data.SRPM.Tasks, &workflow.Tasks[i])
		} else if strings.HasSuffix(task.Id, "-rpm") || strings.HasSuffix(task.Id, "-repo") {
			if strings.HasSuffix(task.Id, "-b2-rpm") || task.Id == "b2-repo" {
				data.JPB.Tasks = append(data.JPB.Tasks, &workflow.Tasks[i])
			} else if strings.HasSuffix(task.Id, "-b3-rpm") || task.Id == "b3-repo" {
				data.Bootstrap.Tasks = append(data.Bootstrap.Tasks, &workflow.Tasks[i])
			} else {
				data.Rebuild.Tasks = append(data.Rebuild.Tasks, &workflow.Tasks[i])
			}
		} else if strings.HasSuffix(task.Id, "-validate") {
			data.Validate.Tasks = append(data.Validate.Tasks, &workflow.Tasks[i])
		} else {
			log.Fatal(fmt.Sprintf("Unmatched task ID: %s", task.Id))
		}
	}

	count("scm", &data.SCM)
	count("srpm", &data.SRPM)
	count("jpb", &data.JPB)
	count("bootstrap", &data.Bootstrap)
	count("rebuild", &data.Rebuild)
	count("validate", &data.Validate)

	w.Header().Add("Content-Type", "text/html")
	err := Template.ExecuteTemplate(w, "pipeline.html", data)
	if err != nil {
		fmt.Println(err)
	}
}

func workflow_handler(w http.ResponseWriter, r *http.Request) {
	workflow := get_workflow()
	w.Header().Add("Content-Type", "text/html")
	err := Template.ExecuteTemplate(w, "workflow.html", workflow)
	if err != nil {
		fmt.Println(err)
	}
}

func subject_handler(w http.ResponseWriter, r *http.Request) {
	subject := get_subject()
	w.Header().Add("Content-Type", "text/html")
	err := Template.ExecuteTemplate(w, "subject.html", subject)
	if err != nil {
		fmt.Println(err)
	}
}

func plan_handler(w http.ResponseWriter, r *http.Request) {
	plan := get_plan()
	w.Header().Add("Content-Type", "text/html")
	err := Template.ExecuteTemplate(w, "plan.html", plan)
	if err != nil {
		fmt.Println(err)
	}
}

func task_handler(w http.ResponseWriter, r *http.Request) {
	task_id := strings.TrimPrefix(r.URL.Path, "/task/")
	workflow := get_workflow()
	var task Task
	found := false
	for _, t := range workflow.Tasks {
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
	mime.AddExtensionType(".log", "text/plain")
	mime.AddExtensionType(".conf", "text/plain")
	mime.AddExtensionType(".spec", "text/plain")
	mime.AddExtensionType(".patch", "text/plain")
	mime.AddExtensionType(".fmf", "text/plain")
	mime.AddExtensionType(".yaml", "text/plain")
	http.Handle("/artifact/", http.StripPrefix("/artifact/", http.FileServer(http.Dir("/mnt/nfs/mbi-result/local/"))))

	http.HandleFunc("/task/", task_handler)
	http.HandleFunc("/workflow", workflow_handler)
	http.HandleFunc("/subject", subject_handler)
	http.HandleFunc("/plan", plan_handler)
	http.HandleFunc("/", pipeline_handler)
	http.ListenAndServe(":8080", nil)
}
