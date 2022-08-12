package main

/*
mobileDevice is the "abstraction" interface of the bridge pattern
*/
type mobileDevice interface {
	/*
	  each abstraction will have a method that allows the abstraction to set an
	  implementation object
	*/
	installMap(mapApp)
	/*
	  the pattern also should have methods that clients will interface with
	*/
	searchAddress(string)
}
