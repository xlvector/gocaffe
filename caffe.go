package gocaffe

/*

#cgo CXXFLAGS: -I/usr/local/include/ -I/usr/local/opt/openblas/include/ -I./deps/ -O3 -Wall
#cgo LDFLAGS: -lcaffe -lboost_system -lstdc++ -lopencv_core -lopencv_highgui -lopencv_imgproc -lboost_filesystem -lopencv_imgcodecs
#include <stdlib.h>
#include "caffe.h"

*/
import "C"

import "unsafe"

type CaffePredictor struct {
	predictor C.CaffePredictor
}

func NewCaffePredictor(model, trained string) *CaffePredictor {
	modelpath := C.CString(model)
	defer C.free(unsafe.Pointer(modelpath))

	trainedpath := C.CString(trained)
	defer C.free(unsafe.Pointer(trainedpath))

	return &CaffePredictor{
		C.NewCaffePredictor(modelpath, trainedpath),
	}
}

func doubleToFloats(in *C.double, size int) []float64 {
	defer C.free(unsafe.Pointer(in))
	out := (*[1 << 30]float64)(unsafe.Pointer(in))[:size:size]
	return out
}

func (p *CaffePredictor) NClass() int {
	return int(C.NClass(p.predictor))
}

func (p *CaffePredictor) Predict(imgfile string) []float64 {
	imgpath := C.CString(imgfile)
	defer C.free(unsafe.Pointer(imgpath))

	ret := C.Predict(p.predictor, imgpath)
	return doubleToFloats(ret, p.NClass())
}
