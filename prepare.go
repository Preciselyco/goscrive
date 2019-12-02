package scrive

import (
	"fmt"
	"net/http"
)

type NewDocumentParams struct {
	FileName string
	File     []byte
	Saved    *bool
}

func (c *Client) NewDocument(p NewDocumentParams) (*Document, *ScriveError) {
	doc := &Document{}
	_, se := c.pwe(
		"documents/new",
		func(req *request) {
			req.writeMFile("file", p.FileName, p.File)
			req.writeMBool("saved", p.Saved)
		},
		http.StatusCreated,
		doc,
	)
	return doc, se
}

type NewDocumentFromTemplateParams struct {
	DocumentID    string
	ObjectVersion *uint64
}

func (c *Client) NewDocumentFromTemplate(p NewDocumentFromTemplateParams) (*Document, *ScriveError) {
	doc := &Document{}
	_, se := c.pwe(
		fmt.Sprintf("documents/newfromtemplate/%s", p.DocumentID),
		func(req *request) {
			req.writeMUInt("object_version", p.ObjectVersion)
		},
		http.StatusCreated,
		doc,
	)
	return doc, se
}

type CloneDocumentParams struct {
	DocumentID    string
	ObjectVersion *uint64
}

func (c *Client) CloneDocument(p CloneDocumentParams) (*Document, *ScriveError) {
	doc := &Document{}
	_, se := c.pwe(
		fmt.Sprintf("documents/%s/clone", p.DocumentID),
		func(req *request) {
			req.writeMUInt("object_version", p.ObjectVersion)
		},
		http.StatusCreated,
		doc,
	)
	return doc, se
}

type UpdateDocumentParams struct {
	DocumentID    string
	Document      *Document
	ObjectVersion *uint64
}

func (c *Client) UpdateDocument(p UpdateDocumentParams) (*Document, *ScriveError) {
	doc := &Document{}
	_, se := c.pw(
		fmt.Sprintf("documents/%s/update", p.DocumentID),
		func(req *request) {
			req.writeMJSON("document", p.Document)
			req.writeMUInt("object_version", p.ObjectVersion)
		},
		doc,
	)
	return doc, se
}

type SetAttachment struct {
	Name            string `json:"name"`
	Required        bool   `json:"required"`
	AddToSealedFile bool   `json:"add_to_sealed_file"`
	FileParam       string `json:"file_param"`
	FileName        string `json:"-"`
	File            []byte `json:"-"`
}

type SetAttachmentsParams struct {
	DocumentID    string
	Attachments   []*SetAttachment
	Incremental   *bool
	ObjectVersion *uint64
}

func (c *Client) SetAttachments(p SetAttachmentsParams) (*Document, *ScriveError) {
	doc := &Document{}
	_, se := c.pw(
		fmt.Sprintf("documents/%s/setattachments", p.DocumentID),
		func(req *request) {
			req.writeMJSON("attachments", p.Attachments)
			req.writeMBool("incremental", p.Incremental)
			req.writeMUInt("object_version", p.ObjectVersion)
			for _, a := range p.Attachments {
				req.writeMFile(a.FileParam, a.FileName, a.File)
			}
		},
		doc,
	)
	return doc, se
}

type RemovePagesParams struct {
	DocumentID    string
	Pages         []uint64
	ObjectVersion *uint64
}

func (c *Client) RemovePages(p RemovePagesParams) (*Document, *ScriveError) {
	doc := &Document{}
	_, se := c.pw(
		fmt.Sprintf("documents/%s/removepages", p.DocumentID),
		func(req *request) {
			req.writeMJSON("pages", p.Pages)
			req.writeMUInt("object_version", p.ObjectVersion)
		},
		doc,
	)
	return doc, se
}

type StartParams struct {
	DocumentID    string
	ObjectVersion *uint64
}

func (c *Client) Start(p StartParams) (*Document, *ScriveError) {
	doc := &Document{}
	_, se := c.pw(
		fmt.Sprintf("documents/%s/start", p.DocumentID),
		func(req *request) {
			req.writeMUInt("object_version", p.ObjectVersion)
		},
		doc,
	)
	return doc, se
}
