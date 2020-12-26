package main

// Module: Food
// User request list of foods. The system should show food with:
// - has_liked (bool): if user liked?
// - liked_count (int): count of total liked by users

// Solution:
// API: /v1/foods
// Repo layer:
// type LikeStore interface {
//	GetUserLikeFood(context, foodIds, userId) (map[int]bool, error)
//}

// 1. Fetch list of food as normal
// 2. Call to module "User Like Food" to get array of "UserLikeFood" model.
// 2.1 Query by array of food_id and current user_id
// 3. Assign to participant food

// Module: User Like Food
// Model UserLikeFood: (food_id, user_id, created_at, updated_at)
// Storage layer:
// Func: GetUserLikeFood(context, foodIds, userId) (map[int]bool, error)

// foods []Food, ulfs []UserLikeFood
// ulfs []UserLikeFood => mapUserLiked map[int]bool
//
// for _, item := range ulfs {
//	mapUserLiked[item.Id] = true
// }
//
// for i := range foods {
//  food := foods[i]
//  foods[i].HasLiked = mapUserLiked[food.Id]
//}
