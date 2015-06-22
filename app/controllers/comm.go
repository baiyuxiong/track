package controllers

import (
	"github.com/baiyuxiong/track/app"
	"github.com/baiyuxiong/track/app/utils"
	"github.com/baiyuxiong/track/app/models"
	"github.com/revel/revel"
	"time"
	"regexp"
	"os"
	"io"
	"path/filepath"
	"fmt"
)

type Comm struct {
	BaseController
}

func (c Comm) SendSms(username string) revel.Result {
	c.Validation.Required(username).Message("手机号不能为空")
	c.Validation.Match(username, regexp.MustCompile("^(1)\\d{10}$")).Message("手机号格式不正确")

	if c.Validation.HasErrors() {
		return c.Err(utils.ValidationErrorToString(c.Validation.Errors))
	}

	//验证是否已被注册
	var s = &models.SmsCode{Username: username}
	has, _ := app.Engine.Get(s)

	code := utils.Sms_code()
	s.Code = code
	s.UpdatedAt = time.Now()

	var err error
	if has {
		_, err = app.Engine.Update(s)
	} else {
		_, err = app.Engine.Insert(s)
	}

	if err != nil {
		return c.Err(err.Error())
	} else {
		return c.OK(nil)
	}
}

type UploadInfo struct {
	RealPath string `json:"realPath"`
	fileName string `json:"fileName"`
	clientName string `json:"clientName"`
	fileExt string `json:"fileExt"`
}

//上传文件
func (c Comm) Upload(fileType string) revel.Result {

	if !utils.StringInSlice(fileType,utils.UPLOAD_ALLOWED_FILE_TYPE){
		return c.Err("此类型附件不允许上传")
	}

	uploadInfo := &UploadInfo{}

	m := c.Request.MultipartForm
	realPath := "";
	for fname, _ := range m.File {
		fheaders := m.File[fname]
		for i, _ := range fheaders {
			//for each fileheader, get a handle to the actual file
			file, err := fheaders[i].Open()
			defer file.Close() //close the source file handle on function return
			if err != nil {
				return c.Err(err.Error())
			}

			uploadInfo.clientName = fheaders[i].Filename
			uploadInfo.fileExt = filepath.Ext(fheaders[i].Filename)

			newFileName := uniqueId(c.User.Id,fileType,fheaders[i].Filename)+uploadInfo.fileExt
			uploadInfo.fileName = newFileName

			realPath = filepath.Join("public","upload",fileType,utils.FormatNow("200601"),newFileName)
			uploadInfo.RealPath = realPath

			dst_path := filepath.Join(revel.BasePath,realPath)
			dst, err := os.Create(dst_path)
			defer dst.Close() //close the destination file handle on function return
			defer os.Chmod(dst_path, (os.FileMode)(0644)) //limit access restrictions
			if err != nil {
				return c.Err(err.Error())
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				return c.Err(err.Error())
			}
		}
	}
	return c.OK(uploadInfo)
}

func uniqueId(uid int,fileType string,fileName string) string{
	res := fmt.Sprintf("%d",uid)+fileType+fmt.Sprintf("%d",time.Now().UnixNano())
	return utils.Md5String(res)
}