package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

func TestNewDataSource(t *testing.T) {
	server, client := gapiTestTools(200, createdDataSourceJSON)
	defer server.Close()

	ds := &DataSource{
		Name:      "foo",
		Type:      "cloudwatch",
		URL:       "http://some-url.com",
		Access:    "access",
		IsDefault: true,
		JSONData: JSONData{
			AssumeRoleArn: "arn:aws:iam::123:role/some-role",
			AuthType:      "keys",
			DefaultRegion: "us-east-1",
		},
		SecureJSONData: SecureJSONData{
			AccessKey: "123",
			SecretKey: "456",
		},
	}

	created, err := client.NewDataSource(ds)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(created))

	if created != 1 {
		t.Error("datasource creation response should return the created datasource ID")
	}
}
