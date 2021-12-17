package service

import "time"

/**
*	列表文件或文件夹对象
 */
type IndexFileType struct {
	Name    string
	IsDir   bool
	Size    int64
	ModTime time.Time
}
