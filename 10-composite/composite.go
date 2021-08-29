package composite

import "os"

type FileSystemNode interface {
	IsFile() bool
	CountNumOfFiles() int
	CountSizeOfFiles() int64
}

type Directory struct {
	path    string
	isFile  bool
	subNode []FileSystemNode
}

func NewDirectory() *Directory {
	dir := &Directory{}
	dir.isFile = false
	dir.path = "./xyz"
	dir.subNode = []FileSystemNode{NewFile("./xyz/a.txt"), NewFile("./xyz/b.txt"), NewFile("./xyz/c.txt")}
	return dir
}

func (d *Directory) IsFile() bool {
	return d.isFile
}

func (d *Directory) CountNumOfFiles() int {
	numOfFiles := 0
	for _, subNode := range d.subNode {
		numOfFiles += subNode.CountNumOfFiles()
	}
	return numOfFiles
}

func (d *Directory) CountSizeOfFiles() int64 {
	var sizeOfFiles int64 = 0
	for _, subNode := range d.subNode {
		sizeOfFiles += subNode.CountSizeOfFiles()
	}
	return sizeOfFiles
}

func (d *Directory) Add(file FileSystemNode) {
	d.subNode = append(d.subNode, file)
}

func (d *Directory) Remove(file FileSystemNode) {
	for i := 0; i < len(d.subNode); i++ {
		if d.subNode[i] == file {
			d.subNode = append(d.subNode[:i], d.subNode[i+1:len(d.subNode)]...)
			return
		}
	}
}

type File struct {
	path   string
	isFile bool
}

func NewFile(path string) *File {
	return &File{path: path, isFile: true}
}

func (f *File) IsFile() bool {
	return f.isFile
}

func (f *File) CountNumOfFiles() int {
	return 1
}

func (f *File) CountSizeOfFiles() int64 {
	file, err := os.Stat(f.path)
	if err != nil {
		return 0
	}
	return file.Size()
}
