package main

import (
	"flag"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"io"
	"fmt"
	"io/ioutil"
	"path"

)

type MyFilePath struct {
	Path    string
	NewPath string
}

// cool function to replace last string https://stackoverflow.com/questions/27942769/replace-strings-n-times-starting-from-the-end
func Rev(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (f *MyFilePath) MatchAndReplacePath(pattern string, repl string) {
	f.NewPath = Rev(strings.Replace(Rev(f.Path), Rev(pattern), Rev(repl), 1))
}

func File(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

func CopyDir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = File(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

func WalkDirectory(dir string, pattern string, repl string) error {
	startDir := dir
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}


		f := MyFilePath{Path: path}

		if !info.IsDir() {
			f.MatchAndReplacePath(pattern, repl)
		}


		e := os.Rename(f.Path, f.NewPath)

		if info.IsDir() && startDir != path {
			_ = WalkDirectory(f.NewPath, pattern, repl)
		}

		if e != nil {
			return err
		}
		return nil
	})
	return err
}

func main() {
	d := flag.String("d", "f", "directory to run rename")
	flag.Parse()
	os.RemoveAll(*d)
	CopyDir("bckp", *d)
	_ = WalkDirectory(*d, "n", "test_out_")
}
