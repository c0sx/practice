package leetcode1912

import (
	"container/heap"
)

type Movie struct {
	shop  int
	movie int
	price int
	index int // index in the heap
}

type MovieHeap []*Movie

func (mh MovieHeap) Len() int {
	return len(mh)
}

func (mh MovieHeap) Less(i, j int) bool {
	if mh[i].price == mh[j].price {
		if mh[i].shop == mh[j].shop {
			return mh[i].movie < mh[j].movie
		}

		return mh[i].shop < mh[j].shop
	}

	return mh[i].price < mh[j].price
}

func (mh MovieHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]

	// hack to watch for index changes
	mh[i].index = i
	mh[j].index = j
}

func (mh *MovieHeap) Push(x interface{}) {
	n := len(*mh)
	item := x.(*Movie)
	item.index = n // hack to watch for index changes
	*mh = append(*mh, item)
}

func (mh *MovieHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*mh = old[0 : n-1]
	x.index = -1 // for safety

	return x
}

type MovieRentingSystem struct {
	unrented map[int]*MovieHeap // movie
	rented   MovieHeap
	details  map[int]map[int]*Movie
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
	unrented := map[int]*MovieHeap{}
	rented := MovieHeap{}
	details := map[int]map[int]*Movie{}

	for _, entry := range entries {
		m := &Movie{
			shop:  entry[0],
			movie: entry[1],
			price: entry[2],
		}

		if _, ok := unrented[entry[1]]; !ok {
			unrented[entry[1]] = &MovieHeap{}
		}

		if _, ok := details[entry[0]]; !ok {
			details[entry[0]] = map[int]*Movie{}
		}

		heap.Push(unrented[entry[1]], m)
		details[entry[0]][entry[1]] = m
	}

	return MovieRentingSystem{
		unrented: unrented,
		rented:   rented,
		details:  details,
	}
}

// Search: Finds the cheapest 5 shops that have an unrented copy of a given movie.
// The shops should be sorted by price in ascending order,
// and in case of a tie, the one with the smaller shopi should appear first.
// If there are less than 5 matching shops, then all of them should be returned.
// If no shop has an unrented copy, then an empty list should be returned.
func (mrs *MovieRentingSystem) Search(movie int) []int {
	movies, ok := mrs.unrented[movie]
	if !ok || movies.Len() == 0 {
		return []int{}
	}

	candidates := []*Movie{}
	for len(candidates) < 5 && movies.Len() > 0 {
		m := heap.Pop(movies).(*Movie)
		candidates = append(candidates, m)
	}

	result := []int{}
	for i := 0; i < len(candidates); i++ {
		result = append(result, candidates[i].shop)
		heap.Push(movies, candidates[i])
	}

	return result
}

// Rent: Rents an unrented copy of a given movie from a given shop.
func (mrs *MovieRentingSystem) Rent(shop int, movie int) {
	details, ok := mrs.details[shop][movie]
	if !ok || details.index == -1 {
		return
	}

	movies, ok := mrs.unrented[movie]
	if !ok || details.index >= movies.Len() {
		return
	}

	heap.Remove(movies, details.index)
	heap.Push(&mrs.rented, details)
}

// Drop: Drops off a previously rented copy of a given movie at a given shop.
func (mrs *MovieRentingSystem) Drop(shop int, movie int) {
	details, ok := mrs.details[shop][movie]
	if !ok || details.index == -1 {
		return
	}

	heap.Remove(&mrs.rented, details.index)
	heap.Push(mrs.unrented[movie], details)
}

// Report: Returns the cheapest 5 rented movies
// (possibly of the same movie ID) as a 2D list res where res[j] = [shopj, moviej]
// describes that the jth cheapest rented movie moviej was rented from the shop shopj.
// The movies in res should be sorted by price in ascending order, and in case of a tie,
// the one with the smaller shopj should appear first, and if there is still tie,
// the one with the smaller moviej should appear first. If there are fewer than 5 rented movies,
// then all of them should be returned. If no movies are currently being rented,
// then an empty list should be returned.
func (mrs *MovieRentingSystem) Report() [][]int {
	candidates := []*Movie{}
	result := [][]int{}

	for len(candidates) < 5 && mrs.rented.Len() > 0 {
		m := heap.Pop(&mrs.rented).(*Movie)
		candidates = append(candidates, m)
	}

	for _, m := range candidates {
		result = append(result, []int{m.shop, m.movie})
		heap.Push(&mrs.rented, m)
	}

	return result
}

/**
 * Your MovieRentingSystem object will be instantiated and called as such:
 * obj := Constructor(n, entries);
 * param_1 := obj.Search(movie);
 * obj.Rent(shop,movie);
 * obj.Drop(shop,movie);
 * param_4 := obj.Report();
 */
