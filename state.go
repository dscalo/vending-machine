package main

type State struct {
	Screen string
	Balance float64
	Snacks *Snacks
	Selection int

}


func newState(screen string, balance float64, snacks *Snacks) *State {
	state := State{Screen: screen, Balance: balance, Snacks: snacks}
	return &state
}
