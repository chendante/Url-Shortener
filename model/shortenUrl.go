package model

import (
	. "Url-Shortener/model/base"
)

var db = make(map[string]string)
var reverseDb = make(map[string]string)

type urlPair struct {
	originUrl string
	shortUrl  string
}

func ShortenUrl(longUrl string) string {
	if shortUrl, ok := SearchShortUrlSQL(longUrl); ok{
		return shortUrl
	}
	var urlP urlPair
	urlP.shortUrl = getShortUrl(longUrl)
	urlP.originUrl = longUrl
	insertUrlSQL(urlP)
	//db[longUrl] = shortUrl
	//reverseDb[shortUrl] = longUrl
	return urlP.shortUrl
}

func insertUrlSQL(urlP urlPair) {
	Db.Exec("insert INTO url(origin_url, short_url) values(?,?)", urlP.originUrl, urlP.shortUrl)
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
	// shortUrl, ok := db [longUrl]
	urlP := new(urlPair)
	row := Db.QueryRow("select short_url from url where origin_url = ?", longUrl)
	if err := row.Scan(&urlP.shortUrl); err == nil{
		if urlP.shortUrl != ""{
			return urlP.shortUrl, true
		}
	}
	return "", false
}

func SearchOriginUrlSQL(shortUrl string) (string, bool) {
	urlP := new(urlPair)
	row := Db.QueryRow("select origin_url from url where short_url = ?", shortUrl)
	if err := row.Scan(&urlP.originUrl); err == nil{
		if urlP.originUrl != ""{
			return urlP.originUrl, true
		}
	}
	return "", false
}
