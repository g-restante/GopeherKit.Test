package internal

import (
	"os"
	"path/filepath"
	"testing"
)

// TestNewGenerator tests the generator creation.
func TestNewGenerator(t *testing.T) {
	gen := NewGenerator("testpkg", "/tmp/test")
	
	if gen.PackageName != "testpkg" {
		t.Errorf("Expected PackageName 'testpkg', got '%s'", gen.PackageName)
	}
	
	if gen.OutputDir != "/tmp/test" {
		t.Errorf("Expected OutputDir '/tmp/test', got '%s'", gen.OutputDir)
	}
	
	if gen.Templates == nil {
		t.Error("Expected Templates to be initialized")
	}
}

// TestGenerateTestBoilerplate tests test boilerplate generation.
func TestGenerateTestBoilerplate(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "gopherkit_test_")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)
	
	gen := NewGenerator("testpkg", tempDir)
	
	err = gen.GenerateTestBoilerplate("mypackage")
	if err != nil {
		t.Fatalf("Failed to generate test boilerplate: %v", err)
	}
	
	// Check if the file was created
	expectedPath := filepath.Join(tempDir, "mypackage_test.go")
	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		t.Errorf("Expected file %s was not created", expectedPath)
	}
	
	// Read and verify content
	content, err := os.ReadFile(expectedPath)
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}
	
	contentStr := string(content)
	if !contains(contentStr, "package mypackage_test") {
		t.Error("Generated file should contain correct package declaration")
	}
	
	if !contains(contentStr, "func TestMypackage(t *testing.T)") {
		t.Error("Generated file should contain test function")
	}
}

// TestGenerateAssertions tests custom assertion generation.
func TestGenerateAssertions(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "gopherkit_test_")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)
	
	gen := NewGenerator("assert", tempDir)
	
	specs := []string{
		"IsPositive:value int:value > 0:expected positive value",
		"IsEmpty:s string:len(s) == 0:expected empty string",
	}
	
	err = gen.GenerateAssertions(specs)
	if err != nil {
		t.Fatalf("Failed to generate assertions: %v", err)
	}
	
	// Check if the file was created
	expectedPath := filepath.Join(tempDir, "custom_assertions.go")
	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		t.Errorf("Expected file %s was not created", expectedPath)
	}
	
	// Read and verify content
	content, err := os.ReadFile(expectedPath)
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}
	
	contentStr := string(content)
	if !contains(contentStr, "func IsPositive(t *testing.T, value int, msg ...string)") {
		t.Error("Generated file should contain IsPositive assertion")
	}
	
	if !contains(contentStr, "func IsEmpty(t *testing.T, s string, msg ...string)") {
		t.Error("Generated file should contain IsEmpty assertion")
	}
}

// TestParseAssertionSpec tests assertion spec parsing.
func TestParseAssertionSpec(t *testing.T) {
	gen := NewGenerator("test", "/tmp")
	
	spec := "IsPositive:value int:value > 0:expected positive value"
	result, err := gen.parseAssertionSpec(spec)
	
	if err != nil {
		t.Fatalf("Failed to parse assertion spec: %v", err)
	}
	
	if result.Name != "IsPositive" {
		t.Errorf("Expected Name 'IsPositive', got '%s'", result.Name)
	}
	
	if result.Params != "value int" {
		t.Errorf("Expected Params 'value int', got '%s'", result.Params)
	}
	
	if result.Condition != "value > 0" {
		t.Errorf("Expected Condition 'value > 0', got '%s'", result.Condition)
	}
	
	if result.DefaultMessage != "expected positive value" {
		t.Errorf("Expected DefaultMessage 'expected positive value', got '%s'", result.DefaultMessage)
	}
}

// TestParseAssertionSpecInvalidFormat tests error handling for invalid spec format.
func TestParseAssertionSpecInvalidFormat(t *testing.T) {
	gen := NewGenerator("test", "/tmp")
	
	invalidSpec := "invalid:spec"
	_, err := gen.parseAssertionSpec(invalidSpec)
	
	if err == nil {
		t.Error("Expected error for invalid spec format")
	}
}

// contains checks if a string contains a substring.
func contains(haystack, needle string) bool {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return true
		}
	}
	return false
}
