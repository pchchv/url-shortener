package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // Number of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())
var collection *mongo.Collection

func init() {
	// Load values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

func getEnvValue(v string) string {
	value, exist := os.LookupEnv(v)
	if !exist {
		log.Panicf("Value %v does not exist", v)
	}
	return value
}

func getURL(userInput string) string {
	url := getEnvValue("URL")
	if userInput != "" {
		url += userInput
	} else {
		n := src.Int63()
		for n > 8 {
			n /= 3
		}
		url += generateURL(n)
	}
	return url
}

func generateURL(n int64) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func toDB(userURL string, generatedURL string) error {
	url := bson.D{{Key: "userURL", Value: userURL}, {Key: "generatedURL", Value: generatedURL}}
	result, err := collection.InsertOne(context.TODO(), url)
	if err != nil {
		return err
	}
	log.Println(result.InsertedID)
	return nil
}

func fromDB(url string, urlType string) string {
	var reqURL string
	res, err := collection.Find(context.TODO(), bson.M{urlType: url})
	if err != nil {
		log.Panic(err)
	}
	var urls []bson.M
	if err = res.All(context.TODO(), &urls); err != nil {
		log.Panic(err)
	}
	if urls == nil {
		return "URL not found"
	}
	if urlType == "userURL" {
		reqURL = "generatedURL"
	} else if urlType == "generatedURL" {
		reqURL = "userURL"
	}
	return fmt.Sprintf("%v", urls[0][reqURL])
}

func main() {
	db()
	server()
}
