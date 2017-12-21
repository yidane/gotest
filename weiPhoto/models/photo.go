package models

import (
	"os"
	"path/filepath"
	"sort"
)

type Photo struct {
	URL  string
	Name string
}

type PhotoThumb struct {
	URL  string
	Name string
}

type PhotoThumbList []*PhotoThumb

func (p PhotoThumbList) Len() int {
	return len(p)
}

func (p PhotoThumbList) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}

func (p PhotoThumbList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func GetPhotoThumbList(path string, i int) (PhotoThumbList, bool, error) {
	var files = []*os.FileInfo{}
	err := filepath.Walk(path, func(path string, file os.FileInfo, err error) error {
		if file != nil && !file.IsDir() {
			files = append(files, &file)
		}
		return nil
	})
	if err != nil {
		return nil, false, err
	}

	if len(files) < i {
		return nil, false, nil
	}

	var photos = PhotoThumbList{}
	for _, file := range files {
		p := &PhotoThumb{Name: (*file).Name()}
		photos = append(photos, p)
	}

	sort.Sort(photos)
	return photos, false, nil
}
