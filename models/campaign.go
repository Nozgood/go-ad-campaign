package models

type Date struct {
	Day int 		`json:"dateDay" bson:"date_day"`
	Month string	`json:"dateMonth" bson:"date_month"`
	Year int		`json:"dateYear" bson:"date_year"`
}

type Campaign struct {
	Name            string	`json:"name" bson:"campaign_name"` // traduction betweeen json and bson format
	StartDate       string	`json:"startDate" bson:"campaign_start"`
	EndDate         string	`json:"endDate" bson:"campaign_end"`
	Price           int		`json:"price" bson:"campaign_price"`
	Objective       int		`json:"objective" bson:"campaign_objective"`
	PricePerDisplay float64	`json:"pricePerDisplay" bson:"campaign_pricePerDisplay"`
}
