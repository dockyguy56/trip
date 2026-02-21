package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RideFareModel struct {
	ID                 primitive.ObjectID
	UserID             primitive.ObjectID
	PackageSlug        string
	TTotalPriceInCents float64
}
