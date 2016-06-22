package main

import (
	"flag"
	"fmt"
	"golang.org/x/sys/unix"
	"io/ioutil"
	"os"
)

var (
	base   = flag.Int("base", 0, "base memory address")
	length = flag.Int("len", 0, "number of bytes to map")
	write  = flag.Bool("w", false, "read from stdin and write to memory")
)

func main() {
	flag.Parse()

	f, err := os.OpenFile("/dev/mem", os.O_RDWR|os.O_SYNC, 0666)
	if err != nil {
		fatal(err)
	}
	defer f.Close()

	length := *length
	prot := unix.PROT_WRITE | unix.PROT_READ
	flags := unix.MAP_SHARED
	fd := int(f.Fd())
	offset := int64(*base)

	mem, err := unix.Mmap(fd, offset, length, prot, flags)
	if err != nil {
		fatal("mmap", fd, offset, length, prot, flags, ":", err)
	}

	if *write {
		data, _ := ioutil.ReadAll(os.Stdin)
		copy(mem, data)
	} else {
		os.Stdout.Write(mem)
	}

}

func fatal(msg ...interface{}) {
	fmt.Fprintln(os.Stderr, msg...)
	os.Exit(1)
}
