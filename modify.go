package scrive

import (
	"fmt"
	"net/http"
)

type RemindSignatoriesParams struct {
	DocumentID    string
	ObjectVersion *uint64
}

func (c *Client) RemindSignatories(p RemindSignatoriesParams) *ScriveError {
	_, se := c.pwb(
		fmt.Sprintf("documents/%s/remind", p.DocumentID),
		func(req *request) {
			req.writeMUInt("object_version", p.ObjectVersion)
		},
	)
	return se
}

type ProlongDocumentParams struct {
	DocumentID    string
	Days          *uint64
	ObjectVersion *uint64
}

func (c *Client) ProlongDocument(p ProlongDocumentParams) (*Document, *ScriveError) {
	doc := &Document{}
	_, se := c.pw(
		fmt.Sprintf("documents/%s/prolong", p.DocumentID),
		func(req *request) {
			req.writeMUInt("days", p.Days)
			req.writeMUInt("object_version", p.ObjectVersion)
		},
		doc,
	)
	return doc, se
}

type CancelDocumentParams struct {
	DocumentID    string
	ObjectVersion *uint64
}

func (c *Client) CancelDocument(p CancelDocumentParams) (*Document, *ScriveError) {
	doc := &Document{}
	_, se := c.pw(
		fmt.Sprintf("documents/%s/cancel", p.DocumentID),
		func(req *request) {
			req.writeMUInt("object_version", p.ObjectVersion)
		},
		doc,
	)
	return doc, se
}

type TrashDocumentParams struct {
	DocumentID    string
	ObjectVersion *uint64
}

func (c *Client) TrashDocument(p TrashDocumentParams) (*Document, *ScriveError) {
	doc := &Document{}
	_, se := c.pw(
		fmt.Sprintf("documents/%s/trash", p.DocumentID),
		func(req *request) {
			req.writeMUInt("object_version", p.ObjectVersion)
		},
		doc,
	)
	return doc, se
}

type TrashDocumentsResponse struct {
	Trashed   uint64      `json:"trashed"`
	Documents []*Document `json:"documents"`
}

func (c *Client) TrashDocuments(documentIDs []string) (*TrashDocumentsResponse, *ScriveError) {
	resp := &TrashDocumentsResponse{}
	_, se := c.pw(
		"documents/trash",
		func(req *request) {
			req.writeMJSON("document_ids", documentIDs)
		},
		resp,
	)
	return resp, se
}

type DeleteDocumentParams struct {
	DocumentID    string
	ObjectVersion *uint64
}

func (c *Client) DeleteDocument(p DeleteDocumentParams) (*Document, *ScriveError) {
	doc := &Document{}
	_, se := c.pw(
		fmt.Sprintf("documents/%s/delete", p.DocumentID),
		func(req *request) {
			req.writeMUInt("object_version", p.ObjectVersion)
		},
		doc,
	)
	return doc, se
}

type DeleteDocumentsResponse struct {
	Deleted   uint64      `json:"deleted"`
	Documents []*Document `json:"documents"`
}

func (c *Client) DeleteDocuments(documentIDs []string) (*DeleteDocumentsResponse, *ScriveError) {
	resp := &DeleteDocumentsResponse{}
	_, se := c.pw(
		"documents/delete",
		func(req *request) {
			req.writeMJSON("document_ids", documentIDs)
		},
		resp,
	)
	return resp, se
}

type ForwardDocumentParams struct {
	DocumentID    string
	Email         string
	NoContent     *bool
	NoAttachments *bool
	ObjectVersion *uint64
}

func (c *Client) ForwardDocument(p ForwardDocumentParams) *ScriveError {
	_, se := c.pwb(
		fmt.Sprintf("documents/%s/forward", p.DocumentID),
		func(req *request) {
			req.writeMString("email", &p.Email)
			req.writeMBool("no_content", p.NoContent)
			req.writeMBool("no_attachments", p.NoAttachments)
			req.writeMUInt("object_version", p.ObjectVersion)
		},
	)
	return se
}

type AutoReminderParams struct {
	DocumentID    string
	Days          *uint64
	ObjectVersion *uint64
}

func (c *Client) SetDocumentAutoReminder(p AutoReminderParams) (*Document, *ScriveError) {
	doc := &Document{}
	_, se := c.pw(
		fmt.Sprintf("documents/%s/setautoreminder", p.DocumentID),
		func(req *request) {
			req.writeMUInt("days", p.Days)
			req.writeMUInt("object_version", p.ObjectVersion)
		},
		doc,
	)
	return doc, se
}

type RestartDocumentParams struct {
	DocumentID    string
	ObjectVersion *uint64
}

func (c *Client) RestartDocument(p RestartDocumentParams) (*Document, *ScriveError) {
	doc := &Document{}
	_, se := c.pwe(
		fmt.Sprintf("documents/%s/restart", p.DocumentID),
		func(req *request) {
			req.writeMUInt("object_version", p.ObjectVersion)
		},
		http.StatusCreated,
		doc,
	)
	return doc, se
}

type SignatoryAuthenticationToViewParams struct {
	DocumentID           string
	SignatoryID          string
	AuthenticationMethod AuthenticationMethodToView
	PersonalNumber       *string
	MobileNumber         *string
	ObjectVersion        *uint64
}

func (c *Client) SetSignatoryAuthenticationToView(p SignatoryAuthenticationToViewParams) (*Document, *ScriveError) {
	doc := &Document{}
	_, se := c.pw(
		fmt.Sprintf("documents/%s/%s/setauthenticationtoview", p.DocumentID, p.SignatoryID),
		func(req *request) {
			req.writeMString("authentication_type", &p.AuthenticationMethod)
			req.writeMString("personal_number", p.PersonalNumber)
			req.writeMString("mobile_number", p.MobileNumber)
			req.writeMUInt("object_version", p.ObjectVersion)
		},
		doc,
	)
	return doc, se
}

type SignatoryAuthenticationToSignParams struct {
	DocumentID           string
	SignatoryID          string
	AuthenticationMethod AuthenticationMethodToSign
	PersonalNumber       *string
	MobileNumber         *string
	ObjectVersion        *uint64
}

func (c *Client) SetSignatoryAuthenticationToSign(p SignatoryAuthenticationToSignParams) (*Document, *ScriveError) {
	doc := &Document{}
	_, se := c.pw(
		fmt.Sprintf("documents/%s/%s/setauthenticationtosign", p.DocumentID, p.SignatoryID),
		func(req *request) {
			req.writeMString("authentication_type", &p.AuthenticationMethod)
			req.writeMString("personal_number", p.PersonalNumber)
			req.writeMString("mobile_number", p.MobileNumber)
			req.writeMUInt("object_version", p.ObjectVersion)
		},
		doc,
	)
	return doc, se
}

type ChangeSignatoryEmailAndMobileParams struct {
	DocumentID    string
	SignatoryID   string
	Email         *string
	MobileNumber  *string
	ObjectVersion *uint64
}

func (c *Client) ChangeSignatoryEmailAndMobile(p ChangeSignatoryEmailAndMobileParams) (*Document, *ScriveError) {
	doc := &Document{}
	_, se := c.pw(
		fmt.Sprintf("documents/%s/%s/changeemailandmobile", p.DocumentID, p.SignatoryID),
		func(req *request) {
			req.writeMString("email", p.Email)
			req.writeMString("mobile_number", p.MobileNumber)
			req.writeMUInt("object_version", p.ObjectVersion)
		},
		doc,
	)
	return doc, se
}
