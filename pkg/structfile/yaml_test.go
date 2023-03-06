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

type someYamlFile struct {
	Text   string             `yaml:"text"`
	Nested someYamlFileNested `yaml:"nested"`
}

type someYamlFileNested struct {
	Text   string `yaml:"text"`
	Number int    `yaml:"number"`
}

func TestYamlLoad(t *testing.T) {
	f := fs.NewFile(t,
		"yamlFile",
		fs.WithBytes([]byte("text: Some Text\nnested:\n    text: Some nested text\n    number: 10\n")),
	)
	defer f.Remove()

	data, err := YamlLoad[someYamlFile](f.Path())
	assert.NilError(t, err)
	assert.DeepEqual(t, data, &someYamlFile{
		Text: "Some Text",
		Nested: someYamlFileNested{
			Text:   "Some nested text",
			Number: 10,
		},
	})
}

func TestYamlLoadNotExistentFile(t *testing.T) {
	notExistentPath := path.Join(".", "not-exist", "no-file-yaml")
	_, err := YamlLoad[someYamlFile](notExistentPath)
	assert.Assert(t, errorx.IsOfType(err, ErrReadFile))
}

func TestYamlLoadBadFormattedFile(t *testing.T) {
	f := fs.NewFile(t, "badFormattedYamlFile", fs.WithBytes([]byte("{")))
	defer f.Remove()
	_, err := YamlLoad[someYamlFile](f.Path())
	assert.Assert(t, errorx.IsOfType(err, ErrUnmarshalFile))
}

func TestYamlFileAll(t *testing.T) {
	tmpDir := fs.NewDir(t, "tmpdir")
	filePath := path.Join(tmpDir.Path(), "someYamlFile.yaml")
	defer tmpDir.Remove()
	f, err := NewFileYaml[someYamlFile](FileYamlProps[someYamlFile]{
		Path: filePath,
		DefaultData: someYamlFile{
			Text: "Some Text",
			Nested: someYamlFileNested{
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
		"text: Some Text\nnested:\n    text: Some nested text\n    number: 10\n",
	)

	f.Data.Nested.Number = -10
	f.Data.Text = "new value"

	err = f.Save()
	assert.NilError(t, err)

	bytes, err = os.ReadFile(filePath)
	assert.NilError(t, err)

	assert.DeepEqual(t,
		string(bytes),
		"text: new value\nnested:\n    text: Some nested text\n    number: -10\n",
	)

}
