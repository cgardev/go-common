package structfile

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
)

func JsonLoad[T any](path string) (*T, error) {
	jsonFile, err := os.ReadFile(path)
	if err != nil {
		return nil, ErrReadFile.WrapWithNoMessage(err)
	}

	var data T
	err = json.Unmarshal(jsonFile, &data)
	if err != nil {
		return nil, ErrUnmarshalFile.WrapWithNoMessage(err)
	}

	return &data, nil
}

func JsonLoadTo[T any](path string, data *T) error {
	jsonFile, err := os.ReadFile(path)
	if err != nil {
		return ErrReadFile.WrapWithNoMessage(err)
	}

	err = json.Unmarshal(jsonFile, &data)
	if err != nil {
		return ErrUnmarshalFile.WrapWithNoMessage(err)
	}
	return nil
}

func JsonSave[T any](path string, data T) error {
	marshal, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return ErrMarshalFile.WrapWithNoMessage(err)
	}
	err = os.WriteFile(path, marshal, fs.ModePerm)
	if err != nil {
		return ErrWriteFile.WrapWithNoMessage(err)
	}

	return nil
}

type FileJson[T any] struct {
	path string
	Data T
}

type FileJsonProps[T any] struct {
	Path        string
	DefaultData T
}

func NewFileJson[T any](props FileJsonProps[T]) (*FileJson[T], error) {
	instance := FileJson[T]{
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

func (j *FileJson[T]) Load() error {
	return JsonLoadTo[T](j.path, &j.Data)
}

func (j *FileJson[T]) Save() error {
	return JsonSave[T](j.path, j.Data)
}
