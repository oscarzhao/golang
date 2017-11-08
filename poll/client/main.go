package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	server string

	times        int64
	frequencySec int64
)

// curl 'http://a4pbww.v.vote8.cn/' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: en-US,zh-CN;q=0.8,zh;q=0.5,en;q=0.3' -H 'Connection: keep-alive' -H 'Content-Type: application/x-www-form-urlencoded' -H 'Cookie: __cfduid=ddb54b33dc145860a2170a8d71323b82e1499827983; ASP.NET_SessionId=4rinftdvdvv0wgxvmgo41hvm; Vote.CheckedOption.2805658=6643052; Vote.CheckedOption.2805603=6642374; Vote.CheckedOption.2805651=; Vote.CheckedOption.2805656=6646616; Vote.CheckedOption.2805657=6643141' -H 'DNT: 1' -H 'Host: a4pbww.v.vote8.cn' -H 'Referer: http://a4pbww.v.vote8.cn/' -H 'Upgrade-Insecure-Requests: 1' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:54.0) Gecko/20100101 Firefox/54.0' --data '__EVENTTARGET=ctl00%24cphMainContent%24lbtnVote&__EVENTARGUMENT=&__VIEWSTATE=%2FwEPDwUJNzI0MDU4ODE4ZGQv2TCzfpTJWODgulZwjDDa63WZ9Q%3D%3D&__VIEWSTATEGENERATOR=9A48AF2D&ctl00%24cphMainContent%24rptTopicList%24ctl00%24tbOptionSearchInput=&VoteOption_2805603=6642374&ctl00%24cphMainContent%24rptTopicList%24ctl00%24rptOptions%24ctl12%24hiddenOptionDisabledMessage=%E6%8A%95%E7%A5%A8%E6%95%B0%E9%87%8F%E5%B7%B2%E8%BE%BE%E4%B8%8A%E9%99%90&ctl00%24cphMainContent%24rptTopicList%24ctl00%24hiddenTopicID=2805603&ctl00%24cphMainContent%24rptTopicList%24ctl01%24tbOptionSearchInput=&ctl00%24cphMainContent%24rptTopicList%24ctl01%24hiddenTopicID=2805651&ctl00%24cphMainContent%24rptTopicList%24ctl02%24tbOptionSearchInput=&ctl00%24cphMainContent%24rptTopicList%24ctl02%24rptOptions%24ctl07%24hiddenOptionDisabledMessage=%E6%8A%95%E7%A5%A8%E6%95%B0%E9%87%8F%E5%B7%B2%E8%BE%BE%E4%B8%8A%E9%99%90&VoteOption_2805656=6646616&ctl00%24cphMainContent%24rptTopicList%24ctl02%24hiddenTopicID=2805656&ctl00%24cphMainContent%24rptTopicList%24ctl03%24tbOptionSearchInput=&VoteOption_2805657=6643141&ctl00%24cphMainContent%24rptTopicList%24ctl03%24hiddenTopicID=2805657&ctl00%24cphMainContent%24rptTopicList%24ctl04%24tbOptionSearchInput=&VoteOption_2805658=6643052&ctl00%24cphMainContent%24rptTopicList%24ctl04%24hiddenTopicID=2805658&ctl00%24cphMainContent%24hiddenWeixinUserInfoJson=%7B%22city%22%3A%22Haidian%22%2C%22province%22%3A%22Beijing%22%2C%22country%22%3A%22CN%22%2C%22openid%22%3A%22ovp7djpCb9ERW0Fa8bOS3FTHadoY%22%2C%22unionid%22%3A%22oypjDt7VMhN5Sncr3WAHD0e57FoA%22%2C%22nickname%22%3A%22Oscar%22%2C%22headimgurl%22%3A%22http%3A%2F%2Fwx.qlogo.cn%2Fmmopen%2F0AXDYELkFI5UcS55ic0KJyaLufOxPbcs6Qk7ftCUP8GtOiau9bjmXBcrraia3yM2zYoHA3P0Ar1vDw6aSEyWTBkL3cibzveEfKpx%2F0%22%7D&ctl00%24cphMainContent%24hiddenWeixinUserInfoJsonEncode=d2e691df21417566f77387749fbc9701&hiddenVote8ClickValidateCode=1499829562%2C4583e970cf928b9d1b5c33cb70d3564a&ctl00%24cphMainContent%24ucVerifyCode%24hiddenVerifyCodeModeInfo=8%2C50af00639360fdccf2b6e9d32d397e87&ctl00%24cphMainContent%24hiddenRefererUrl=http%3A%2F%2Fa4pbww.v.vote8.cn%2Fm%2FShortcut%2F6643052&ctl00%24cphMainContent%24hiddenTimeStampEncodeString=1499829344%2C174ae87791be90143dfceaea1cb76ba3&ctl00%24cphMainContent%24hiddenLatitude=&ctl00%24cphMainContent%24hiddenLongitude=&ctl00%24cphMainContent%24hiddenGeoLocationEncode='

func init() {
	flag.StringVar(&server, "server", "http://a4pbww.v.vote8.cn/", "server url")
	flag.Int64Var(&times, "times", 1, "try to send request for how many times")
	flag.Int64Var(&frequencySec, "frequency", 20, "连续两次请求的最大发送间隔（单位：秒），实际发送间隔在 0和这个值之间")

	rand.Seed(time.Now().Unix())
}

func SendRequest() {
	req, _ := http.NewRequest("POST", server, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:54.0) Gecko/20100101 Firefox/54.0")
	_ = req
}

func main() {
	for i := int64(0); i < times; i++ {
		sleepSec := rand.Int63n(frequencySec)
		log.Printf("will poll after %d minutes\n", sleepSec)
		time.Sleep(time.Duration(sleepSec) * time.Second)
	}
}
