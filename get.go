package scrive

import (
	"fmt"
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

func (c *Client) GetMainFile(documentID string) ([]byte, *ScriveError) {
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
