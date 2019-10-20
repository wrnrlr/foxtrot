package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/corywalker/expreduce/expreduce"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/wrnrlr/foxtrot"
	"log"
	"net/http"
	"os"
	"runtime/pprof"
)

func main() {
	var debug = flag.Bool("debug", false, "Debug mode. No initial definitions.")
	//var rawterm = flag.Bool("rawterm", false, "Do not use readline. Useful for pexpect integration.")
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	var netprofile = flag.Bool("netprofile", false, "Enable live profiling at http://localhost:8080/debug/pprof/")
	//var scriptfile = flag.String("script", "", "script `file` to read from")
	var initfile = flag.String("initfile", "", "A script to run on initialization.")
	var preloadRubi = flag.Bool("preloadrubi", false, "Preload the Rubi definitions for integral support on startup.")
	var runui = flag.Bool("ui", false, "Start UI.")

	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	if *netprofile {
		go http.ListenAndServe(":8080", nil)
	}

	fmt.Printf("Welcome to Expreduce!\n\n")

	es := expreduce.NewEvalState()
	if *preloadRubi {
		fmt.Println("Pre-loading Rubi snapshot for integral support. Disable with -preloadrubi=false.")
		es.Eval(atoms.E(atoms.S("LoadRubiBundledSnapshot")))
		fmt.Println("Done loading Rubi snapshot.")
		fmt.Print("\n")
	}
	if *debug {
		es.NoInit = true
		es.ClearAll()
	}

	if *initfile != "" {
		f, err := os.Open(*initfile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		buf := new(bytes.Buffer)
		buf.ReadFrom(f)
		scriptText := buf.String()
		expreduce.EvalInterpMany(scriptText, *initfile, es)
	}

	if *runui {
		foxtrot.RunUI(es)
	}
}
