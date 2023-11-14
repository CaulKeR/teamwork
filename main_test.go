package teamwork

import (
	"os"
	"testing"
)

func TestSortEmailsWithOccurs(t *testing.T) {

	result := sortEmailsWithOccurs("customers_test.csv")
	expected := []DomainOccur{
		{"example.com", 2},
		{"example.org", 1},
	}
	if !equalDomainOccurs(result, expected) {
		t.Errorf(
			"Result does not match the expected output.\nExpected: %v\nActual: %v",
			expected,
			result)
	}

}

func equalDomainOccurs(a, b []DomainOccur) bool {

	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].domain != b[i].domain || a[i].occur != b[i].occur {
			return false
		}
	}
	return true

}

func TestOpenFile(t *testing.T) {

	// case one: open an existing file
	existingFile := "customers_test.csv"
	file, err := os.Open(existingFile)
	if err != nil {
		t.Fatal(err)
	}
	defer closeFile(file)
	openedFile := openFile(existingFile)
	if openedFile == nil {
		t.Error("Expected non-nil file, got nil")
	}

	// case two: open non-existing file
	nonExistingFile := "some_file.csv"
	openedFile = openFile(nonExistingFile)
	if openedFile != nil {
		t.Error("Expected nil file, got non-nil")
	}

}

func TestExtractEmailDomain(t *testing.T) {

	testCases := []struct {
		input    string
		expected string
	}{
		{
			"user@example.com", "example.com",
		},
		{
			"john.doe123@test.org", "test.org",
		},
		{
			"invalid_email", emptyString,
		},
		{
			"no_at_symbol.com", emptyString,
		},
		{
			"user@user@user@.com", emptyString,
		},
	}

	for _, testCase := range testCases {
		actual := extractEmailDomain(testCase.input)
		if actual != testCase.expected {
			t.Errorf(
				"For input '%s', expected '%s' but got '%s'",
				testCase.input,
				testCase.expected,
				actual)
		}
	}

}
