package main

import "time"

/*
	submitter is an interface with a submit method
*/
type submitter interface {
	submit(values formValues) error
}

type submitterInstance struct{}

func newSubmitter() submitter {
	return submitterInstance{}
}

func (i submitterInstance) submit(values formValues) error {
	// here we'll just mock the submit
	time.Sleep(time.Second * 2)
	return nil
}
