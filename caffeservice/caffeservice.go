package main

import (
	"crypto/md5"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/xlvector/dlog"
	"github.com/xlvector/gocaffe"
)

func loadLabel(f string) []string {
	buf, err := ioutil.ReadFile(f)
	if err != nil {
		return nil
	}
	lines := strings.Split(string(buf), "\n")
	return lines
}

type IntStringPair struct {
	index int
	str   string
}

func Download(index int, url string, ch chan IntStringPair, wg *sync.WaitGroup) {
	defer wg.Done()
	c := &http.Client{
		Timeout: time.Second * 2,
	}
	dlog.Println("begin download ", url)
	resp, err := c.Get(url)
	if resp == nil || resp.Body == nil {
		dlog.Warn("nil resp")
		return
	}
	defer resp.Body.Close()
	if err != nil {
		dlog.Warn("download err: %v", err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		dlog.Warn("download err: %v", err)
		return
	}
	out := randomFile(url)
	err = ioutil.WriteFile(out, b, 0655)
	if err != nil {
		dlog.Warn("download err: %v", err)
		return
	}
	dlog.Println("download image ", url, " and save to ", out)
	ch <- IntStringPair{index, out}
}

func DownloadAll(urls []string) []string {
	wg := &sync.WaitGroup{}
	ch := make(chan IntStringPair, 100)
	for i, url := range urls {
		wg.Add(1)
		go func(index int, link string) {
			Download(index, link, ch, wg)
		}(i, url)
	}
	wg.Wait()
	close(ch)

	ret := make([]string, len(urls))
	for p := range ch {
		ret[p.index] = p.str
	}
	return ret
}

func randomFile(url string) string {
	return fmt.Sprintf("%d_%x.jpg", time.Now().UnixNano(), md5.Sum([]byte(url)))
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

func (p *CaffeService) Label(i int) string {
	if i < 0 || i >= len(p.labels) {
		return "unknown"
	}
	return p.labels[i]
}

func (p *CaffeService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tmpImgs := strings.Split(r.FormValue("imgs"), "|")
	imgs := make([]string, 0, len(tmpImgs))
	for _, img := range tmpImgs {
		if len(img) > 0 {
			imgs = append(imgs, img)
		}
	}
	if len(imgs) == 0 {
		Json(w, map[string]interface{}{
			"status": 100,
			"msg":    "no image to predict",
		}, 500)
		return
	}

	fs := DownloadAll(imgs)
	for k, f := range fs {
		if len(f) == 0 {
			Json(w, map[string]interface{}{
				"status": 101,
				"msg":    "fail to download image: " + imgs[k],
			}, 500)
			return
		}
	}

	probs := p.predictor.PredictBatch(fs)

	for _, f := range fs {
		os.Remove(f)
	}

	for k, ps := range probs {
		if ps == nil || len(ps) == 0 {
			Json(w, map[string]interface{}{
				"status": 102,
				"msg":    "fail to predict for image: " + imgs[k],
			}, 500)
		}
	}
	bestMatch := p.predictor.GreedyMatch(probs)
	results := make([]map[string]interface{}, len(bestMatch))
	for k, bm := range bestMatch {
		dis := make(map[string]float64)
		for j, v := range probs[k] {
			dis[p.Label(j)] = v
		}
		results[k] = map[string]interface{}{
			"img":          imgs[k],
			"label":        p.Label(bm),
			"distribution": dis,
		}
	}
	Json(w, map[string]interface{}{
		"status":  0,
		"msg":     "ok",
		"results": results,
	}, 200)
}

func main() {
	dlog.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	model := flag.String("model", "", "model path")
	trained := flag.String("trained", "", "trained model path")
	label := flag.String("label", "", "label file")
	flag.Parse()
	cs := NewCaffeService(*model, *trained, *label)
	http.Handle("/predict", cs)
	log.Fatal(http.ListenAndServe(":8011", nil))
}
