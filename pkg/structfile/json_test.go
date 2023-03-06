//go:build unit_test

package structfile

import (
	"github.com/joomcode/errorx"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/fs"
	"os"
	"path"
	"testing"
)

type someJsonFile struct {
	Text   string             `json:"text"`
	Nested someJsonFileNested `json:"nested"`
}

type someJsonFileNested struct {
	Text   string `json:"text"`
	Number int    `json:"number"`
}

func TestJsonLoad(t *testing.T) {
	f := fs.NewFile(t,
		"jsonFile",
		fs.WithBytes([]byte("{\"Text\":\"Some Text\",\"Nested\":{\"Text\":\"Some nested text\",\"Number\":10}}")),
	)
	defer f.Remove()

	data, err := JsonLoad[someJsonFile](f.Path())
	assert.NilError(t, err)
	assert.DeepEqual(t, data, &someJsonFile{
		Text: "Some Text",
		Nested: someJsonFileNested{
			Text:   "Some nested text",
			Number: 10,
		},
	})
}

func TestJsonLoadNotExistentFile(t *testing.T) {
	notExistentPath := path.Join(".", "not-exist", "no-file-json")
	_, err := JsonLoad[someJsonFile](notExistentPath)
	assert.Assert(t, errorx.IsOfType(err, ErrReadFile))
}

func TestJsonLoadBadFormattedFile(t *testing.T) {
	f := fs.NewFile(t, "badFormattedJsonFile", fs.WithBytes([]byte("{")))
	defer f.Remove()
	_, err := JsonLoad[someJsonFile](f.Path())
	assert.Assert(t, errorx.IsOfType(err, ErrUnmarshalFile))
}

func TestJsonFileAll(t *testing.T) {
	tmpDir := fs.NewDir(t, "tmpdir")
	filePath := path.Join(tmpDir.Path(), "someJsonFile.json")
	defer tmpDir.Remove()
	f, err := NewFileJson[someJsonFile](FileJsonProps[someJsonFile]{
		Path: filePath,
		DefaultData: someJsonFile{
			Text: "Some Text",
			Nested: someJsonFileNested{
				Text:   "Some nested text",
				Number: 10,
			},
		},
	})
	assert.NilError(t, err)

	_, err = os.Stat(filePath)
	assert.NilError(t, err)

	bytes, err := os.ReadFile(filePath)
	assert.NilError(t, err)

	assert.DeepEqual(t,
		string(bytes),
		"{\n\t\"text\": \"Some Text\",\n\t\"nested\": {\n\t\t\"text\": \"Some nested text\",\n\t\t\"number\": 10\n\t}\n}",
	)

	f.Data.Nested.Number = -10
	f.Data.Text = "new value"

	err = f.Save()
	assert.NilError(t, err)

	bytes, err = os.ReadFile(filePath)
	assert.NilError(t, err)

	assert.DeepEqual(t,
		string(bytes),
		"{\n\t\"text\": \"new value\",\n\t\"nested\": {\n\t\t\"text\": \"Some nested text\",\n\t\t\"number\": -10\n\t}\n}",
	)

}
