package main

func main() {
	/*
	  in our example we have a collection of readers that want to receive messages from
	  the New York Timews
	*/
	reader1 := newReader("Frank")
	reader2 := newReader("Jill")
	nyt := newNYTimes()

	/*
	  the observer pattern allows the reader objects to subscribe to the publisher (nyt)
	*/
	nyt.subscribe(reader1)
	nyt.subscribe(reader2)

	a := &article{
		name:    "Click Bait Title Here",
		content: "This just in, politics!",
	}

	nyt.post(a)
	/*
	  once an article has been posted, then the publisher can notify only the subscribers
	*/
	nyt.notify()
}
