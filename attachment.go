package scrive

import (
	"fmt"
	"net/http"
)

type AttachmentListFilter = map[string]string

type ListAttachmentsParams struct {
	Domain  *string
	Filter  []*AttachmentListFilter
	Sorting []*ListSortParam
}

func (c *Client) ListAttachments(p ListAttachmentsParams) (*[]*Attachment, *ScriveError) {
	resp := &[]*Attachment{}
	_, se := c.gw(
		"attachments/list",
		func(req *request) {
			req.writeMString("domain", p.Domain)
			req.writeMJSON("filter", p.Filter)
			req.writeMJSON("sorting", p.Sorting)
		},
		resp,
	)
	return resp, se
}

func (c *Client) DownloadAttachment(attachmentID string) ([]byte, string, *ScriveError) {
	resp, se := c.gwb(
		fmt.Sprintf("attachments/%s/download/file", attachmentID),
		nil,
	)
	if se != nil {
		return nil, "", se
	}
	return resp.body, resp.getHeader("content-type"), nil
}

type CreateAttachmentParams struct {
	Title    string
	FileName string
	File     []byte
}

func (c *Client) CreateAttachment(p CreateAttachmentParams) *ScriveError {
	_, se := c.pwbe(
		"attachments/create",
		func(req *request) {
			req.writeMString("title", &p.Title)
			req.writeMFile("file", p.FileName, p.File)
		},
		http.StatusCreated,
	)
	return se
}

func (c *Client) DeleteAttachments(attachmentIDs []string) *ScriveError {
	_, se := c.pwb(
		"attachments/delete",
		func(req *request) {
			req.writeMJSON("attachment_ids", attachmentIDs)
		},
	)
	return se
}

func (c *Client) SetAttachmentsSharing(attachmentIDs []string, shared bool) *ScriveError {
	_, se := c.pwb(
		"attachments/setsharing",
		func(req *request) {
			req.writeMJSON("attachment_ids", attachmentIDs)
			req.writeMBool("shared", &shared)
		},
	)
	return se
}
