package local

import (

"github.com/stretchr/testify/assert"
"testing"
"time"

)

var fs = New(Config{
	BasePath: "../../mocks/fixtures",
})

func TestFiles(t *testing.T) {
	files, _ := fs.Files()
	assert.Equal(t, "../../mocks/fixtures/dummy-2.txt", files[1].GetPath())
}

func TestDirectories(t *testing.T) {
	dirs, _ := fs.Directories()
	assert.Equal(t, "../../mocks/fixtures/aDir", dirs[0].GetPath())
}

func TestWrite(t *testing.T) {
	b := []byte("abc")
	n, err := fs.File("test.txt").Write(b)
	if n != len(b) {
		t.Fatalf("Wrote %v bytes of %v bytes", n, len(b))
	}
	assert.NoError(t, err)
}

func TestRead(t *testing.T) {
	b := make([]byte, 3)
	n, err := fs.File("test.txt").Read(b)
	if n != len(b) {
		t.Fatalf("Read %v bytes of %v bytes", n, len(b))
	}
	assert.NoError(t, err)
}

func TestStat(t *testing.T)  {
	info, err := fs.File("test.txt").Stat()
	if err != nil {
		t.Errorf("Couldn't get stat: %v", err)
	}

	if info.Size != 3 {
		t.Errorf("Invalid file size expected %v got %v", 3, info.Size)
	}

	if info.LastModified.Day() != time.Now().Day() {
		t.Errorf("Invalid file time expected %v got %v", time.Now().Day(), info.LastModified.Day())
	}
}


func TestExist(t *testing.T) {
	if ! fs.File("test.txt").Exist() {
		t.Fatalf("File does not exist")
	}
}

func TestDelete(t *testing.T) {
	assert.NoError(t, fs.File("test.txt").Delete())
}

func TestNotExist(t *testing.T) {
	if fs.File("test.txt").Exist() {
		t.Fatalf("File exist")
	}
}

func TestExistDir(t *testing.T) {
	if ! fs.Directory("aDir").Exist() {
		t.Fatalf("Dir does not exist")
	}
}

func TestCreateDir(t *testing.T) {
	assert.NoError(t, fs.Directory("dDir").Create())
}

func TestStatDir(t *testing.T)  {
	info, err := fs.Directory("dDir").Stat()
	if err != nil {
		t.Errorf("Couldn't get stat: %v", err)
	}

	if info.Size != 64 {
		t.Errorf("Invalid dir size expected %v got %v", 64, info.Size)
	}

	if info.LastModified.Day() != time.Now().Day() {
		t.Errorf("Invalid dir time expected %v got %v", time.Now().Day(), info.LastModified.Day())
	}
}

func TestDeleteDir(t *testing.T) {
	assert.NoError(t, fs.Directory("dDir").Delete())
}

