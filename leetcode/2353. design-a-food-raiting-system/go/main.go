package leetcode2353

import (
	"container/heap"
)

type Food struct {
	name   string
	rating int
}

type FoodHeap []Food

func (fh FoodHeap) Len() int { return len(fh) }
func (fh FoodHeap) Less(i, j int) bool {
	if fh[i].rating == fh[j].rating {
		return fh[i].name < fh[j].name
	}
	return fh[i].rating > fh[j].rating
}
func (fh FoodHeap) Swap(i, j int) { fh[i], fh[j] = fh[j], fh[i] }

func (fh *FoodHeap) Push(x any) {
	*fh = append(*fh, x.(Food))
}

func (fh *FoodHeap) Pop() any {
	old := *fh
	n := len(old)
	x := old[n-1]
	*fh = old[0 : n-1]
	return x
}

type FoodRatings struct {
	ratings       map[string]int
	foodByCuisine map[string]*FoodHeap
	cuisineOfFood map[string]string
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	ratingsMap := make(map[string]int)
	foodByCuisine := make(map[string]*FoodHeap)
	cuisineOfFood := make(map[string]string)

	for i, food := range foods {
		ratingsMap[food] = ratings[i]
		cuisine := cuisines[i]
		cuisineOfFood[food] = cuisine

		if _, exists := foodByCuisine[cuisine]; !exists {
			foodByCuisine[cuisine] = &FoodHeap{}
		}

		heap.Push(foodByCuisine[cuisine], Food{name: food, rating: ratings[i]})
	}

	return FoodRatings{
		ratings:       ratingsMap,
		foodByCuisine: foodByCuisine,
		cuisineOfFood: cuisineOfFood,
	}
}

func (fr *FoodRatings) ChangeRating(food string, newRating int) {
	fr.ratings[food] = newRating
	cuisine := fr.cuisineOfFood[food]
	heap.Push(fr.foodByCuisine[cuisine], Food{name: food, rating: newRating})
}

func (fr *FoodRatings) HighestRated(cuisine string) string {
	foodHeap := fr.foodByCuisine[cuisine]
	for {
		topFood := (*foodHeap)[0]
		if fr.ratings[topFood.name] == topFood.rating {
			return topFood.name
		}

		heap.Pop(foodHeap)
	}
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
