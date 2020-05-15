package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// $GOPATH/bin/goconvey

func TestUploadTest(t *testing.T) {
	// 테스트 할 이미지를 가져온다.
	assert := assert.New(t)
	dirname, _ := os.Getwd()
	path := dirname + "/testImage.png"
	file, _ := os.Open(path)
	defer file.Close()

	// upload 폴더를 비운다.
	os.RemoveAll("./uploads")

	// formfile을 만든다.
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)
	io.Copy(multi, file)
	writer.Close()

	// request 요청을 한다.
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-type", writer.FormDataContentType())

	// 요청이 잘 갔는지 검증한다.
	uploadsHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	uploadFilePath := "./uploads/" + filepath.Base(path)
	_, err = os.Stat(uploadFilePath)
	assert.NoError(err)

	// 요청으로 업로드 된 파일이 맞는지 검증한다.
	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}
	uploadFile.Read(uploadData)
	originFile.Read(originData)

	assert.Equal(originData, uploadData)
}
