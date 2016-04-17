#ifndef CAFFE_HPP
#define CAFFE_HPP

#define CPU_ONLY 1
#include <caffe/caffe.hpp>
#include <iostream>
#include <string>
#include <vector>
#include <opencv2/imgproc/imgproc.hpp>
#include <opencv2/highgui/highgui.hpp>
#include <opencv2/core/core.hpp>

namespace cppcaffe {
class CaffePredictor {
 private:
  int width_, height_, nclass_;
  caffe::shared_ptr<caffe::Net<float> > net_;

 public:
  CaffePredictor(){};

  CaffePredictor(const char* model_file, const char* trained_file, int width,
                 int height) {
    caffe::Caffe::set_mode(caffe::Caffe::CPU);
    net_.reset(new caffe::Net<float>(model_file, caffe::TEST));
    net_->CopyTrainedLayersFrom(trained_file);
    width_ = width;
    height_ = height;
  }

  std::vector<double> predict(const char* imgname) {
    cv::Mat img;
    img = cv::imread(imgname, CV_LOAD_IMAGE_COLOR);
    caffe::Blob<float>* blob = new caffe::Blob<float>(1, 3, height_, width_);
    float* data = new float[3 * height_ * width_];
    int cn = 3;
    uint8_t* pixel_ptr = (uint8_t*)img.data;
    for (int i = 0; i < img.rows; i++) {
      for (int j = 0; j < img.cols; j++) {
        for (int c = 0; c < cn; c++) {
          data[c * img.rows * img.cols + i * img.cols + j] =
              (float)(pixel_ptr[i * img.cols * cn + j * cn + c]) / 256.0;
        }
      }
    }
    blob->set_cpu_data(data);
    std::vector<caffe::Blob<float>*> bottom;
    bottom.push_back(blob);
    const std::vector<caffe::Blob<float>*>& rr = net_->Forward(bottom);
    delete blob;
    delete[] data;
    std::vector<double> ret(rr[0]->count());
    for (int i = 0; i < ret.size(); i++) {
      ret[i] = rr[0]->cpu_data()[i];
    }
    nclass_ = rr[0]->count();
    return ret;
  }

  int nclass() { return nclass_; }
};
};

#endif
