package main

type form struct {
	values formValues
	validator validator
	submitter submitter
	submitting bool
	submitted bool
	touched map[string]bool
}

func newForm() *form {
	return &form{
		validator: newValidator(),
		submitter: newSubmitter(),
	}
}

func (f *form) setValues(values formValues) {
	f.values = values
}

