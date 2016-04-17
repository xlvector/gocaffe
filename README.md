# gocaffe

Implement caffe prediction by golang use CGO.

you can test by:

	go test -v -x

in caffe.go, you can change CXXFLAGS in code:

	#cgo CXXFLAGS: -I/Users/xiangliang/Code/caffe/distribute/include -I/usr/local/opt/openblas/include/ -I./deps/ -O3 -Wall

you can use caffepred to test images from web:

	go install github.com/xlvector/gocaffe/caffepred

	caffepred --model ../examples/car_quick.prototxt --trained ../examples/car_quick_iter_8000.caffemodel.h5 --img "http://image.guazistatic.com/gz01160412/16/54/77fa28cea0d8e7547509b2719a82b046.jpg"
