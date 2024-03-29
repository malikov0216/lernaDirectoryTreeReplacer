package fileMethods

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func CopyDirectory (src, dst string) {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		log.Fatalln(err)
	}

	for _, f := range files {
		newSrc := path.Join(src, f.Name())
		newDst := path.Join(dst, f.Name())


		if f.IsDir() && f.Name() != "build" {

			CopyDirectory(newSrc, newDst)

		} else if f.Name() == "build" {

			a := strings.Split(newDst, "/")
			b := "./app/" + a[len(a) - 2]

			_ = os.MkdirAll(b, f.Mode())
			CopyDirectory(newSrc, b)

		} else if !f.IsDir() {

			a := strings.Split(newSrc, "/")
			b := a[len(a) - 2]

			if b == "build" {
				err := CopyFile(newSrc, newDst)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}
}