package export

type Exporter interface {
	GetOutFileName() string
	GetExportData() [][]string
	GetOutStyleFormat() string
}
