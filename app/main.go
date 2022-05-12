package main

import (
	"bytes"
	"flag"
	"fmt"
	"hash/crc64"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func getCRC64(t uint64, b []byte) string {
	h := crc64.New(crc64.MakeTable(t))
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	src := flag.String("src", "", "input file")
	dst := flag.String("dst", "", "target file")
	out := flag.String("out", "output1", "output file")
	flag.Parse()

	if src == nil || dst == nil || out == nil || len(*src) == 0 || len(*dst) == 0 {
		panic("some input is nil")
	}

	// load input src
	b1, err := ioutil.ReadFile(*src)
	if err != nil {
		panic(err)
	}
	c1 := getCRC64(crc64.ECMA, b1)
	fmt.Printf("src crc64 ECMA is %v, len %v, filename is %v\n", c1, len(b1), *src)

	// load dst
	b2, err := ioutil.ReadFile(*dst)
	if err != nil {
		panic(err)
	}
	c2 := getCRC64(crc64.ECMA, b2)
	fmt.Printf("dst crc64 ECMA is %v, len %v, filename is %v\n", c2, len(b2), *dst)

	// len of dst must longer than src
	if len(b2)-len(b1) < 8 {
		fmt.Println(len(b1), len(b2))
		panic("dst file length should longer than src, at least 8 bytes")
	}

	b1Filled := make([]byte, len(b2))
	for idx := range b1 {
		b1Filled[idx] = b1[idx]
	}
	tmpFileName := ".tmp_file"
	err = ioutil.WriteFile(tmpFileName, b1Filled, 0666)
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFileName)

	start := len(b1)
	if len(b2)-start > 256 {
		start = len(b2) - 256
	}
	changeRange := fmt.Sprintf("%v.0:%v.8", start, len(b2))

	// open output
	f3, err := os.OpenFile(*out, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f3.Close()

	bf3 := bytes.NewBuffer(nil)
	wt := io.MultiWriter(bf3, f3)
	// run
	cmd := exec.Command("./crchack", "-w64", "-p42f0e1eba9ea3693", "-iffffffffffffffff", "-rR", "-xffffffffffffffff", "-b", changeRange, tmpFileName, c2)
	// fmt.Println(cmd.String())
	cmd.Stdout = wt
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	// generate output
	b3, err := ioutil.ReadAll(bf3)
	if err != nil {
		panic(err)
	}
	getCRC64(crc64.ECMA, b3)
	c3 := getCRC64(crc64.ECMA, b3)
	fmt.Printf("out crc64 ECMA is %v, len %v, filename is %v\n", c3, len(b3), *out)
}
