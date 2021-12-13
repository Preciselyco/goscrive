package scrive

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (c *Client) GetDocument(documentID string) (*Document, *ScriveError) {
	doc := &Document{}
	_, se := c.gw(
		fmt.Sprintf("documents/%s/get", documentID),
		nil,
		doc,
	)
	return doc, se
}

func (c *Client) GetDocumentByShortID(documentID string) (*Document, *ScriveError) {
	doc := &Document{}
	_, se := c.gw(
		fmt.Sprintf("documents/%s/getbyshortid", documentID),
		nil,
		doc,
	)
	return doc, se
}

func (c *Client) GetSignLinkQR(documentID, signatoryID string) ([]byte, string, *ScriveError) {
	resp, se := c.gwb(
		fmt.Sprintf("documents/%s/%s/getqrcode", documentID, signatoryID),
		nil,
	)
	if se != nil {
		return nil, "", se
	}
	return resp.body, resp.getHeader("content-type"), nil
}

type GetDocumentListParams struct {
	Offset  *uint64
	Max     *uint64
	Filter  []map[string]interface{}
	Sorting []map[string]string
}

type GetDocumentsListResponse struct {
	TotalMatching uint64      `json:"total_matching"`
	Documents     []*Document `json:"documents"`
}

func (c *Client) GetDocumentList(p GetDocumentListParams) (*GetDocumentsListResponse, *ScriveError) {
	resp := &GetDocumentsListResponse{}
	_, se := c.gw(
		"documents/list",
		func(req *request) {
			req.setQueryUint("offset", p.Offset)
			req.setQueryUint("max", p.Max)
			req.setQueryJSON("filter", p.Filter)
			req.setQueryJSON("sorting", p.Sorting)
		},
		resp,
	)
	return resp, se
}

type FileResponse struct {
	Length int64 // Length of file in bytes, iff reported by server `Content-Length` header. Otherwise -1
	Body   io.ReadCloser
}

func (c *Client) GetMainFile(documentID string) (*FileResponse, *ScriveError) {
	resp, err := c.getFile(fmt.Sprintf("documents/%s/files/main/file", documentID), nil)
	if err != nil {
		return nil, localError(err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, localError(fmt.Errorf("expected status code: 200, got: %d", resp.StatusCode))
	}

	lengthStr := resp.Header.Get("Content-Length")
	length, err := strconv.ParseInt(lengthStr, 10, 64)
	if err != nil {
		length = -1
	}

	return &FileResponse{
		Length: length,
		Body:   resp.Body,
	}, nil
}

func (c *Client) GetMainFileBytes(documentID string) ([]byte, *ScriveError) {
	resp, se := c.gwb(
		fmt.Sprintf("documents/%s/files/main/file", documentID),
		nil,
	)
	if se != nil {
		return nil, se
	}
	response := make([]byte, len(resp.body))
	copy(response, resp.body)
	return response, nil
}

func (c *Client) GetRelatedFile(documentID, fileID string) ([]byte, string, *ScriveError) {
	resp, se := c.gwb(
		fmt.Sprintf("documents/%s/files/%s/file", documentID, fileID),
		nil,
	)
	if se != nil {
		return nil, "", se
	}
	response := make([]byte, len(resp.body))
	copy(response, resp.body)
	return response, resp.getHeader("content-type"), nil
}

func (c *Client) GetDocumentHistory(documentID string) (*DocumentHistory, *ScriveError) {
	hist := &DocumentHistory{}
	_, se := c.gw(
		fmt.Sprintf("documents/%s/history", documentID),
		nil,
		hist,
	)
	return hist, se
}

type TriggerAPICallbackParams struct {
	DocumentID    string
	ObjectVersion *uint64
}

func (c *Client) TriggerAPICallback(p TriggerAPICallbackParams) *ScriveError {
	_, se := c.pwb(
		fmt.Sprintf("documents/%s/callback", p.DocumentID),
		func(req *request) {
			req.writeMUInt("object_version", p.ObjectVersion)
		},
	)
	return se
}
