package job

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	JobCode_NoneModify      = 0
	JobCode_BussinessModify = 1
	JobCode_ProduceModify   = 2
)

type JobCode struct {
	date               string
	companyJobCnt      string
	bussinessModifyNum int
	produceModifyNum   int
	companyCode        string
	codeType           uint8
	isNew              bool
}

func NewJobCode(code, companyCode string, codeType uint8) *JobCode {
	jobCode := JobCode{
		companyCode: companyCode,
		codeType:    codeType,
	}
	if code == "" {
		jobCode.isNew = true
	}

	jc := &jobCode
	jc.parseDate(code).parseCompany(code).parseBussiness(code).parseProduce(code)

	return jc
}

func (jobCode *JobCode) parseDate(code string) *JobCode {
	if len(code) < 9 {
		jobCode.date = time.Now().Format("20060102")
		return jobCode
	}
	jobCode.date = code[0:8]

	return jobCode
}

func (jobCode *JobCode) parseCompany(code string) *JobCode {
	if len(code) < 9 {
		return jobCode
	}

	pos := strings.Index(code, "-")
	// 不存在
	if pos < 0 {
		jobCode.companyJobCnt = code[8:]
	} else if pos > 8 {
		jobCode.companyJobCnt = code[8:pos]
	}

	return jobCode
}

func (jobCode *JobCode) parseBussiness(code string) *JobCode {
	reg := regexp.MustCompile(`-Y[0-9]+`)
	words := reg.FindString(code)
	if len(words) < 2 {
		return jobCode
	}

	reg = regexp.MustCompile(`[0-9]+$`)
	jobCode.bussinessModifyNum, _ = strconv.Atoi(reg.FindString(words))

	return jobCode
}

func (jobCode *JobCode) parseProduce(code string) *JobCode {
	reg := regexp.MustCompile(`-Z[0-9]+`)
	words := reg.FindString(code)
	if len(words) < 2 {
		return jobCode
	}

	reg = regexp.MustCompile(`[0-9]+$`)
	jobCode.produceModifyNum, _ = strconv.Atoi(reg.FindString(words))

	return jobCode
}

func (jobCode *JobCode) GetCode() (code string) {
	if jobCode.isNew {
		code = jobCode.date + jobCode.companyCode + strconv.FormatInt(GetTodayJobCount()+1, 10)
	} else {
		code = jobCode.date + jobCode.companyJobCnt
		if jobCode.codeType == JobCode_ProduceModify {
			jobCode.produceModifyNum += 1
		} else if jobCode.codeType == JobCode_BussinessModify {
			jobCode.bussinessModifyNum += 1
		}

		if jobCode.bussinessModifyNum > 0 {
			code += "-Y" + strconv.Itoa(jobCode.bussinessModifyNum)
		}

		if jobCode.produceModifyNum > 0 {
			code += "-Z" + strconv.Itoa(jobCode.produceModifyNum)
		}
	}

	return
}
