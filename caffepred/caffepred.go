package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"github.com/xlvector/gocaffe"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func loadLabel(f string) []string {
	buf, err := ioutil.ReadFile(f)
	if err != nil {
		return nil
	}
	lines := strings.Split(string(buf), "\n")
	return lines
}

func randomFile(url string) string {
	return fmt.Sprintf("%d_%x.tmp", time.Now().UnixNano(), md5.Sum([]byte(url)))
}

func main() {
	model := flag.String("model", "", "model path")
	trained := flag.String("trained", "", "trained model path")
	img := flag.String("img", "", "image path, include url")
	label := flag.String("label", "", "label file")
	flag.Parse()

	predictor := gocaffe.NewCaffePredictor(*model, *trained)
	prob := []float64{}
	if strings.HasPrefix(*img, "http://") || strings.HasPrefix(*img, "https://") {
		resp, err := http.Get(*img)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		imgname := randomFile(*img)
		ioutil.WriteFile(imgname, b, 0655)
		prob = predictor.Predict(imgname)
		os.Remove(imgname)
	} else {
		prob = predictor.Predict(*img)
	}

	labels := loadLabel(*label)
	for k, v := range prob {
		if labels == nil {
			fmt.Println(k, "\t", v)
		} else {
			fmt.Println(labels[k], "\t", v)
		}
	}
}
