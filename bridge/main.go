package main

/*
  the bridge pattern is a design meant
  to break down large structs or muddled
  business logic into smaller structs
  that can be managed independently of each other
*/
func main() {
  gmapsApp := &gmaps{}
  appleMapsApp := &appleMaps{}

  iphoneDevice := &iphone{}
  iphoneDevice.installMap(appleMapsApp)

  iphoneDevice.searchAddress("1000 Vin Scully Ave, Los Angeles, CA 90012")

  galexySDevice := &galexyS{}
  galexySDevice.searchAddress("1000 Vin Scully Ave, Los Angeles, CA 90012")

  galexySDevice.installMap(gmapsApp)
  galexySDevice.searchAddress("1000 Vin Scully Ave, Los Angeles, CA 90012")
}
