package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

type ScenarioCost struct {
	ScenarioLine   int     `json:"scenario_line"`
	TaskGroupName  string  `json:"task_group_name"`
	MaxRuntimeSecs float64 `json:"max_runtime_secs"`
}

func TestGet(t *testing.T) {
	body, err := Get("https://localhost/", "user", "pw")
	if err != nil {
		t.Error("get failed.")
	}
	defer body.Close()
	content, _ := ioutil.ReadAll(body)
	log.Println(string(content))
}

func TestPost(t *testing.T) {
	s1 := ScenarioCost{
		ScenarioLine:   1,
		TaskGroupName:  "task1",
		MaxRuntimeSecs: 1,
	}
	s2 := ScenarioCost{
		ScenarioLine:   2,
		TaskGroupName:  "task2",
		MaxRuntimeSecs: 2,
	}
	scenarioCosts := []ScenarioCost{s1, s2}
	fmt.Println(scenarioCosts)
	payload, err := json.Marshal(scenarioCosts)
	if err != nil {
		t.Error("obj to json failed.")
		fmt.Println(err.Error())
	}
	fmt.Println(string(payload))

	body, _ := Post(
		"https://localhost/", payload, "user", "pw")
	defer body.Close()
	content, _ := ioutil.ReadAll(body)
	fmt.Println(string(content))
}
