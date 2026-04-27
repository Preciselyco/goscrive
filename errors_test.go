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
