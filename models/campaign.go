package models

type Date struct {
	Day int 		`json:"dateDay" bson:"date_day"`
	Month int	`json:"dateMonth" bson:"date_month"`
	Year int		`json:"dateYear" bson:"date_year"`
}

type Campaign struct {
	Name            string	`json:"name" bson:"campaign_name"` // traduction betweeen json and bson format
	StartDate       Date	`json:"startDate" bson:"campaign_start"`
	EndDate         Date	`json:"endDate" bson:"campaign_end"`
	Price           int		`json:"price" bson:"campaign_price"`
	Objective       int		`json:"objective" bson:"campaign_objective"`
	PricePerDisplay float64	`json:"pricePerDisplay" bson:"campaign_pricePerDisplay"`
}
