package main

import (
	"crypto/md5"
	"encoding/json"
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

type CaffeService struct {
	predictor *gocaffe.CaffePredictor
	labels    []string
}

func NewCaffeService(model, trained, label string) *CaffeService {
	ret := &CaffeService{
		predictor: gocaffe.NewCaffePredictor(model, trained),
		labels:    loadLabel(label),
	}
	if ret.labels == nil {
		log.Fatalln("label file empty")
	}
	return ret
}

func Json(w http.ResponseWriter, data map[string]interface{}, code int) {
	b, _ := json.Marshal(data)
	http.Error(w, string(b), code)
}

func (p *CaffeService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	img := queries.Get("img")
	if len(img) == 0 {
		Json(w, map[string]interface{}{
			"status": 100,
			"msg":    "empty parameter: img",
		}, 500)
		return
	}
	resp, err := http.Get(img)
	if err != nil {
		Json(w, map[string]interface{}{
			"status": 101,
			"msg":    err.Error(),
		}, 500)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Json(w, map[string]interface{}{
			"status": 102,
			"msg":    err.Error(),
		}, 500)
		return
	}
	imgname := randomFile(img)
	ioutil.WriteFile(imgname, b, 0655)
	prob := p.predictor.Predict(imgname)
	if prob == nil {
		Json(w, map[string]interface{}{
			"status": 103,
			"msg":    "fail to predict",
		}, 500)
		return
	}
	os.Remove(imgname)
	mprob := make(map[string]float64)
	maxProb := 0.0
	bestLabel := ""
	for k, v := range prob {
		mprob[p.labels[k]] = v
		if maxProb < v {
			maxProb = v
			bestLabel = p.labels[k]
		}
	}
	Json(w, map[string]interface{}{
		"status":            0,
		"prob_distribution": mprob,
		"best_label":        bestLabel,
	}, 200)
}

func main() {
	model := flag.String("model", "", "model path")
	trained := flag.String("trained", "", "trained model path")
	label := flag.String("label", "", "label file")
	flag.Parse()
	cs := NewCaffeService(*model, *trained, *label)
	http.Handle("/predict", cs)
	log.Fatal(http.ListenAndServe(":8011", nil))
}
