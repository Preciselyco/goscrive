package scrive_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	scrive "github.com/Preciselyco/goscrive"
)

func TestNewDocument(t *testing.T) {
	cli := getClient()
	doc, se := cli.NewDocument(scrive.NewDocumentParams{})
	failIfScriveE(t, "TestPrepare: cli.NewDocument", se)
	fmt.Printf("got document response: %s\n", marshalIndentFail(t, "TestPrepare: document response", doc))
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
	fmt.Printf("got document response: %s\n", marshalIndentFail(t, "TestPrepareWithDoc: document response", doc))
}
