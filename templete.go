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

func writeHeapProfile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("mem profile %v", err)
	}
	pprof.WriteHeapProfile(f)
	f.Close()
}

func startCPUProfile(filename string) func() {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("cpu profile %v", err)
	}
	pprof.StartCPUProfile(f)
	return pprof.StopCPUProfile
}

func main() {
	log.Printf("go # %v", runtime.NumGoroutine())

	var cpuprofilename = flag.String("cpuprofilename", "", "cpu profile filename")
	var memprofilename = flag.String("memprofilename", "", "memory profile filename")
	flag.Parse()
	args := flag.Args()

	if *cpuprofilename != "" {
		fn := startCPUProfile(*cpuprofilename)
		defer fn()
	}

	sttime := time.Now()
	doMain(args)
	fmt.Printf("%v\n", time.Now().Sub(sttime))
	log.Printf("go # %v", runtime.NumGoroutine())

	if *memprofilename != "" {
		writeHeapProfile(*memprofilename)
	}
}

func doMain(args []string) {
	fmt.Printf("%v\n", args)
}
