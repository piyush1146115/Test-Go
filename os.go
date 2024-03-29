package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	files, err := ioutil.ReadDir("/home/piyush/go/src/github.com/piyush1146115/etcd-experiments/demo/etcd-0/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name() != "restore" {
			fmt.Println(file.Name())
			//os.Remove(file.Name())
			err = os.RemoveAll("/home/piyush/go/src/github.com/piyush1146115/etcd-experiments/demo/etcd-0/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	//err = os.Rename("/home/piyush/go/src/github.com/piyush1146115/etcd-experiments/demo/etcd-0/restore", "/home/piyush/go/src/github.com/piyush1146115/etcd-experiments/demo/etcd-0")
	//if err != nil {
	//	log.Fatal(err)
	//}

	files, err = ioutil.ReadDir("/home/piyush/go/src/github.com/piyush1146115/etcd-experiments/demo/etcd-0/restore")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		//os.Remove(file.Name())
		err = os.Rename("/home/piyush/go/src/github.com/piyush1146115/etcd-experiments/demo/etcd-0/restore/"+file.Name(), "/home/piyush/go/src/github.com/piyush1146115/etcd-experiments/demo/etcd-0/"+file.Name())
		if err != nil {
			log.Fatal(err)
		}
	}

	err = os.RemoveAll("/home/piyush/go/src/github.com/piyush1146115/etcd-experiments/demo/etcd-0/restore")
	if err != nil {
		log.Fatal(err)
	}

	//err = Dir("/home/piyush/go/src/github.com/piyush1146115/etcd-experiments/demo/etcd-0/restore", "/home/piyush/go/src/github.com/piyush1146115/etcd-experiments/demo/etcd-0")
	//if err != nil {
	//	log.Fatal(err)
	//}

}

// Dir copies a whole directory recursively
func Dir(src string, dst string) error {
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
			if err = Dir(srcfp, dstfp); err != nil {
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

// File copies a single file from src to dst
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
