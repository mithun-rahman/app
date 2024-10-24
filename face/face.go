package face

//func (rs *Resource) GetImage(w http.ResponseWriter, r *http.Request) {
//	imgName := strings.TrimSpace(r.URL.Query().Get("i"))
//	imgBucket := strings.TrimSpace(r.URL.Query().Get("b"))
//	reqID := r.Header.Get("reqID")
//
//	if len(imgName) == 0 || len(imgBucket) == 0 {
//		logging.Logger.WithFields(logrus.Fields{"reqID": reqID}).Error("Object Not Found")
//		w.WriteHeader(http.StatusNotFound)
//		w.Header().Set("Content-Type", "plain/text")
//		_, _ = w.Write(nil)
//		return
//	}
//
//	res, err := miniofs.GetObject(imgName, imgBucket)
//
//	if err != nil {
//		logging.Logger.WithFields(logrus.Fields{"reqID": reqID}).Error(err)
//		_ = render.Render(w, r, response.ErrorResponse(reserrors.ErrInternalServerError, nil))
//		return
//	}
//	logging.Logger.WithFields(logrus.Fields{"reqID": reqID}).Info("Getting Object From Minio")
//
//	fileBytes, errs := io.ReadAll(res)
//	if errs != nil {
//		logging.Logger.WithFields(logrus.Fields{"reqID": reqID}).Error(errs)
//		_ = render.Render(w, r, response.ErrorResponse(reserrors.ErrInternalServerError, nil))
//		return
//	}
//
//	//base64Encoding := base64.StdEncoding.EncodeToString(fileBytes)
//
//	//render.Respond(w, r, response.Response{
//	//	HTTPStatusCode: http.StatusOK,
//	//	Data: response.Object{
//	//		"base64": base64Encoding,
//	//	},
//	//})
//
//	w.WriteHeader(http.StatusOK)
//	w.Header().Set("Content-Type", "application/octet-stream")
//	_, err = w.Write(fileBytes)
//	if err != nil {
//		logging.Logger.WithFields(logrus.Fields{"reqID": reqID}).Error(err)
//		_ = render.Render(w, r, response.ErrorResponse(reserrors.ErrInternalServerError, nil))
//		return
//	}
//	defer func(res *minio.Object) {
//		err := res.Close()
//		if err != nil {
//			logging.Logger.WithFields(logrus.Fields{"reqID": reqID}).Error("IO Closed error ", err)
//		}
//	}(res)
//	return
//}
//
//func ImageToBase64(path string) string {
//	bytes, err := ioutil.ReadFile(path)
//	if err != nil {
//		return ""
//	}
//	var base64Encoding string
//	mimeType := http.DetectContentType(bytes)
//	switch mimeType {
//	case "image/jpeg":
//		base64Encoding += "data:image/jpeg;base64,"
//	case "image/png":
//		base64Encoding += "data:image/png;base64,"
//	}
//	base64Encoding += base64.StdEncoding.EncodeToString(bytes)
//	return base64Encoding
//}
