package vscr

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func init() {
	// Default configure the global log to also print
	// line number since this will be used as a script...
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}

// Write the entire string to a temp file and return the filename
func MustWriteTemp(data []byte) string {
	dir := os.TempDir()
	dir = filepath.Join(dir, "kalle.vim")
	e := ioutil.WriteFile(dir, data, 0777)
	if e != nil {
		log.Fatalln(e)
	}
	return dir
}

func MustCopy(src, dst string) {
	fi, e := os.Open(src)
	if e != nil {
		log.Fatalln(e)
	}
	defer fi.Close()
	fo, e := os.Create(dst)
	if e != nil {
		log.Fatalln(e)
	}
	defer fo.Close()
	_, e = io.Copy(fo, fi)
	if e != nil {
		log.Fatalln(e)
	}
}

// Run the cmd and exit if it fails
func MustRun(cmd string, arg ...string) {
	c := exec.Command(cmd, arg...)
	e := c.Run()
	if e != nil {
		log.Fatalln(e)
	}
}
