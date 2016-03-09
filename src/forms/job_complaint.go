package forms

type JobComplaint struct {
	jobId string
}

func (jc *JobComplaint) Valid() (err error) {
	return nil
}

func (jc *JobComplaint) isJobIdValid() bool {
	return true
}
