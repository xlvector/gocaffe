package gocaffe_test

import (
	"github.com/xlvector/gocaffe"
	"testing"
)

func TestCaffe(t *testing.T) {
	pred := gocaffe.NewCaffePredictor("./examples/car_quick.prototxt", "./examples/car_quick_iter_8000.caffemodel.h5")
	t.Log(pred.Predict("./examples/car.png"))

	t.Log(pred.Predict("./examples/car.jpg"))
}
