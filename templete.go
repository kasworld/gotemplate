package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	log.Printf("go # %v", runtime.NumGoroutine())

	var cpuprofilename = flag.String("cpuprofilename", "", "cpu profile filename")
	var memprofilename = flag.String("memprofilename", "", "memory profile filename")
	flag.Parse()
	args := flag.Args()

	if *cpuprofilename != "" {
		f, err := os.Create(*cpuprofilename)
		if err != nil {
			log.Fatalf("profile %v", err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	if *memprofilename != "" {
		f, err := os.Create(*memprofilename)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}

	sttime := time.Now()
	doMain(args)
	fmt.Printf("%v\n", time.Now().Sub(sttime))
	log.Printf("go # %v", runtime.NumGoroutine())
}

func doMain(args []string) {
	fmt.Printf("%v\n", args)
}
