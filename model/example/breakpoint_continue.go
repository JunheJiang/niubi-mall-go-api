package example

import "niubi-mall/model/common"

// ExaFile --- file struct, 文件结构体
type ExaFile struct {
	ID           int
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
	CreateTime   common.JSONTime
	UpdateTime   common.JSONTime
}

// ExaFileChunk --- file chunk struct, 切片结构体
type ExaFileChunk struct {
	ID              int
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
	CreateTime      common.JSONTime
	UpdateTime      common.JSONTime
}
