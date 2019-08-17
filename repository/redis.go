package repository

import (
	"github.com/go-redis/redis"
	"os"
	"time"
)

var client *redis.Client

func init(){
	redisURI := os.Getenv("REDIS_URI")
	client = redis.NewClient(&redis.Options{
		Addr:     redisURI,
		Password: "",
		DB:       0,
	})
}

func SetOTP(phone string, otp string) error{
	err := client.Set(phone, otp, 600*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

func ValidateOTP(phone string, otp string) (bool, error){
	storedOTP, err := client.Get(phone).Result()

	if err != nil{
		return false, err
	}

	if otp != storedOTP{
		return false, nil
	}

	return true, nil
}
