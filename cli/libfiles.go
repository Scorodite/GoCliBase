package cli

import (
	"io/fs"
	"os"
	"strings"
)

type filesStdLib struct{}

func (c *CliEngine) RegisterLibFileSystem() {
	lib := &filesStdLib{}

	c.RegisterMethod(lib, "Fs.Exists", "Checks if path exists", "Path")
	c.RegisterMethod(lib, "Fs.IsDir", "Checks if path is directory", "Path")
	c.RegisterMethod(lib, "Fs.IsFile", "Checks if path is File", "Path")
	c.RegisterMethod(lib, "Fs.Cwd", "Changes current working directory", "Path")
	c.RegisterMethod(lib, "Fs.Gwd", "Returns current working directory")
	c.RegisterMethod(lib, "Fs.Mkdir", "Creates directory", "Path")
	c.RegisterMethod(lib, "Fs.MkdirAll", "Creates directory with any necessary parents", "Path")
	c.RegisterMethod(lib, "Fs.Delete", "Deletes file or directory", "Path")
	c.RegisterMethod(lib, "Fs.DeleteAll", "Removes path and any children it contains", "Path")
	c.RegisterMethod(lib, "Fs.Rename", "Renames (moves) file or directory from old path to new path", "OldPath", "NewPath")
	c.RegisterMethod(lib, "Fs.WriteFile", "Creates or overrides file with provided content", "Path", "Text")
	c.RegisterMethod(lib, "Fs.ReadFile", "Returns file contents", "Path")
	c.RegisterMethod(lib, "Fs.Ls", "Lists directory", "Path")
	c.RegisterMethod(lib, "Fs.Lss", "Lists directory for files or dirs", "Path", "Dirs")
}

// Checks if path exists
func (l *filesStdLib) Exists(path string) bool {
	_, e := os.Stat(path)
	return !os.IsNotExist(e)
}

// Checks if path is directory
func (l *filesStdLib) IsDir(path string) bool {
	r, e := os.Stat(path)
	return e == nil && r.IsDir()
}

// Checks if path is File
func (l *filesStdLib) IsFile(path string) bool {
	r, e := os.Stat(path)
	return e == nil && !r.IsDir()
}

// Changes current working directory
func (l *filesStdLib) Cwd(path string) error {
	return os.Chdir(path)
}

// Returns current working directory
func (l *filesStdLib) Gwd() (string, error) {
	return os.Getwd()
}

// Creates directory
func (l *filesStdLib) Mkdir(path string) error {
	return os.Mkdir(path, fs.ModeDir)
}

// Creates directory with any necessary parents
func (l *filesStdLib) MkdirAll(path string) error {
	return os.MkdirAll(path, fs.ModeDir)
}

// Deletes file or directory
func (l *filesStdLib) Delete(path string) error {
	return os.Remove(path)
}

// Removes path and any children it contains
func (l *filesStdLib) DeleteAll(path string) error {
	return os.RemoveAll(path)
}

// Renames (moves) file or directory from old path to new path
func (l *filesStdLib) Rename(old, new string) error {
	return os.Rename(old, new)
}

// Creates or overrides file with provided content
func (l *filesStdLib) WriteFile(path, text string) error {
	return os.WriteFile(path, []byte(text), 0644)
}

// Returns file contents
func (l *filesStdLib) ReadFile(path string) (string, error) {
	r, e := os.ReadFile(path)
	return string(r), e
}

func (l *filesStdLib) Ls(path string) (string, error) {
	ents, e := os.ReadDir(path)
	if e != nil {
		return "", e
	}
	var entNames []string
	for _, ent := range ents {
		entNames = append(entNames, ent.Name())
	}
	return strings.Join(entNames, "\n"), nil
}

func (l *filesStdLib) Lss(path string, dirs bool) (string, error) {
	ents, e := os.ReadDir(path)
	if e != nil {
		return "", e
	}
	var entNames []string
	for _, ent := range ents {
		if ent.IsDir() != dirs {
			continue
		}
		entNames = append(entNames, ent.Name())
	}
	return strings.Join(entNames, "\n"), nil
}
