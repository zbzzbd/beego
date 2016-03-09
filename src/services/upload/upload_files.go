package upload

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"models"
	"net/http"
	"os"
	"path"
	"utils"
)

type UploadFileType uint

const (
	_                 UploadFileType = iota
	FT_Create                        // 创建作业
	FT_Modify                        // 修改作业
	FT_Valid                         // 审核作业
	FT_Claim                         // 认领作业
	FT_Submit                        // 认领作业
	FT_Complain                      // 作业投诉
	FT_ComplainReplay                //投诉回复
)

const (
	StaticPathPhoto  = "photo"
	ErrorMissingFile = "无法读取上传文件"
	ErrorSavingFile  = "服务器保存文件错误"
)

type UploadFile struct {
	request  *http.Request
	formName string
	savePath string
	url      string
	cdn      string
	files    []models.File
	fileType UploadFileType
	relId    uint
}

func NewUploadFile(request *http.Request, formName, confKey string) *UploadFile {
	return &UploadFile{
		request:  request,
		formName: formName,
		savePath: utils.GetConf().String("upload::root") + utils.GetConf().String("upload::"+confKey),
		url:      path.Join(StaticPathPhoto, utils.GetConf().String("upload::"+confKey)),
	}
}

func (uf *UploadFile) SetCdn(cdn string) *UploadFile {
	uf.cdn = cdn
	return uf
}

func (uf *UploadFile) SetFileType(fileType UploadFileType) *UploadFile {
	uf.fileType = fileType
	return uf
}

func (uf *UploadFile) SetRelId(relId uint) *UploadFile {
	uf.relId = relId
	return uf
}

func (uf *UploadFile) Do() error {
	files, ok := uf.request.MultipartForm.File[uf.formName]
	fmt.Printf("%v", files)
	if !ok {
		return errors.New(ErrorMissingFile)
	}

	for _, file := range files {
		uf.uploadFileToLocal(file)
	}

	return uf.saveFileToDb()
}

func (uf *UploadFile) uploadFileToLocal(fileHeader *multipart.FileHeader) (err error) {
	fullpath, name := uf.getPath(fileHeader.Filename)

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	if err := uf.saveToLocal(file, fullpath); err != nil {
		return err
	}

	uf.files = append(uf.files, models.File{
		Name:  fileHeader.Filename,
		Url:   uf.cdn + "/" + path.Join(uf.url, name),
		Type:  uint(uf.fileType),
		RelId: uf.relId,
	})

	return nil
}

func (uf *UploadFile) saveToLocal(file io.Reader, fullpath string) (err error) {
	f, err := os.OpenFile(fullpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)

	return
}

func (uf *UploadFile) getPath(filename string) (fullpath, name string) {
	name = utils.GetRandomName()
	ext := path.Ext(filename)
	if ext != "" && ext != "." {
		name += ext
	}

	fullpath = path.Join(uf.savePath, name)

	return
}

func (uf *UploadFile) saveFileToDb() error {
	_, err := models.GetDB().InsertMulti(len(uf.files), uf.files)

	return err
}
