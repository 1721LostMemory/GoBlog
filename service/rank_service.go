package service

import (
	"context"
	"goblog/config"
	"goblog/models"
)

func RankByPosts(top int64) ([]models.UserRank, error) {
	zs, err := config.RedisDB.ZRevRangeWithScores(context.Background(), "rank:user:post", 0, top-1).Result()
	if err != nil {
		return nil, err
	}

	var ranks []models.UserRank
	for i, z := range zs {
		ranks = append(ranks, models.UserRank{
			Username:  z.Member.(string),
			PostCount: uint(z.Score),
			Rank:      uint(i + 1),
		})
	}
	return ranks, nil
}
