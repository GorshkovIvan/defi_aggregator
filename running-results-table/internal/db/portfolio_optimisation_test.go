package db

import (
	"testing"
)

func TestPortfolioOptimisation(t *testing.T) {
	database := New()

	optimised := NewOptimisedPortfolio(&database)

	if len(optimised) == 0 {
		t.Errorf("No optimal portfolio returned")
	}

	optimised2 := OptimisePortfolio(&database)

	if len(optimised2) == 0 {
		t.Errorf("No optimal portfolio returned")
	}

	database.AddRecordfromAPI()
	optimised3 := OptimisePortfolio(&database)

	if len(optimised3) == 0 {
		t.Errorf("No optimal portfolio returned")
	}

}
