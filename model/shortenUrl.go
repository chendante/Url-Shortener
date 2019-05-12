package model

import (
	. "Url-Shortener/model/base"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strings"
)

type Url struct {
	Id			uint	`gorm:"primary_key"`
	OriginUrl	string
	ShortUrl	string
	Visits		int
}
func (Url) TableName() string {
	return "url"
}

func ShortenUrl(longUrl string) string {
	if shortUrl, ok := SearchShortUrlSQL(longUrl); ok{
		return shortUrl
	}
	var urlP Url
	urlP.ShortUrl = getShortUrl(longUrl)
	urlP.OriginUrl = longUrl
	insertUrlSQL(urlP)
	return urlP.ShortUrl
}

func insertUrlSQL(urlP Url) {
	Db.Create(&urlP)
	if !Db.NewRecord(urlP){
		fmt.Print("yes")
	}
}

func getShortUrl(longUrl string) string {
	shortUrlList, _ := Transform(longUrl)
	var res string
	for _, v := range shortUrlList{
		if _, ok := SearchOriginUrlSQL(v); !ok{
			res = v
			break
		}
	}
	if res == ""{
		res = getShortUrl(longUrl + "a")
	}
	return res
}

func SearchShortUrlSQL(longUrl string) (string, bool) {
	urlP := Url{OriginUrl:longUrl}
	var result Url
	Db.Where(&urlP).First(&result)
	if result.Id == 0{
		return "", false
	} else {
		return result.ShortUrl, true
	}
}

func SearchOriginUrlSQL(shortUrl string) (string, bool) {
	urlP := Url{ShortUrl:shortUrl}
	var res Url
	Db.Where(&urlP).First(&res)
	if res.Id == 0{
		return "", false
	} else {
		return res.OriginUrl, true
	}
}

func SearchOriginUrlRedis(shortUrl string) (string, bool) {
	originUrl, err := redis.String(MRedis.Do("GET", shortUrl))
	if err != nil {
		return "", false
	} else {
		return originUrl, true
	}
}

func SearchOriginUrl(shortUrl string) (string, bool) {
	originUrl, ok := SearchOriginUrlRedis(shortUrl)
	if !ok{
		originUrl, ok := SearchOriginUrlSQL(shortUrl)
		if ok{
			insertUrlRedis(Url{ShortUrl:shortUrl, OriginUrl:originUrl})
		}
		return originUrl, ok
	}
	return originUrl, true
}

func insertUrlRedis(urlP Url) {
	_, _ = MRedis.Do("SET", urlP.ShortUrl, urlP.OriginUrl)
}

//func updateUrlVisits(origin_url string) {
//
//}

func AddScript(url string) string {
	var res = "<head><meta http-equiv=\"refresh\" content=\"0;url="
	if !strings.HasPrefix(url, "http"){
		res += "https://"
	}
	res += url
	res += "\"></head>"
	return res
}
