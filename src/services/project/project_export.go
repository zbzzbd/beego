package project

import (
	"encoding/csv"
	"log"
	"os"
	"utils"
)

type ProjectExport struct {
	filename string
	fullpath string
}

func NewProjectExport() *ProjectExport {
	return &ProjectExport{
		filename: GetFileName(),
		fullpath: utils.GetConf().String("upload::root") + utils.GetConf().String("upload::export"),
	}
}

func (self *ProjectExport) MakeFile(records [][]string) *ProjectExport {
	var err error
	f, err := os.Create(self.GetFilePath())
	w := csv.NewWriter(f)
	w.WriteAll(records)
	if err = w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
	return self
}

func GetFileName() string {
	return utils.GetRandomName()
}

func (self *ProjectExport) GetFilePath() string {
	return self.fullpath + "/" + self.filename
}
