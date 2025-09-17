package leetcode2353

import "testing"

func TestFoodRatings(t *testing.T) {
	foods := []string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}
	cuisines := []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}
	ratings := []int{9, 12, 8, 15, 14, 7}
	foodRatings := Constructor(foods, cuisines, ratings)

	if highest := foodRatings.HighestRated("korean"); highest != "kimchi" {
		t.Errorf("Expected 'kimchi', got '%s'", highest)
	}

	foodRatings.ChangeRating("sushi", 16)

	if highest := foodRatings.HighestRated("japanese"); highest != "sushi" {
		t.Errorf("Expected 'sushi', got '%s'", highest)
	}

	foodRatings.ChangeRating("ramen", 16)

	if highest := foodRatings.HighestRated("japanese"); highest != "ramen" {
		t.Errorf("Expected 'ramen', got '%s'", highest)
	}
}
