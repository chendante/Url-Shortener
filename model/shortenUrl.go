package model

var db = make(map[string]string)
var reverseDb = make(map[string]string)

func ShortenUrl(longUrl string) string {
	if shortUrl, ok := SearchShortUrl(longUrl); ok{
		return shortUrl
	}
	shortUrl := getShortUrl(longUrl)
	db[longUrl] = shortUrl
	reverseDb[shortUrl] = longUrl
	return shortUrl
}

func getShortUrl(longUrl string) string {
	shortUrlList, _ := Transform(longUrl)
	var res string
	for _, v := range shortUrlList{
		if _, ok := SearchOriginUrl(v); !ok{
			res = v
			break
		}
	}
	if res == ""{
		res = getShortUrl(longUrl + "a")
	}
	return res
}

func SearchShortUrl(longUrl string) (string, bool) {
	shortUrl, ok := db [longUrl]
	return shortUrl, ok
}

func SearchOriginUrl(shortUrl string) (string, bool) {
	originUrl, ok := reverseDb [shortUrl]
	return originUrl, ok
}
