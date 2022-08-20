package main

/*
	formFacade keeps reference to a form implementation
*/
type formFacade struct {
	form *form
}

func newFormFacade() *formFacade {
	return &formFacade{
		form: newForm(),
	}
}

/*
	the formFacade submit is where values are submitted to the form, validated,
	and submitted
*/
func (facade formFacade) submit(values formValues) error {
	facade.form.setValues(values)
	err := facade.form.validator.validate(facade.form.values)
	if err != nil {
		return err
	}
	err = facade.form.submitter.submit(facade.form.values)
	if err != nil {
		return err
	}
	return nil
}
