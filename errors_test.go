package scrive

import (
	"strings"
	"testing"
)

func TestUnparseableResponseError(t *testing.T) {
	t.Run("preserves the HTTP status code", func(t *testing.T) {
		se := unparseableResponseError(401, []byte("No valid access credentials were provided."))
		if se.HttpCode != 401 {
			t.Fatalf("expected HttpCode 401, got %d", se.HttpCode)
		}
	})

	t.Run("includes the body in the error message", func(t *testing.T) {
		se := unparseableResponseError(401, []byte("No valid access credentials were provided."))
		if !strings.Contains(se.ErrorMessage, "No valid access credentials") {
			t.Fatalf("expected body in ErrorMessage, got %q", se.ErrorMessage)
		}
	})

	t.Run("uses the unparseable_response error type", func(t *testing.T) {
		se := unparseableResponseError(404, []byte("Not Found"))
		if se.ErrorType != ErrorTypeUnparseableResponse {
			t.Fatalf("expected ErrorType %q, got %q", ErrorTypeUnparseableResponse, se.ErrorType)
		}
	})

	t.Run("trims whitespace from the body", func(t *testing.T) {
		se := unparseableResponseError(500, []byte("\n\n  Internal Server Error  \n"))
		if se.ErrorMessage != "Internal Server Error" {
			t.Fatalf("expected trimmed body, got %q", se.ErrorMessage)
		}
	})

	t.Run("truncates long bodies", func(t *testing.T) {
		body := strings.Repeat("x", maxUnparseableBodyLen+100)
		se := unparseableResponseError(500, []byte(body))
		if !strings.HasSuffix(se.ErrorMessage, "…") {
			t.Fatalf("expected truncated body to end with ellipsis, got %q", se.ErrorMessage)
		}
		if len(se.ErrorMessage) > maxUnparseableBodyLen+10 {
			t.Fatalf("expected truncated message length around %d, got %d", maxUnparseableBodyLen, len(se.ErrorMessage))
		}
	})

	t.Run("falls back to a placeholder for empty bodies", func(t *testing.T) {
		se := unparseableResponseError(502, nil)
		if !strings.Contains(se.ErrorMessage, "empty response body") {
			t.Fatalf("expected placeholder for empty body, got %q", se.ErrorMessage)
		}
	})
}

