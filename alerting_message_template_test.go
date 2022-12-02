package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

func TestMessageTemplates(t *testing.T) {
	t.Run("get message templates succeeds", func(t *testing.T) {
		client := gapiTestTools(t, 200, getMessageTemplatesJSON)

		ts, err := client.MessageTemplates()

		if err != nil {
			t.Error(err)
		}
		t.Log(pretty.PrettyFormat(ts))
		if len(ts) != 2 {
			t.Errorf("wrong number of templates returned, got %#v", ts)
		}
		if ts[0].Name != "template-one" {
			t.Errorf("incorrect name - expected %s on element %d, got %#v", "template-one", 0, ts)
		}
		if ts[1].Name != "template-two" {
			t.Errorf("incorrect name - expected %s on element %d, got %#v", "template-two", 0, ts)
		}
	})

	t.Run("get message template succeeds", func(t *testing.T) {
		client := gapiTestTools(t, 200, messageTemplateJSON)

		tmpl, err := client.MessageTemplate("template-one")

		if err != nil {
			t.Error(err)
		}
		t.Log(pretty.PrettyFormat(tmpl))
		if tmpl.Name != "template-one" {
			t.Errorf("incorrect name - expected %s, got %#v", "template-one", tmpl)
		}
	})

	t.Run("get non-existent message template fails", func(t *testing.T) {
		client := gapiTestTools(t, 404, ``)

		tmpl, err := client.MessageTemplate("does not exist")

		if err == nil {
			t.Errorf("expected error but got nil")
			t.Log(pretty.PrettyFormat(tmpl))
		}
	})

	t.Run("put message template succeeds", func(t *testing.T) {
		client := gapiTestTools(t, 202, messageTemplateJSON)

		err := client.SetMessageTemplate("template-three", "{{define \"template-one\" }}\n  content three\n{{ end }}")

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("delete message template succeeds", func(t *testing.T) {
		client := gapiTestTools(t, 204, ``)

		err := client.DeleteMessageTemplate("template-three")

		if err != nil {
			t.Error(err)
		}
	})
}

const getMessageTemplatesJSON = `
[
	{
		"name": "template-one",
		"template": "{{define \"template-one\" }}\n  content one\n{{ end }}"
	},
	{
		"name": "template-two",
		"template": "{{define \"template-one\" }}\n  content two\n{{ end }}"
	}
]
`

const messageTemplateJSON = `
{
	"name": "template-one",
	"template": "{{define \"template-one\" }}\n  content one\n{{ end }}"
}
`
