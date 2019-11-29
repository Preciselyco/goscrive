package scrive

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type NewDocumentParams struct {
	FileName string
	File     []byte
	Saved    *bool
}

func (c *Client) NewDocument(p NewDocumentParams) (*Document, *ScriveError) {
	if len(p.File) == 0 && p.Saved == nil {
		out := &Document{}
		if se := c.postExpect(
			"documents/new",
			nil,
			nil,
			http.StatusCreated,
			out); se != nil {
			return nil, se
		}
		return out, nil
	}
	var buf bytes.Buffer
	wr := multipart.NewWriter(&buf)
	if len(p.File) > 0 {
		fw, err := wr.CreateFormFile("file", p.FileName)
		if err != nil {
			return nil, localError(err)
		}
		if _, err := io.Copy(fw, bytes.NewReader(p.File)); err != nil {
			return nil, localError(err)
		}
	}
	if p.Saved != nil {
		if err := wr.WriteField("saved", boolToStr(*p.Saved)); err != nil {
			return nil, localError(err)
		}
	}
	wr.Close()
	headers := map[string]string{
		"Content-Type": wr.FormDataContentType(),
	}
	doc := &Document{}
	if se := c.postExpect(
		"documents/new",
		ioutil.NopCloser(
			bytes.NewReader(buf.Bytes()),
		),
		&headers,
		http.StatusCreated,
		doc); se != nil {
		return nil, se
	}
	return doc, nil
}

type NewDocumentFromTemplateParams struct {
	DocumentID    uint64
	ObjectVersion *uint64
}

func (c *Client) NewDocumentFromTemplate(p *NewDocumentFromTemplateParams) (*Document, *ScriveError) {
	doc := &Document{}
	url := fmt.Sprintf("documents/newfromtemplate/%d", p.DocumentID)
	if p.ObjectVersion == nil {
		if se := c.postExpect(
			url,
			nil,
			nil,
			http.StatusCreated,
			doc); se != nil {
			return nil, se
		}
		return doc, nil
	}
	var buf bytes.Buffer
	wr := multipart.NewWriter(&buf)
	if err := wr.WriteField("object_version", fmt.Sprintf("%d", *p.ObjectVersion)); err != nil {
		return nil, localError(err)
	}
	wr.Close()
	headers := map[string]string{
		"Content-Type": wr.FormDataContentType(),
	}
	if se := c.postExpect(
		url,
		ioutil.NopCloser(
			bytes.NewReader(buf.Bytes()),
		),
		&headers,
		http.StatusCreated,
		doc); se != nil {
		return nil, se
	}
	return doc, nil
}