// TestClient_expect covers the four code paths through Client.expect, the
// function that routes between localError, parseResponseError, and
// unparseableResponseError. Same-package test so we can call the unexported
// method directly.
func TestClient_expect(t *testing.T) {
	c := &Client{}

	t.Run("success: returns nil when code matches and out is decoded", func(t *testing.T) {
		var out struct {
			Hello string `json:"hello"`
		}
		se := c.expect(200, []byte(`{"hello":"world"}`), 200, &out, false)
		if se != nil {
			t.Fatalf("expected nil ScriveError, got %+v", se)
		}
		if out.Hello != "world" {
			t.Fatalf("expected out.Hello=\"world\", got %q", out.Hello)
		}
	})

	t.Run("success with bin=true skips JSON parsing", func(t *testing.T) {
		se := c.expect(200, []byte("not json at all"), 200, nil, true)
		if se != nil {
			t.Fatalf("expected nil ScriveError for bin=true success, got %+v", se)
		}
	})

	t.Run("success but body fails to parse returns localError", func(t *testing.T) {
		var out struct {
			Hello string `json:"hello"`
		}
		se := c.expect(200, []byte("not json"), 200, &out, false)
		if se == nil {
			t.Fatal("expected non-nil ScriveError when success body fails to parse")
		}
		if se.ErrorType != ErrorTypeLocal {
			t.Fatalf("expected ErrorType %q, got %q", ErrorTypeLocal, se.ErrorType)
		}
		if se.HttpCode != -1 {
			t.Fatalf("expected HttpCode -1 for local parse failure, got %d", se.HttpCode)
		}
	})

	t.Run("error: parses well-formed Scrive error envelope", func(t *testing.T) {
		body := []byte(`{
			"error_type":    "request_parameters_invalid",
			"error_message": "The 'document' parameter is missing.",
			"http_code":     400
		}`)
		se := c.expect(400, body, 200, nil, false)
		if se == nil {
			t.Fatal("expected non-nil ScriveError for non-success status")
		}
		if se.HttpCode != 400 {
			t.Fatalf("expected HttpCode 400, got %d", se.HttpCode)
		}
		if se.ErrorType != "request_parameters_invalid" {
			t.Fatalf("expected ErrorType from JSON envelope, got %q", se.ErrorType)
		}
		if !strings.Contains(se.ErrorMessage, "missing") {
			t.Fatalf("expected envelope error message, got %q", se.ErrorMessage)
		}
	})

	t.Run("error with non-JSON body surfaces unparseable_response with real status", func(t *testing.T) {
		// Regression test for the previous behavior: when Scrive returned a
		// non-JSON 401 (e.g. plain-text "No valid access credentials..."),
		// expect() would localError() the JSON parse failure and drop the
		// HTTP status. Now it must surface httpCode + body.
		body := []byte("No valid access credentials were provided. Please refer to our API documentation.")
		se := c.expect(401, body, 200, nil, false)
		if se == nil {
			t.Fatal("expected non-nil ScriveError for non-JSON error body")
		}
		if se.HttpCode != 401 {
			t.Fatalf("expected real HttpCode 401, got %d", se.HttpCode)
		}
		if se.ErrorType != ErrorTypeUnparseableResponse {
			t.Fatalf("expected ErrorType %q, got %q", ErrorTypeUnparseableResponse, se.ErrorType)
		}
		if !strings.Contains(se.ErrorMessage, "No valid access credentials") {
			t.Fatalf("expected body in ErrorMessage, got %q", se.ErrorMessage)
		}
	})

	t.Run("error with HTML body still surfaces unparseable_response", func(t *testing.T) {
		body := []byte("<html><body>502 Bad Gateway</body></html>")
		se := c.expect(502, body, 200, nil, false)
		if se == nil {
			t.Fatal("expected non-nil ScriveError")
		}
		if se.HttpCode != 502 {
			t.Fatalf("expected HttpCode 502, got %d", se.HttpCode)
		}
		if se.ErrorType != ErrorTypeUnparseableResponse {
			t.Fatalf("expected ErrorType %q, got %q", ErrorTypeUnparseableResponse, se.ErrorType)
		}
	})

	t.Run("error with valid JSON but missing required envelope fields surfaces unparseable_response", func(t *testing.T) {
		// parseResponseError requires error_type, error_message, and http_code
		// to all be present. A JSON body that's syntactically valid but missing
		// any of those should fall back to unparseable_response so the real
		// HTTP status survives.
		body := []byte(`{"some_other_shape": "value"}`)
		se := c.expect(403, body, 200, nil, false)
		if se == nil {
			t.Fatal("expected non-nil ScriveError")
		}
		if se.HttpCode != 403 {
			t.Fatalf("expected HttpCode 403, got %d", se.HttpCode)
		}
		if se.ErrorType != ErrorTypeUnparseableResponse {
			t.Fatalf("expected ErrorType %q, got %q", ErrorTypeUnparseableResponse, se.ErrorType)
		}
	})

	t.Run("ScriveErrorToError formats unparseable_response usefully", func(t *testing.T) {
		body := []byte("No valid access credentials were provided.")
		se := c.expect(401, body, 200, nil, false)
		err := ScriveErrorToError(se)
		msg := err.Error()
		if !strings.Contains(msg, "401") || !strings.Contains(msg, "unparseable_response") || !strings.Contains(msg, "No valid access credentials") {
			t.Fatalf("expected formatted error to include status/type/body, got %q", msg)
		}
	})
}
