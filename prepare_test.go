package scrive_test

import (
	"io/ioutil"
	"os"
	"testing"

	scrive "github.com/Preciselyco/goscrive"
	"github.com/stretchr/testify/assert"
)

func TestNewDocument(t *testing.T) {
	cli := getClient()
	doc, se := cli.NewDocument(scrive.NewDocumentParams{})
	failIfScriveE(t, "TestPrepare: cli.NewDocument", se)
	log("got document response: %s", marshalIndentFail(t, "TestNewDocument: document response", doc))
}

func TestNewDocumentWithDoc(t *testing.T) {
	cli := getClient()
	documentBody, err := ioutil.ReadFile("testdata/document.pdf")
	failIfE(t, "TestNewDocumentWithDoc: readFile", err)
	saved := true
	doc, se := cli.NewDocument(scrive.NewDocumentParams{
		FileName: "document.pdf",
		File:     documentBody,
		Saved:    &saved,
	})
	failIfScriveE(t, "TestNewDocumentWithDoc: cli.NewDocument", se)
	log("got document response: %s", marshalIndentFail(t, "TestNewDocumentWithDoc: document response", doc))
}

func TestNewDocumentFromTemplate(t *testing.T) {
	cli := getClient()
	documentID := os.Getenv("TEST_TEMPLATE_ID")
	doc, se := cli.NewDocumentFromTemplate(scrive.NewDocumentFromTemplateParams{
		DocumentID: documentID,
	})
	failIfScriveE(t, "TestNewDocumentFromTemplate: cli.NewDocumentFromTemplate", se)
	log("got document response: %s", marshalIndentFail(t, "TestNewDocumentFromTemplate: document response", doc))
}

func TestDocumentClone(t *testing.T) {
	cli := getClient()
	documentID := os.Getenv("TEST_TEMPLATE_ID")
	doc, se := cli.CloneDocument(scrive.CloneDocumentParams{
		DocumentID: documentID,
	})
	failIfScriveE(t, "TestDocumentClone: cli.CloneDocument", se)
	log("got document response: %s", marshalIndentFail(t, "TestDocumentClone: document response", doc))
}

func TestDocumentUpdate(t *testing.T) {
	cli := getClient()
	documentID := os.Getenv("TEST_TEMPLATE_ID")
	newTitle := "some nice new title"
	doc, se := cli.UpdateDocument(scrive.UpdateDocumentParams{
		DocumentID: documentID,
		Document: &scrive.Document{
			Title: scrive.String(newTitle),
		},
	})
	failIfScriveE(t, "TestDocumentUpdate: cli.UpdateDocument", se)
	log("got document response: %s", marshalIndentFail(t, "TestDocumentUpdate: document response", doc))
	assert.Equal(t, newTitle, *doc.Title)
}

func TestSetAttachments(t *testing.T) {
	cli := getClient()
	documentID := os.Getenv("TEST_TEMPLATE_ID")
	attachment, err := ioutil.ReadFile("testdata/document.pdf")
	failIfE(t, "readFile", err)
	a1Name := "Attachment 1"
	a2Name := "Attachment 2"
	doc, se := cli.SetAttachments(scrive.SetAttachmentsParams{
		DocumentID: documentID,
		Attachments: []*scrive.SetAttachment{
			{
				Name:            "Attachment 1",
				Required:        true,
				AddToSealedFile: true,
				FileParam:       "attachment_1",
				File:            attachment,
			},
			{
				Name:            "Attachment 2",
				Required:        false,
				AddToSealedFile: false,
				FileParam:       "attachment_2",
				File:            attachment,
			},
		},
	})
	failIfScriveE(t, "TestSetAttachments: cli.SetAttachments", se)
	log("got document response: %s", marshalIndentFail(t, "TestSetAttachments: document response", doc))
	assert.NotNil(t, doc.AuthorAttachments)
	assert.Equal(t, len(*doc.AuthorAttachments), 2)
	a1 := (*doc.AuthorAttachments)[0]
	assert.Equal(t, a1.Name, a1Name)
	assert.Equal(t, a1.Required, true)
	assert.Equal(t, a1.AddToSealedFile, true)
	a2 := (*doc.AuthorAttachments)[1]
	assert.Equal(t, a2.Name, a2Name)
	assert.Equal(t, a2.Required, false)
	assert.Equal(t, a2.AddToSealedFile, false)
}
