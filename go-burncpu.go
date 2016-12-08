package main

import "runtime"

func main() {
	numCPU := runtime.NumCPU()
	done := make(chan error, 0)
	for i := 0; i < numCPU; i++ {
		go func() {
			n := 0
			for true {
				n = n + 1
				_ = n
			}
		}()
	}
	<-done
}
