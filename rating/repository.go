package rating

import (
	"github.com/tenahubapi/entity"
)

// RatingRepository is
type RatingRepository interface{
	Rating(id uint) (float64, []error)
	StoreRating(rating *entity.Comment) (*entity.Comment, []error)
}
