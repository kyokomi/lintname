package lintname

import "testing"

func TestLintName(t *testing.T) {
	if "HogeID" != LintName("HogeId") {
		t.Errorf("ERROR: %s", "lint error")
	}
}
