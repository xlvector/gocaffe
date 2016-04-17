#ifndef CAFFE_H
#define CAFFE_H

typedef void* CaffePredictor;

CaffePredictor NewCaffePredictor(const char * model_file, const char * trained_file);
double * Predict(CaffePredictor predictor, const char * imgname);
int NClass(CaffePredictor predictor);

#endif