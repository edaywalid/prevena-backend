package dto

type AnalyzeRequest struct {
	Ingerdients []string `json:"ingerdients"`
}

type AnalyzedIngredient struct {
	Name      string  `json:"name"`
	RiskScore float64 `json:"risk_score"`
	RiskType  string  `json:"risk_type"`
}

type AnalyzedResponse struct {
	Ingredients []AnalyzedIngredient `json:"ingredients"`
	OverallRisk float64              `json:"overall_risk"`
}
