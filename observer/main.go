package main

func main() {
  reader1 := newReader("Frank")
  reader2 := newReader("Jill")
  nyt := newNYTimes()

  nyt.subscribe(reader1)
  nyt.subscribe(reader2)

  a := &article{
    name: "Click Bait Title Here",
    content: "This just in, politics!",
  }

  nyt.post(a)
  nyt.notify()
}
