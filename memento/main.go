package main

import "fmt"

/*
  An originator is a struct that has an internal state
  and can produce snapshots of that internal state called "Mementos"
*/
type originator struct {
  state string
}
func (o *originator) save() *memento {
  return &memento{state: o.state}
}
func (o *originator) restore(m *memento) {
  o.state = m.getState()
}
func (o *originator) setState(state string) {
  o.state = state
}
func(o *originator) getState() string {
  return o.state
}

/*
  A memento is a snapshot object that represents a snapshot of
  the originators state. Note, that the memento has no "setter"
  since the state should only be set when the memento is initialized.
*/
type memento struct {
  state string
}
func (m *memento) getState() string {
  return m.state
}

/*
  A caretaker stores a stack of mementos and has means of adding to the
  stack as well as accessing all the menentos that were added previously.
*/
type caretaker struct {
  mementos []*memento
}
func(c *caretaker) add(m *memento) {
  c.mementos = append(c.mementos, m) 
}
func(c *caretaker) get(index int) *memento {
  return c.mementos[index]
}

func main() {
  ct := &caretaker{
    mementos: make([]*memento, 0),
  }
  o := &originator{
    state: "A",
  }
  fmt.Printf("initial state: %s\n", o.getState())
  ct.add(o.save())
  
  o.setState("B")
  fmt.Printf("state: %s\n", o.getState())
  ct.add(o.save())

  o.setState("C")
  fmt.Printf("state: %s\n", o.getState())
  ct.add(o.save())

  o.restore(ct.get(1))
  fmt.Printf("restored: %s\n", o.getState())

  o.restore(ct.get(0))
  fmt.Printf("restored: %s\n", o.getState())

}
