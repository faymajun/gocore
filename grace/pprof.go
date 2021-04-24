package grace

import (
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/sirupsen/logrus"
)

func SetupProfiling(cpuProfile, memProfile string) {
	if cpuProfile != "" {
		f, err := os.Create(cpuProfile)
		if err != nil {
			logrus.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		OnInterrupt(func() {
			pprof.StopCPUProfile()
		})
	}
	if memProfile != "" {
		runtime.MemProfileRate = 1
		f, err := os.Create(memProfile)
		if err != nil {
			logrus.Fatal(err)
		}
		OnInterrupt(func() {
			pprof.WriteHeapProfile(f)
			f.Close()
		})
	}

}
