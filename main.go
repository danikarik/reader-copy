package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type process struct {
	Name string
	Data io.Reader
}

func doSomeStuff(pp ...process) error {
	for _, p := range pp {
		content, err := ioutil.ReadAll(p.Data)
		if err != nil {
			return err
		}

		fmt.Println("%s -> %v", p.Name, string(content))
	}

	return nil
}

func openFile(path string) (io.ReadCloser, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return ioutil.NopCloser(f), nil
}

func main() {
	pp := []process{}

	for _, path := range []string{"log1.txt", "log2.txt"} {
		r, err := openFile(path)
		if err == nil {
			pp = append(pp, process{Name: path, Data: r})

			if err := r.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}

	if err := doSomeStuff(pp...); err != nil {
		log.Fatal(err)
	}
}
