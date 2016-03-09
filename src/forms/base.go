package forms

type Base struct {
	errMsg map[string]error
}

func (b *Base) AddError(field string, err error) {
	b.errMsg[field] = err
}

func (b *Base) ClearErrors() {
	b.errMsg = make(map[string]error)
}

func (b *Base) Errors() (errs map[string]string) {
	errs = make(map[string]string)
	for field, err := range b.errMsg {
		errs[field] = err.Error()
	}
	return
}
