package cartrepo

import (
	"context"
	"fooddlv/cart/cartmodel"
)

//type GetFoodStorage interface {
//	GetFoods(ctx context.Context, cond map[string]interface{}, id []int) ([]Food, error)
//}

type CartStorage interface {
	Create(ctx context.Context, createCartsData *[]cartmodel.Cart) (int, error)
}

type createCartRepo struct {
	//foodStore GetFoodStorage
	store CartStorage
}

func NewCreateCartRepo(store CartStorage) *createCartRepo {
	return &createCartRepo{store: store}
}

func (repo *createCartRepo) AddToCart(ctx context.Context, createCartsData *[]cartmodel.Cart) (result int, err error) {
	// Get foods info
	//foodsId := make([]int, len(createCartsData))
	//for i := range foodsId {
	//	foodsId[i] = createCartsData[i].FoodId
	//}

	//if _, err := repo.foodStore.GetFoods(ctx, nil, foodsId); err != nil {
	//	return 0, common.NewCustomError(err, "Item not exists", "FoodNotExists")
	//}
	// Get User id,

	// otherwise handle the add to the cart
	result, err = repo.store.Create(ctx, createCartsData)

	return result, nil
}
