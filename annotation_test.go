package gapi

import (
	"testing"

	"github.com/gobs/pretty"
	"github.com/grafana/grafana-api-golang-client/goclient/client/annotations"
	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

const (
	annotationsJSON = `[{
		"id": 1124,
		"alertId": 0,
		"dashboardId": 468,
		"panelId": 2,
		"userId": 1,
		"userName": "",
		"newState": "",
		"prevState": "",
		"time": 1507266395000,
		"text": "test",
		"metric": "",
		"regionId": 1123,
		"type": "event",
		"tags": [
			"tag1",
			"tag2"
		],
		"data": {}
	}]`

	newAnnotationJSON = `{
		"message":"Annotation added",
		"id": 1,
		"endId": 2
	}`

	newGraphiteAnnotationJSON = `{
		"message":"Annotation added",
		"id": 1
	}`

	updateAnnotationJSON = `{"message":"Annotation updated"}`

	patchAnnotationJSON = `{"message":"Annotation patched"}`

	deleteAnnotationJSON = `{"message":"Annotation deleted"}`
)

func TestAnnotations(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, annotationsJSON)
	defer mocksrv.Close()

	from := int64(1506676478816)
	to := int64(1507281278816)
	limit := int64(100)

	resp, err := client.Annotations.GetAnnotations(
		annotations.NewGetAnnotationsParams().
			WithFrom(&from).
			WithTo(&to).
			WithLimit(&limit),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp.Payload))

	if resp.Payload[0].ID != 1124 {
		t.Error("annotations response should contain annotations with an ID")
	}
}

func TestNewAnnotation(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, newAnnotationJSON)
	defer mocksrv.Close()

	body := models.PostAnnotationsCmd{
		DashboardID: 123,
		PanelID:     456,
		Time:        1507037197339,
		//IsRegion:    true,
		TimeEnd: 1507180805056,
		Tags:    []string{"tag1", "tag2"},
		Text:    "text description",
	}
	res, err := client.Annotations.PostAnnotation(
		annotations.NewPostAnnotationParams().
			WithBody(&body),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(res))

	if *res.Payload.ID != 1 {
		t.Error("new annotation response should contain the ID of the new annotation")
	}
}

func TestUpdateAnnotation(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, updateAnnotationJSON)
	defer mocksrv.Close()

	body := models.UpdateAnnotationsCmd{
		Text: "new text description",
	}
	res, err := client.Annotations.UpdateAnnotation(
		annotations.NewUpdateAnnotationParams().
			WithAnnotationID("1").
			WithBody(&body),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(res))

	if res.Payload.Message != "Annotation updated" {
		t.Error("update annotation response should contain the correct response message")
	}
}

func TestPatchAnnotation(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, patchAnnotationJSON)
	defer mocksrv.Close()

	body := models.PatchAnnotationsCmd{
		Text: "new text description",
	}
	res, err := client.Annotations.PatchAnnotation(
		annotations.NewPatchAnnotationParams().
			WithAnnotationID("1").
			WithBody(&body),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(res.Payload))

	if res.Payload.Message != "Annotation patched" {
		t.Error("patch annotation response should contain the correct response message")
	}
}

func TestNewGraphiteAnnotation(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, newGraphiteAnnotationJSON)
	defer mocksrv.Close()

	a := models.PostGraphiteAnnotationsCmd{
		What: "what",
		When: 1507180805056,
		Tags: []string{"tag1", "tag2"},
		Data: "data",
	}
	res, err := client.Annotations.PostGraphiteAnnotation(
		annotations.NewPostGraphiteAnnotationParams().
			WithBody(&a),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(res.Payload))

	if *res.Payload.ID != 1 {
		t.Error("new annotation response should contain the ID of the new annotation")
	}
}

func TestDeleteAnnotation(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, deleteAnnotationJSON)
	defer mocksrv.Close()

	res, err := client.Annotations.DeleteAnnotationByID(
		annotations.NewDeleteAnnotationByIDParams().
			WithAnnotationID("1"),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(res.Payload))

	if res.Payload.Message != "Annotation deleted" {
		t.Error("delete annotation response should contain the correct response message")
	}
}

/*
This endpoint is not supported by Grafana API
func TestDeleteAnnotationByRegionID(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, deleteAnnotationJSON)
	defer mocksrv.Close()

	res, err := client.DeleteAnnotationByRegionID(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(res))

	if res != "Annotation deleted" {
		t.Error("delete annotation by region ID response should contain the correct response message")
	}
}
*/
