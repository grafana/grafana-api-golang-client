package gapi

const (
	createdDataSourceJSON = `{"id":1,"message":"Datasource added", "name": "test_datasource"}`

	getOrgsJSON = `[{"id":1,"name":"Main Org."},{"id":2,"name":"Test Org."}]`
	getOrgJSON = `{"id":1,"name":"Main Org.","address":{"address1":"","address2":"","city":"","zipCode":"","state":"","country":""}}`
	createdOrgJSON = `{"message":"Organization created","orgId":1}`
	updatedOrgJSON = `{"message":"Organization updated"}`
	deletedOrgJSON = `{"message":"Organization deleted"}`

	getOrgUsersJSON = `[{"orgId":1,"userId":1,"email":"admin@localhost","avatarUrl":"/avatar/46d229b033af06a191ff2267bca9ae56","login":"admin","role":"Admin","lastSeenAt":"2018-06-28T14:16:11Z","lastSeenAtAge":"\u003c 1m"}]`
	addOrgUserJSON = `{"message":"User added to organization"}`
	updateOrgUserJSON = `{"message":"Organization user updated"}`
	removeOrgUserJSON = `{"message":"User removed from organization"}`

	getUsersJSON = `[{"id":1,"name":"","login":"admin","email":"admin@localhost","avatarUrl":"/avatar/46d229b033af06a191ff2267bca9ae56","isAdmin":true,"lastSeenAt":"2018-06-28T14:42:24Z","lastSeenAtAge":"\u003c 1m"}]`
	getUserByEmailJSON = `{"id":1,"email":"admin@localhost","name":"","login":"admin","theme":"","orgId":1,"isGrafanaAdmin":true}`
)
