package gocaffe

/*

#cgo CXXFLAGS: -I/usr/local/include/ -I/usr/local/opt/openblas/include/ -I./deps/ -O3 -Wall
#cgo LDFLAGS: -lcaffe -lboost_system -lstdc++ -lopencv_core -lopencv_highgui -lopencv_imgproc -lboost_filesystem -lopencv_imgcodecs
#include <stdlib.h>
#include "caffe.h"

*/
import "C"

import (
	"sort"
	"sync"
	"time"
	"unsafe"

	"github.com/xlvector/dlog"
)

type CaffePredictor struct {
	predictor C.CaffePredictor
	lock      *sync.Mutex
}

func NewCaffePredictor(model, trained string) *CaffePredictor {
	modelpath := C.CString(model)
	defer C.free(unsafe.Pointer(modelpath))

	trainedpath := C.CString(trained)
	defer C.free(unsafe.Pointer(trainedpath))

	return &CaffePredictor{
		predictor: C.NewCaffePredictor(modelpath, trainedpath),
		lock:      &sync.Mutex{},
	}
}

func doubleToFloats(in *C.double, size int) []float64 {
	defer C.free(unsafe.Pointer(in))
	out := (*[1 << 30]float64)(unsafe.Pointer(in))[:size:size]
	out2 := make([]float64, size)
	for i, v := range out {
		out2[i] = v
	}
	return out2
}

func (p *CaffePredictor) NClass() int {
	return int(C.NClass(p.predictor))
}

func (p *CaffePredictor) Predict(imgfile string) []float64 {
	p.lock.Lock()
	defer p.lock.Unlock()
	imgpath := C.CString(imgfile)
	defer C.free(unsafe.Pointer(imgpath))

	ret := C.Predict(p.predictor, imgpath)
	if ret == nil {
		return nil
	}
	return doubleToFloats(ret, p.NClass())
}

type Triple struct {
	index, label int
	prob         float64
}

type TripleSlice []Triple

func (c TripleSlice) Len() int {
	return len(c)
}
func (c TripleSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c TripleSlice) Less(i, j int) bool {
	return c[i].prob > c[j].prob
}

func (p *CaffePredictor) GreedyMatch(probs [][]float64) []int {
	trs := make(TripleSlice, 0, len(probs)*len(probs[0]))
	for i, ps := range probs {
		for l, v := range ps {
			trs = append(trs, Triple{i, l, v})
		}
	}
	sort.Sort(trs)
	ul := make([]byte, len(probs))
	ur := make([]byte, len(probs[0]))
	ret := make([]int, len(probs))
	for i := 0; i < len(ret); i++ {
		ret[i] = -1
	}
	for _, t := range trs {
		if ul[t.index] > 0 || ur[t.label] > 0 {
			continue
		}
		ul[t.index] = 1
		ur[t.label] = 1
		ret[t.index] = t.label
	}
	return ret
}

func (p *CaffePredictor) PredictBatch(imgs []string) [][]float64 {
	start := time.Now().UnixNano()
	ret := make([][]float64, 0, len(imgs))
	for _, img := range imgs {
		out := p.Predict(img)
		ret = append(ret, out)
	}
	dlog.Println("predict all used(ms) : ", (time.Now().UnixNano()-start)/1000000)
	return ret
}
