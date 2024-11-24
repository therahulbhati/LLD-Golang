package restaurant

type Restaurant struct {
	ID                  string
	Name                string
	ServiceablePincodes []string
	FoodItemName        string
	FoodItemPrice       float64
	Quantity            int
	Ratings             []Rating
}

type Rating struct {
	UserID  string
	Score   int
	Comment string
}

func (r *Restaurant) AverageRating() float64 {
	if len(r.Ratings) == 0 {
		return 0
	}
	sum := 0
	for _, rating := range r.Ratings {
		sum += rating.Score
	}
	return float64(sum) / float64(len(r.Ratings))
}
