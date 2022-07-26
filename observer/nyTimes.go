package main

import "fmt"

type nyTimes struct {
  article *article
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
  if _, exists := pub.subscribers[sub]; exists {
    delete(pub.subscribers, sub)
  }
}
func (pub *nyTimes) notify() {
  fmt.Println("publishing...")
  fmt.Println()
  for sub := range pub.subscribers {
    sub.update(pub.article)
  }
}
func (pub *nyTimes) post(a *article) {
  fmt.Println()
  fmt.Println("posting:", a.name)
  pub.article = a
}

