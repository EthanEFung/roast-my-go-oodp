package main

import "fmt"

/*
  nyTimes is a concrete publisher that keeps reference to a map of readers
*/
type nyTimes struct {
	article     *article
	subscribers map[*reader]bool
}

func newNYTimes() *nyTimes {
	return &nyTimes{
		subscribers: make(map[*reader]bool, 0),
	}
}
func (pub *nyTimes) subscribe(sub *reader) {
	pub.subscribers[sub] = true
	fmt.Printf("Thank you %s for subscribing!\n", sub.name)
}
func (pub *nyTimes) unsubscribe(sub *reader) {
	delete(pub.subscribers, sub)
}
func (pub *nyTimes) notify() {
	fmt.Println("publishing...")
	fmt.Println()
	for sub := range pub.subscribers {
		/*
		   notice that the publisher only has reference to the readers who've subscribed
		*/
		sub.update(pub.article)
	}
}
func (pub *nyTimes) post(a *article) {
	fmt.Println()
	fmt.Println("posting:", a.name)
	pub.article = a
}
