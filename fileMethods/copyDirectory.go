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

			parentPath := getParentDst(newDst, "./app/")

			_ = os.MkdirAll(parentPath, f.Mode())
			CopyDirectory(newSrc, parentPath)

		} else if !f.IsDir() {

			parentPath := getParentDst(newDst)

			if parentPath == "build" {
				err := CopyFile(newSrc, newDst)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}
}
func getParentDst (dst ...string) string{
	a := strings.Split(dst[0], "/")
	b := dst[1] + a[len(a) -2]
	return b
}