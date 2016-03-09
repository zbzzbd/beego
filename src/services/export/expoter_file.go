package export

import (
	"encoding/csv"
	"log"
	"os"
	"utils"
)

type ExporteFile struct {
	Exporter Exporter
	filename string
	fullpath string
}

func NewExporteFile(ep Exporter) *ExporteFile {
	return &ExporteFile{
		Exporter: ep,
		filename: Getfilename(),
		fullpath: utils.GetConf().String("upload::root") + utils.GetConf().String("upload::export"),
	}
}

func Getfilename() string {
	return utils.GetRandomName()
}

func (ex *ExporteFile) GetFilePath() string {
	return ex.fullpath + "/" + ex.filename
}

func (ex *ExporteFile) MakeCsvFile(records [][]string) *ExporteFile {
	var err error
	f, err := os.Create(ex.GetFilePath())
	w := csv.NewWriter(f)
	w.WriteAll(records)
	if err = w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
	return ex
}
