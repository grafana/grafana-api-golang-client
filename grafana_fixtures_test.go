package gapi

const (
	createdDataSourceJSON = `{"id":1,"message":"Datasource added", "name": "test_datasource"}`

	getOrgsJSON = `[{"id":1,"name":"Main Org."},{"id":2,"name":"Test Org."}]`
	getOrgJSON = `{"id":1,"name":"Main Org.","address":{"address1":"","address2":"","city":"","zipCode":"","state":"","country":""}}`
	createdOrgJSON = `{"message":"Organization created","orgId":1}`
	updatedOrgJSON = `{"message":"Organization updated"}`
	deletedOrgJSON = `{"message":"Organization deleted"}`

)
