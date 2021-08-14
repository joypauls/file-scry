// File System Tools
// Dealing with directories (defining as a collection of files/directories here)
package fst

import (
	"fmt"
	"log"
	"os"
)

func readFiles(p *Path) []File {
	rawFiles, err := os.ReadDir(p.Cur()) // DirEntry slice
	if err != nil {
		log.Fatal(err)
	}
	// is this the right way to build this slice?
	files := make([]File, len(rawFiles))
	for i, f := range rawFiles {
		files[i] = MakeFile(f)
	}
	return files
}

type Directory struct {
	files []File
	size  int
}

func (d *Directory) File(i int) *File {
	if i < 0 || i >= d.size {
		fmt.Println(i)
		log.Fatal("Requested a file from index that doesn't exist")
	}
	return &d.files[i]
}

func (d *Directory) Files() []File {
	return d.files
}

func (d *Directory) Size() int {
	return d.size
}

func NewDirectory(p *Path) *Directory {
	d := new(Directory)
	d.files = readFiles(p)
	d.size = len(d.files)
	return d
}