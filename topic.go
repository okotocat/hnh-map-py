package main

import "sync"

type topic struct {
	c  []chan *TileData
	mu sync.Mutex
}

// sets "topic" structure
// c = slice of the "*TileData"
// for acces mutex
func (t *topic) watch(c chan *TileData) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.c = append(t.c, c)
}

// method watch add c (dataType "*TileData") to slice t.c
// sync acces with mutex
func (t *topic) send(b *TileData) {
	t.mu.Lock()
	defer t.mu.Unlock()
	for i := 0; i < len(t.c); i++ {
		select {
		case t.c[i] <- b:
		default:
			close(t.c[i])
			t.c[i] = t.c[len(t.c)-1]
			t.c = t.c[:len(t.c)-1]
		}
	}
}

// sends b (dataType *TileData) to every channel of slice t.c. ;
// if channels busy, it closes channel and deletes it from slice
func (t *topic) close() {
	for _, c := range t.c {
		close(c)
	}
	t.c = t.c[:0]
}

// closes every channel in slice t.c.
type Merge struct {
	From, To int
	Shift    Coord
}

// sets From to int
// sets Shift to datatype Coord

//analog of topic struct
type mergeTopic struct {
	c  []chan *Merge
	mu sync.Mutex
}

func (t *mergeTopic) watch(c chan *Merge) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.c = append(t.c, c)
}
func (t *mergeTopic) send(b *Merge) {
	t.mu.Lock()
	defer t.mu.Unlock()
	for i := 0; i < len(t.c); i++ {
		select {
		case t.c[i] <- b:
		default:
			close(t.c[i])
			t.c[i] = t.c[len(t.c)-1]
			t.c = t.c[:len(t.c)-1]
		}
	}
}
func (t *mergeTopic) close() {
	for _, c := range t.c {
		close(c)
	}
	t.c = t.c[:0]
}
