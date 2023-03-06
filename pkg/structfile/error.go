package structfile

import "github.com/joomcode/errorx"

var (
	nsStructFile     = errorx.NewNamespace("StructFile")
	ErrReadFile      = nsStructFile.NewType("ReadFile")
	ErrWriteFile     = nsStructFile.NewType("WriteFile")
	ErrUnmarshalFile = nsStructFile.NewType("UnmarshalFile")
	ErrMarshalFile   = nsStructFile.NewType("MarshalFile")
)
