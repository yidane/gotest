package main

import (
//"fmt"
//"log"
//"runtime"
//"time"
//
//"github.com/garyburd/redigo/redis"
)

//LocationInfo store 302 info
type LocationInfo struct {
	Key      string
	Location string
}

var (
	codeChan         = make(chan string, 1000)
	locationInfoChan = make(chan LocationInfo, 100)
)

func main() {

	//con, err := redis.Dial("tcp", "127.0.0.1:6379")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//defer con.Close()
	//log.Println("connect redis success")
	//
	//cpuNum := runtime.NumCPU() * 5
	//runtime.GOMAXPROCS(cpuNum)
	//
	//go func() {
	//	log.Println("begin generate code")
	//	for i := 0; i < 100000001; i++ {
	//		newCode := generate(i)
	//		codeChan <- newCode
	//	}
	//}()
	//
	//getRedirect := func(i int) {
	//	log.Println("goruntine ", i, " start")
	//	for {
	//		select {
	//		case code := <-codeChan:
	//			url := "http://dwz.cn/" + code
	//			location, err := getRedirectURL(url)
	//			if err != nil {
	//				// fmt.Println(url)
	//				// fmt.Println(err)
	//			}
	//
	//			locationInfoChan <- LocationInfo{Location: location, Key: code}
	//			break
	//		default:
	//			break
	//		}
	//	}
	//}
	//
	//for i := 0; i < cpuNum; i++ {
	//	go getRedirect(i)
	//}
	//
	//log.Println("start goruntine save to redis")
	//time.Sleep(time.Second)
	//for {
	//	select {
	//	case _ = <-locationInfoChan:
	//		//con.Do("SET", locationInfo.Key, locationInfo.Location)
	//		break
	//	default:
	//		break
	//	}
	//}
}
