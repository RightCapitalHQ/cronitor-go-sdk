package cronitor

import (
	"os"
	"testing"
)

func getClient() Cronitor {
	return Cronitor{
		ApiKey: os.Getenv("CRONITOR_API_KEY"),
	}
}

func Test_PutMonitor(t *testing.T) {
	var c = getClient()

	err := c.PutMonitors([]Monitor{
		{
			Key:      "hello_world",
			Name:     "Hello World",
			Type:     "job",
			Schedule: "* * * * * *",
			Notify: []string{
				"webhook:https://google.com",
			},
			Metadata: `{"foo": {"aaa": "bbb"}}`,
		},
	})

	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_GetMonitor(t *testing.T) {
	var c = getClient()

	_, err := c.GetMonitor("hello_world")

	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_DeleteMonitor(t *testing.T) {
	var c = getClient()

	err := c.DeleteMonitor("hello_world")

	if err != nil {
		t.Errorf(err.Error())
	}
}
