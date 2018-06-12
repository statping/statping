package main

import (
	"fmt"
	"strings"
)

func (f *Failure) ParseError() string {
	err := strings.Contains(f.Issue, "operation timed out")
	if err {
		return fmt.Sprintf("HTTP Request timed out after x seconds")
	}
	err = strings.Contains(f.Issue, "x509: certificate is valid")
	if err {
		return fmt.Sprintf("SSL Certificate invalid")
	}
	return f.Issue
}
