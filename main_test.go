package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Helper to test a route
func testRoute(t *testing.T, path string, expectedContent string) {
	req := httptest.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()

	// Simulate routing
	switch path {
	case "/":
		homeHandler(w, req)
	case "/course":
		courseHandler(w, req)
	case "/about":
		aboutHandler(w, req)
	case "/blog":
		blogHandler(w, req)
	default:
		http.NotFound(w, req)
	}

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK for %s, got %d", path, res.StatusCode)
	}

	body := w.Body.String()
	if !strings.Contains(body, expectedContent) {
		t.Errorf("Expected content %q in %s, but not found", expectedContent, path)
	}
}

func TestHomeRoute(t *testing.T) {
	testRoute(t, "/", "Learn and Grow")
}

func TestCourseRoute(t *testing.T) {
	testRoute(t, "/course", "Courses")
}

func TestAboutRoute(t *testing.T) {
	testRoute(t, "/about", "About Us")
}

func TestBlogRoute(t *testing.T) {
	testRoute(t, "/blog", "Tech Blog")
}
