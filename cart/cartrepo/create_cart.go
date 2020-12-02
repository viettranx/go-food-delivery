package cartrepo

import (
	"context"
	"fooddlv/cart/cartmodel"
	"fooddlv/common"
)

type CartStorage interface {
	Create(ctx context.Context, createCartData *cartmodel.CartCreation) (int, error)
}

type createCartRepo struct {
	store CartStorage
}
type Food struct {
	id          int
	name        string
	description string
	price       float32
	images      string
	status      int
}

func NewCreateCartRepo(store CartStorage) *createCartRepo {
	return &createCartRepo{store: store}
}

func (repo *createCartRepo) AddToCart(ctx context.Context, createCartData *cartmodel.CartCreation) (result int, err error) {
	//var item interface{}
	// find item in foods if exits
	// item, err := repo.store.FindFoodsByCondition(ctx, map[string]interface{}{"id": createCartData.FoodId})
	// so Dummy the Food response
	item := Food{
		id:          1,
		name:        "Dummy",
		description: "Dummy des",
		price:       100,
		images:      "http://...png",
	}

	// If food item not exit, or out of stock, should return error to user
	if item == nil {
		return 0, common.NewCustomError(err, "Item not exists", "FoodNotExists")
	}
	// otherwise handle the add to the cart
	result, err = repo.store.Create(ctx, createCartData)

	return result, nil
}
