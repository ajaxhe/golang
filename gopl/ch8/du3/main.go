package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	var vflag = flag.Bool("v", false, "show verbose progress message")
	flag.Parse()

	roots := flag.Args()
	var n sync.WaitGroup
	fileSizes := make(chan int64)
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *vflag {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			} else {
				nfiles++
				nbytes += size
			}
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			go walkDir(filepath.Join(dir, entry.Name()), n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sem = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sem <- struct{}{}
	defer func() { <-sem }()
	entrys, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entrys
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
