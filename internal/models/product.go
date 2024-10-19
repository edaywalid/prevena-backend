package models

type Product struct {
	Barcode     string       `json:"barcode" bson:"_id"`
	Name        string       `json:"name" bson:"name"`
	Brand       string       `json:"brand" bson:"brand"`
	Ingredients []Ingredient `json:"ingredients" bson:"ingredients"`
	OverallRisk float64      `json:"overall_risk" bson:"overall_risk"`
}

type Ingredient struct {
	ID        string  `json:"id" bson:"_id"`
	Name      string  `json:"name" bson:"name"`
	RiskScore float64 `json:"risk_score" bson:"risk_score"`
	RiskType  string  `json:"risk_type" bson:"risk_type"`
}
