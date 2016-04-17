extern "C" {
#include "caffe.h"
}

#include <vector>
#include "cppcaffe/caffe.hpp"

CaffePredictor NewCaffePredictor(const char *model_file,
                                 const char *trained_file, int width,
                                 int height) {
  return (CaffePredictor)(
      new cppcaffe::CaffePredictor(model_file, trained_file, width, height));
}

double *Predict(CaffePredictor predictor, const char *imgname) {
  std::vector<double> ret =
      ((cppcaffe::CaffePredictor *)predictor)->predict(imgname);
  double *fret = (double *)malloc(sizeof(double) * ret.size());
  for (int i = 0; i < ret.size(); i++) {
    fret[i] = ret[i];
  }
  return fret;
}

int NClass(CaffePredictor predictor) {
  return ((cppcaffe::CaffePredictor *)predictor)->nclass();
}