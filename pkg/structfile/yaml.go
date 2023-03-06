package structfile

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
)

func YamlLoad[T any](path string) (*T, error) {
	var data T
	err := YamlLoadTo[T](path, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func YamlLoadTo[T any](path string, data *T) error {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return ErrReadFile.WrapWithNoMessage(err)
	}

	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		return ErrUnmarshalFile.WrapWithNoMessage(err)
	}
	return nil
}

func YamlSave[T any](path string, data T) error {
	marshal, err := yaml.Marshal(data)
	if err != nil {
		return ErrMarshalFile.WrapWithNoMessage(err)
	}
	err = os.WriteFile(path, marshal, fs.ModePerm)
	if err != nil {
		return ErrWriteFile.WrapWithNoMessage(err)
	}

	return nil
}

type FileYaml[T any] struct {
	path string
	Data T
}

type FileYamlProps[T any] struct {
	Path        string
	DefaultData T
}

func NewFileYaml[T any](props FileYamlProps[T]) (*FileYaml[T], error) {
	instance := FileYaml[T]{
		path: props.Path,
	}
	_, err := os.Stat(props.Path)

	if errors.Is(err, os.ErrNotExist) {
		instance.Data = props.DefaultData
		err = instance.Save()
	} else {
		err = instance.Load()
	}

	if err != nil {
		return nil, err
	}

	return &instance, nil
}

func (y *FileYaml[T]) Load() error {
	return YamlLoadTo[T](y.path, &y.Data)
}

func (y *FileYaml[T]) Save() error {
	return YamlSave[T](y.path, y.Data)
}
