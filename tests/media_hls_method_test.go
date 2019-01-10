package test

import (
	"github.com/xfrr/goffmpeg/models"
	"strings"
	"testing"
)

func TestMediaHlsStartNumber(t *testing.T) {
	var results []string = []string{}
	var asserts []string = []string{"-start_number", "1"}
	m := models.Mediafile{}
	m.SetHlsStartNumber(1)

	if !assertOutput(m.ToStrCommand(), results, asserts) {
		t.Fail()
	}
}

func TestMediaHlsWrapSize(t *testing.T) {
	var results []string = []string{}
	var asserts []string = []string{"-hls_wrap", "1"}
	m := models.Mediafile{}
	m.SetHlsWrapSize(1)

	if !assertOutput(m.ToStrCommand(), results, asserts) {
		t.Fail()
	}
}

func TestMediaHlsSegmentFilename(t *testing.T) {
	var results []string = []string{}
	var asserts []string = []string{"-hls_segment_filename", "file%03d.ts"}
	m := models.Mediafile{}
	m.SetHlsSegmentFilename("file%03d.ts")

	if !assertOutput(m.ToStrCommand(), results, asserts) {
		t.Fail()
	}
}

func assertOutput(cmds []string, results []string, asserts []string) bool {
	for id, cmd := range cmds {
		if strings.Contains(cmd, asserts[0]) {
			results = append(results, cmd)
			results = append(results, cmds[id+1])
			break
		}
	}

	if len(results) != len(asserts) {
		return false
	}
	for id, assert := range asserts {
		result := results[id]
		if result != assert {
			return false
		}
	}

	return true
}
