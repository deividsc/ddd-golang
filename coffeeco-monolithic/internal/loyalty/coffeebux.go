package loyalty

import (
	coffeeco "coffeeco-monolithic/internal"
	"coffeeco-monolithic/internal/store"

	"github.com/google/uuid"
)

type Coffeebux struct {
	ID                                    uuid.UUID
	store                                 store.Store
	coffeelover                           coffeeco.CoffeeLover
	FreeDrinkAvailable                    int
	RemainingDrinkPurchasesUntilFreeDrink int
}
