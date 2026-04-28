package scrive

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

// stubScrive returns a *Client wired to a httptest.Server whose handler is
// supplied by the test. The handler can record incoming requests and respond
// with whatever body/status the test wants to assert against.
func stubScrive(t *testing.T, handler http.HandlerFunc) (*Client, *httptest.Server) {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	root := srv.URL
	cli, err := NewClient(Config{
		APIRoot: &root,
		PAC: &PAC{
			ClientCredentialsIdentifier: "cci",
			ClientCredentialsSecret:     "ccs",
			TokenCredentialsIdentifier:  "tci",
			TokenCredentialsSecret:      "tcs",
		},
		HTTPClient: srv.Client(),
	})
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	return cli, srv
}

// recordedRequest is what the stub captures from each incoming request so the
// test can assert on it without re-reading the body.
type recordedRequest struct {
	Method      string
	Path        string
	Auth        string
	ContentType string
	Body        []byte
}

func recordingHandler(captured *[]recordedRequest, mu *sync.Mutex, respond func(w http.ResponseWriter, r *http.Request, body []byte)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		mu.Lock()
		*captured = append(*captured, recordedRequest{
			Method:      r.Method,
			Path:        r.URL.Path,
			Auth:        r.Header.Get("Authorization"),
			ContentType: r.Header.Get("Content-Type"),
			Body:        body,
		})
		mu.Unlock()
		respond(w, r, body)
	}
}

func TestClient_NewDocument_Success(t *testing.T) {
	var (
		captured []recordedRequest
		mu       sync.Mutex
	)
	cli, _ := stubScrive(t, recordingHandler(&captured, &mu, func(w http.ResponseWriter, r *http.Request, body []byte) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"id":"doc-123","status":"preparation"}`))
	}))

	doc, se := cli.NewDocument(NewDocumentParams{
		FileName: "contract.pdf",
		File:     []byte("%PDF-fake"),
	})
	if se != nil {
		t.Fatalf("expected nil ScriveError, got %+v", se)
	}
	if doc.ID == nil || *doc.ID != "doc-123" {
		t.Fatalf("expected ID doc-123, got %v", doc.ID)
	}

	mu.Lock()
	defer mu.Unlock()
	if len(captured) != 1 {
		t.Fatalf("expected 1 request, got %d", len(captured))
	}
	got := captured[0]
	if got.Method != http.MethodPost {
		t.Errorf("expected POST, got %s", got.Method)
	}
	if got.Path != "/api/v2/documents/new" {
		t.Errorf("expected path /api/v2/documents/new, got %s", got.Path)
	}
	if !strings.HasPrefix(got.ContentType, "multipart/form-data") {
		t.Errorf("expected multipart Content-Type, got %s", got.ContentType)
	}
	if !strings.Contains(string(got.Body), "contract.pdf") {
		t.Errorf("expected file name in multipart body, got %s", got.Body)
	}
	if !strings.Contains(string(got.Body), "%PDF-fake") {
		t.Errorf("expected file bytes in multipart body, got %s", got.Body)
	}
	if got.Auth == "" {
		t.Errorf("expected Authorization header to be set")
	}
}

func TestClient_NewDocument_NonSuccess(t *testing.T) {
	t.Run("400 with valid Scrive envelope surfaces typed error", func(t *testing.T) {
		cli, _ := stubScrive(t, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{
				"error_type":    "request_parameters_invalid",
				"error_message": "The 'file' parameter is missing.",
				"http_code":     400
			}`))
		})

		doc, se := cli.NewDocument(NewDocumentParams{})
		if doc.ID != nil {
			t.Errorf("expected nil ID, got %v", *doc.ID)
		}
		if se == nil {
			t.Fatal("expected non-nil ScriveError")
		}
		if se.HttpCode != 400 {
			t.Errorf("expected HttpCode 400, got %d", se.HttpCode)
		}
		if se.ErrorType != "request_parameters_invalid" {
			t.Errorf("expected typed envelope, got %q", se.ErrorType)
		}
	})

	t.Run("401 plain text surfaces unparseable_response with real status", func(t *testing.T) {
		// Regression: end-to-end check that the plain-text 401 from
		// Scrive's auth layer (the production failure mode that motivated
		// the unparseableResponseError fallback) bubbles up cleanly.
		cli, _ := stubScrive(t, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("No valid access credentials were provided. Please refer to our API documentation."))
		})

		_, se := cli.NewDocument(NewDocumentParams{})
		if se == nil {
			t.Fatal("expected non-nil ScriveError")
		}
		if se.HttpCode != http.StatusUnauthorized {
			t.Errorf("expected HttpCode 401, got %d", se.HttpCode)
		}
		if se.ErrorType != ErrorTypeUnparseableResponse {
			t.Errorf("expected ErrorType %q, got %q", ErrorTypeUnparseableResponse, se.ErrorType)
		}
		if !strings.Contains(se.ErrorMessage, "No valid access credentials") {
			t.Errorf("expected body in ErrorMessage, got %q", se.ErrorMessage)
		}
	})

	t.Run("502 HTML surfaces unparseable_response with real status", func(t *testing.T) {
		cli, _ := stubScrive(t, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusBadGateway)
			_, _ = w.Write([]byte("<html><body>502 Bad Gateway</body></html>"))
		})

		_, se := cli.NewDocument(NewDocumentParams{})
		if se == nil {
			t.Fatal("expected non-nil ScriveError")
		}
		if se.HttpCode != http.StatusBadGateway {
			t.Errorf("expected HttpCode 502, got %d", se.HttpCode)
		}
		if se.ErrorType != ErrorTypeUnparseableResponse {
			t.Errorf("expected ErrorType %q, got %q", ErrorTypeUnparseableResponse, se.ErrorType)
		}
	})
}

