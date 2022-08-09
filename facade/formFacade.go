package main

type formFacade struct {
	form *form
}

func newFormFacade() *formFacade {
	return &formFacade{
	  form: newForm(),
	}
}

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