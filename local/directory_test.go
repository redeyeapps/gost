package local

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestExistDir(t *testing.T) {
	if !fs.Directory("aDir").Exists() {
		t.Fatalf("Dir does not exist")
	}
}

func TestCreateDir(t *testing.T) {
	assert.NoError(t, fs.Directory("dDir").Create())
}

func TestStatDir(t *testing.T) {
	info, err := fs.Directory("dDir").Stat()
	if err != nil {
		t.Errorf("Couldn't get stat: %v", err)
	}

	if info.Size < 1 {
		t.Errorf("Invalid dir size expected more than 0 got %v", info.Size)
	}

	if info.LastModified.Day() != time.Now().Day() {
		t.Errorf("Invalid dir time expected %v got %v", time.Now().Day(), info.LastModified.Day())
	}
}

func TestDeleteDir(t *testing.T) {
	assert.NoError(t, fs.Directory("dDir").Delete())
}
