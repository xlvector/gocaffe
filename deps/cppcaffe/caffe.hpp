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
  int width_, height_, channel_, nclass_;
  caffe::shared_ptr<caffe::Net<double> > net_;

 public:
  CaffePredictor(){};

  CaffePredictor(const char* model_file, const char* trained_file) {
    caffe::Caffe::set_mode(caffe::Caffe::CPU);
    net_.reset(new caffe::Net<double>(model_file, caffe::TEST));
    net_->CopyTrainedLayersFrom(trained_file);
    caffe::shared_ptr<caffe::Blob<double> > input = net_->blob_by_name("data");
    std::vector<int> shape = input->shape();
    width_ = shape[3];
    height_ = shape[2];
    channel_ = shape[1];
  }

  std::vector<double> predict(const char* imgname) {
    cv::Mat img;
    img = cv::imread(imgname, CV_LOAD_IMAGE_COLOR);
    if(img.rows < height_ || img.cols < width_) {
      return std::vector<double>();
    }
    if(img.rows != height_ || img.cols != width_) {
      std::cout << "resize image from " << img.cols << "x" << img.rows << " to " << width_ << "x" << height_ << std::endl;
      cv::Mat dst;
      cv::Size size(width_, height_);
      cv::resize(img, dst, size);
      img = dst;
    }
    caffe::Blob<double>* blob =
        new caffe::Blob<double>(1, channel_, height_, width_);
    double* data = new double[channel_ * height_ * width_];
    int cn = channel_;
    uint8_t* pixel_ptr = (uint8_t*)img.data;
    for (int i = 0; i < img.rows; i++) {
      for (int j = 0; j < img.cols; j++) {
        for (int c = 0; c < cn; c++) {
          data[c * img.rows * img.cols + i * img.cols + j] =
              (double)(pixel_ptr[i * img.cols * cn + j * cn + c]) / 255.0;
        }
      }
    }
    blob->set_cpu_data(data);
    std::vector<caffe::Blob<double>*> bottom;
    bottom.push_back(blob);
    const std::vector<caffe::Blob<double>*>& rr = net_->Forward(bottom);
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
