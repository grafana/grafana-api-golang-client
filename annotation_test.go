package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

func TestAnnotations(t *testing.T) {
	server, client := gapiTestTools(200, annotationsJSON)
	defer server.Close()

	params := map[string]string{
		"from":  "1506676478816",
		"to":    "1507281278816",
		"limit": "100",
	}
	as, err := client.Annotations(params)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(as))

	if as[0].ID != 1124 {
		t.Error("annotations response should contain annotations with an ID")
	}
}

func TestNewAnnotation(t *testing.T) {
	server, client := gapiTestTools(200, newAnnotationJSON)
	defer server.Close()

	a := Annotation{
		DashboardID: 123,
		PanelID:     456,
		Time:        1507037197339,
		IsRegion:    true,
		TimeEnd:     1507180805056,
		Tags:        []string{"tag1", "tag2"},
		Text:        "text description",
	}
	res, err := client.NewAnnotation(&a)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))

	if res != 1 {
		t.Error("new annotation response should contain the ID of the new annotation")
	}
}
