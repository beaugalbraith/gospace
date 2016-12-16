package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

var counter = struct {
	sync.RWMutex
	m map[byte]int
}{m: make(map[byte]int)}

type pair struct {
	byteValue byte
	quantity  int
}

type pairList []pair

func (p pairList) Len() int           { return len(p) }
func (p pairList) Less(i, j int) bool { return p[i].quantity < p[j].quantity }
func (p pairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func bytesFromFile(f *os.File) {
	// byteMap := make(map[byte]int)

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("can't read")
	}
	for _, b := range bs {
		//fmt.Printf("%#x \n", b)
		//byteMap[b]++
		counter.Lock()
		counter.m[b]++
		counter.Unlock()
	}
	var keys []byte
	for value, _ := range counter.m {
		counter.RLock()
		keys = append(keys, value)
		//fmt.Printf("%#x : %d\n", value, count)
	}

	byteList := make(pairList, len(counter.m))
	i := 0
	for k, v := range counter.m {
		byteList[i] = pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(byteList))
	for _, bytep := range byteList {
		fmt.Printf("%#x : %d\n", bytep.byteValue, bytep.quantity)
	}
}

func main() {
	fmt.Println("vim-go")
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln("couldn't open")
	}
	defer f.Close()
	bytesFromFile(f)
}
