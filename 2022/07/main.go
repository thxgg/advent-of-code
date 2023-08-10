package main

import (
	_ "embed"
	"fmt"
	_ "github.com/emirpasic/gods/stacks/arraystack"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type File struct {
	name string
	size int
}

type Directory struct {
	name           string
	size           int
	files          []*File
	subdirectories []*Directory
	parent         *Directory
}

func calcDirSize(dir *Directory) int {
	size := 0
	for _, subdir := range dir.subdirectories {
		size += calcDirSize(subdir)
	}
	for _, file := range dir.files {
		size += file.size
	}
	dir.size = size
	return size
}

func sumDirsLt100k(dir *Directory) int {
	sum := 0
	if dir.size < 100000 {
		sum += dir.size
	}
	for _, subdir := range dir.subdirectories {
		sum += sumDirsLt100k(subdir)
	}
	return sum
}

func partOne() {
	rootDir := &Directory{name: "/"}
	currentDir := rootDir
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "$") { // Command
			if strings.HasPrefix(line, "$ cd ") { // $ cd
				targetDir := line[5:]
				switch targetDir {
				case "/":
					currentDir = rootDir
				case "..":
					currentDir = currentDir.parent
				default:
					currentDir = currentDir.subdirectories[slices.IndexFunc(currentDir.subdirectories, func(d *Directory) bool {
						return d.name == targetDir
					})]
				}
			} else { // $ ls
				continue
			}
		} else { // ls output
			data := strings.Split(line, " ")
			switch data[0] {
			case "dir":
				currentDir.subdirectories = append(currentDir.subdirectories, &Directory{name: data[1], parent: currentDir})
			default:
				size, _ := strconv.Atoi(data[0])
				currentDir.files = append(currentDir.files, &File{name: data[1], size: size})
			}
		}
	}

	// Calculate directory sizes
	calcDirSize(rootDir)

	// Sum directory sizes less than 100k
	fmt.Println(sumDirsLt100k(rootDir))
}

func calcDirSizeAndIndex(dir *Directory, index map[int]*Directory) int {
	size := 0
	for _, subdir := range dir.subdirectories {
		size += calcDirSizeAndIndex(subdir, index)
	}
	for _, file := range dir.files {
		size += file.size
	}
	dir.size = size
	index[dir.size] = dir
	return size
}

func partTwo() {
	rootDir := &Directory{name: "/"}
	currentDir := rootDir
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "$") { // Command
			if strings.HasPrefix(line, "$ cd ") { // $ cd
				targetDir := line[5:]
				switch targetDir {
				case "/":
					currentDir = rootDir
				case "..":
					currentDir = currentDir.parent
				default:
					currentDir = currentDir.subdirectories[slices.IndexFunc(currentDir.subdirectories, func(d *Directory) bool {
						return d.name == targetDir
					})]
				}
			} else { // $ ls
				continue
			}
		} else { // ls output
			data := strings.Split(line, " ")
			switch data[0] {
			case "dir":
				currentDir.subdirectories = append(currentDir.subdirectories, &Directory{name: data[1], parent: currentDir})
			default:
				size, _ := strconv.Atoi(data[0])
				currentDir.files = append(currentDir.files, &File{name: data[1], size: size})
			}
		}
	}

	// Calculate directory sizes
	sizeIndex := make(map[int]*Directory)
	calcDirSizeAndIndex(rootDir, sizeIndex)

	// Choose smallest folder to delete
	reqSize := rootDir.size - 40000000
	//fmt.Printf("Required size: %v\n", reqSize)
	minSufSize := math.MaxInt
	for key := range sizeIndex {
		//fmt.Printf("Dir: %v, Size: %v, Min: %v\n", sizeIndex[key].name, key, minSufSize)
		if key >= reqSize && key < minSufSize {
			minSufSize = key
		}
	}
	fmt.Println(minSufSize)
}

func main() {
	//partOne()
	partTwo()
}
