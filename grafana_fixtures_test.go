package gapi

const (
	createdDataSourceJSON = `{"id":1,"message":"Datasource added", "name": "test_datasource"}`
	annotationsJSON       = `[
			{
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
			}
	]`
	newAnnotationJSON = `{
			"message":"Annotation added",
			"id": 1,
			"endId": 2
	}`
	newGraphiteAnnotationJSON = `{
			"message":"Annotation added",
			"id": 1
	}`
	deleteAnnotationJSON = `{"message":"Annotation deleted"}`
)
