package cartrepo

import (
	"context"
	"fooddlv/cart/cartmodel"
	"fooddlv/common"
)

type Food struct {
	id          int
	name        string
	description string
	price       float32
	images      string
	status      int
}
type GetFoodStorage interface {
	GetFoods(ctx context.Context, cond map[string]interface{}, id []int) ([]Food, error)
}

type CartStorage interface {
	Create(ctx context.Context, createCartData *cartmodel.CreateCart) (int, error)
}

type createCartRepo struct {
	foodStore GetFoodStorage
	store     CartStorage
}

func NewCreateCartRepo(store CartStorage) *createCartRepo {
	return &createCartRepo{store: store}
}

func (repo *createCartRepo) AddToCart(ctx context.Context, createCartData *cartmodel.CreateCart) (result int, err error) {
	// Get foods info
	foodsId := make([]int, len(createCartData.CartItems))
	for i := range foodsId {
		foodsId[i] = createCartData.CartItems[i].FoodId
	}

	if _, err := repo.foodStore.GetFoods(ctx, nil, foodsId); err != nil {
		return 0, common.NewCustomError(err, "Item not exists", "FoodNotExists")
	}
	// Get User id,

	// otherwise handle the add to the cart
	result, err = repo.store.Create(ctx, createCartData)

	return result, nil
}
