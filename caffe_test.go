package gocaffe_test

import (
	"github.com/xlvector/gocaffe"
	"testing"
)

func TestCaffe(t *testing.T) {
	pred := gocaffe.NewCaffePredictor("car_quick.prototxt", "car_quick_iter_8000.caffemodel.h5", 150, 100)
	t.Log(pred.Predict("car.png"))
}