func TestClient_UpdateDocument_Success(t *testing.T) {
	var (
		captured []recordedRequest
		mu       sync.Mutex
	)
	cli, _ := stubScrive(t, recordingHandler(&captured, &mu, func(w http.ResponseWriter, r *http.Request, body []byte) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"doc-123","status":"preparation"}`))
	}))

	id := "doc-123"
	doc := &Document{ID: &id, Title: String("My Contract")}
	got, se := cli.UpdateDocument(UpdateDocumentParams{
		DocumentID: "doc-123",
		Document:   doc,
	})
	if se != nil {
		t.Fatalf("expected nil ScriveError, got %+v", se)
	}
	if got.ID == nil || *got.ID != "doc-123" {
		t.Fatalf("expected ID doc-123, got %v", got.ID)
	}

	mu.Lock()
	defer mu.Unlock()
	if len(captured) != 1 {
		t.Fatalf("expected 1 request, got %d", len(captured))
	}
	rec := captured[0]
	if rec.Path != "/api/v2/documents/doc-123/update" {
		t.Errorf("expected path /api/v2/documents/doc-123/update, got %s", rec.Path)
	}
	// The document JSON is sent as a multipart form field. We don't roundtrip
	// the entire Document struct here; the contract is just that the title
	// we set ends up on the wire.
	if !strings.Contains(string(rec.Body), "My Contract") {
		t.Errorf("expected document JSON to include title, got %s", rec.Body)
	}
}

func TestClient_composeURL(t *testing.T) {
	mkClient := func(root string) *Client {
		c, err := NewClient(Config{
			APIRoot: &root,
			PAC:     &PAC{ClientCredentialsIdentifier: "x"},
		})
		if err != nil {
			t.Fatalf("NewClient: %v", err)
		}
		return c
	}

	t.Run("bare hostname implies https", func(t *testing.T) {
		c := mkClient("api-scrive.com")
		got := c.composeURL("documents/new")
		want := "https://api-scrive.com/api/v2/documents/new"
		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})

	t.Run("explicit http scheme is honored", func(t *testing.T) {
		c := mkClient("http://localhost:1234")
		got := c.composeURL("documents/new")
		want := "http://localhost:1234/api/v2/documents/new"
		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})

	t.Run("explicit https scheme is honored", func(t *testing.T) {
		c := mkClient("https://api-scrive.com")
		got := c.composeURL("documents/new")
		want := "https://api-scrive.com/api/v2/documents/new"
		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})

	t.Run("trailing slash on APIRoot is stripped", func(t *testing.T) {
		c := mkClient("http://localhost:1234/")
		got := c.composeURL("documents/new")
		want := "http://localhost:1234/api/v2/documents/new"
		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})
}
