package service

import (
	"embed"
	"html/template"
	"os"
	"sort"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var f embed.FS

// 获取文件、文件夹信息
func getDirInfo(args string) []IndexFileType {
	var indexFileTypeList []IndexFileType
	var dot = "."
	entry, err := os.ReadDir(args)
	if err == nil {
		for _, dir := range entry {
			var indexFileType IndexFileType
			info, _ := dir.Info()
			indexFileType = IndexFileType{Name: info.Name(), IsDir: info.IsDir(), Size: info.Size(), ModTime: info.ModTime()}
			if len(info.Name()) > 0 && string(indexFileType.Name[0]) != dot && indexFileType.IsDir {
				indexFileTypeList = append(indexFileTypeList, indexFileType)
			}
		}
	}
	sort.SliceStable(indexFileTypeList, func(i int, j int) bool {
		var sa = []rune(indexFileTypeList[i].Name)[0]
		var sb = []rune(indexFileTypeList[j].Name)[0]
		return sb < sa
	})
	sort.SliceStable(indexFileTypeList, func(i int, j int) bool {
		return indexFileTypeList[i].IsDir
	})

	return indexFileTypeList
}

func ParseTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func Router(gitlabToken string, dir string) {
	r := gin.Default()
	templ := template.Must(template.New("").ParseFS(f, "templates/*"))
	r.SetHTMLTemplate(templ)
	r.SetFuncMap(template.FuncMap{"parseTime": ParseTime})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"dirs": getDirInfo(dir), "gitlabToken": gitlabToken})
	})
	r.GET("/projects", func(c *gin.Context) {
		c.JSON(http.StatusOK, getDirInfo(dir))
	})
	r.Run(":8080")
}
