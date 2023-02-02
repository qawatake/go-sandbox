package template

import (
	"bytes"
	"flag"
	"io"
	"os"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "update golden files")

const testdataDir = "testdata"

// golden fileを更新したい場合は`update`フラグを付けてテストを実行する。
func TestAgeName(t *testing.T) {
	t.Parallel()
	goldenFilePath := filepath.Join(testdataDir, "age_name.golden.html")

	params := struct {
		Name string
		Age  int
	}{
		Name: "qawatake",
		Age:  100,
	}

	if *update {
		file, err := os.Create(goldenFilePath)
		if err != nil {
			t.Error(err)
			return
		}
		if err := AgeName.Execute(file, params); err != nil {
			t.Error(err)
		}
		if err := file.Close(); err != nil {
			t.Error(err)
			return
		}
	}

	file, err := os.Open(goldenFilePath)
	if err != nil {
		t.Error(err)
		return
	}
	t.Cleanup(func() {
		file.Close()
	})

	want := &bytes.Buffer{}
	got := &bytes.Buffer{}
	if _, err := io.Copy(want, file); err != nil {
		t.Error(err)
		return
	}
	if err = AgeName.Execute(got, params); err != nil {
		t.Error(err)
		return
	}

	if !bytes.Equal(want.Bytes(), got.Bytes()) {
		t.Errorf("want: %q, got: %q", want, got)
	}
}
