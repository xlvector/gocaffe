# gocaffe

Implement caffe prediction by golang use CGO.

you can test by:

	go test -v -x

in caffe.go, you can change CXXFLAGS in code:

	#cgo CXXFLAGS: -I/Users/xiangliang/Code/caffe/distribute/include -I/usr/local/opt/openblas/include/ -I./deps/ -O3 -Wall

