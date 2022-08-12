/*
In the bridge pattern, a client works with an abstraction layer.
The abstractions delegate responsibilities to an implementation hierarchy.
*/
package main

func main() {
	gmapsApp := &gmaps{}
	appleMapsApp := &appleMaps{}
	/*
	  in our example each mobile device is an abstraction and the mapping apps
	  are the implementations that can be used interchangeably.
	*/
	iphoneDevice := &iphone{}
	iphoneDevice.installMap(appleMapsApp)

	iphoneDevice.searchAddress("1000 Vin Scully Ave, Los Angeles, CA 90012")

	galexySDevice := &galaxyS{}
	galexySDevice.searchAddress("1000 Vin Scully Ave, Los Angeles, CA 90012")

	galexySDevice.installMap(gmapsApp)
	galexySDevice.searchAddress("1000 Vin Scully Ave, Los Angeles, CA 90012")
}
