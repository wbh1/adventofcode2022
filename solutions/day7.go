package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type DaySeven struct {
	input []string
}

type Directory struct {
	Name           string
	Size           int
	Subdirectories []*Directory
	Parent         *Directory
	Files          []string
}

type server struct {
	topDir *Directory
	dirs   []*Directory
	pwd    *Directory
}

func (dir *Directory) addFile(lsEntry string) {
	entry := strings.Split(lsEntry, " ")
	dir.Files = append(dir.Files, entry[1])

	size, err := strconv.Atoi(entry[0])
	if err != nil {
		panic(err)
	}
	dir.addSize(size)
}

func (dir *Directory) addSize(size int) {
	dir.Size += size
	if dir.Parent != nil {
		dir.Parent.addSize(size)
	}
}

func (dir *Directory) addSubdirectory(lsEntry string) *Directory {
	newDir := &Directory{Name: strings.Split(lsEntry, " ")[1], Parent: dir}
	dir.Subdirectories = append(dir.Subdirectories, newDir)
	return newDir
}

func isDir(line string) bool {
	return line[:3] == "dir"
}

func isFile(line string) bool {
	if _, err := strconv.Atoi(strings.Split(line, " ")[0]); err != nil {
		return false
	}
	return true
}

func isCommand(line string) bool {
	return line[:2] == "$ "
}

func (s *server) runCommand(cmd string) {
	cmdParts := strings.Split(cmd, " ")
	switch cmdParts[1] {
	case "cd":
		switch cmdParts[2] {
		case "..":
			s.pwd = s.pwd.Parent
		case "/":
			s.pwd = s.topDir
		default:
			matched := false
			for _, dir := range s.pwd.Subdirectories {
				if dir.Name == cmdParts[2] {
					s.pwd = dir
					matched = true
					break
				}
			}

			if !matched {
				panic(fmt.Sprintf("I don't know about this directory! (%s)", cmd))
			}
		}
	case "ls":
		break
	default:
		panic(fmt.Sprintf("I don't know what this command is! (%s)", cmd))
	}
}

var daySevenServer = &server{topDir: &Directory{Name: "/"}}

func (d *DaySeven) SetInput(input []string) {
	d.input = input
}

func (d *DaySeven) PartOne() string {
	s := daySevenServer
	for _, line := range d.input {
		if isCommand(line) {
			s.runCommand(line)
		} else if isFile(line) {
			s.pwd.addFile(line)
		} else if isDir(line) {
			newDir := s.pwd.addSubdirectory(line)
			s.dirs = append(s.dirs, newDir)
		} else {
			panic(fmt.Sprintf("I don't know what this line does! (%s) ", line))
		}
	}

	result := 0
	for _, dir := range s.dirs {
		if dir.Size <= 100000 {
			result += dir.Size
		}
	}
	return fmt.Sprintf("%d", result)
}

func (d *DaySeven) PartTwo() string {
	const (
		totalDiskSize     = 70000000
		freeSpaceRequired = 30000000
	)
	currentlyFreeSpace := totalDiskSize - daySevenServer.topDir.Size

	candidateDir := daySevenServer.topDir
	for _, dir := range daySevenServer.dirs {
		if dir.Size < candidateDir.Size && (currentlyFreeSpace+dir.Size >= freeSpaceRequired) {
			candidateDir = dir
		}
	}

	return fmt.Sprintf("%d", candidateDir.Size)
}
