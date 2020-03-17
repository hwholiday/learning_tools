package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	GlobalClient := redis.NewClient(
		&redis.Options{
			Addr: "127.0.0.1:6379",
		},
	)
	err := GlobalClient.Ping().Err()
	if nil != err {
		panic(err)
	}
	fmt.Println("链接redis成功")
	res, err := GlobalClient.GeoAdd("geo_hash_test", &redis.GeoLocation{
		Name:      "天府广场",
		Longitude: 104.072833,
		Latitude:  30.663422,
	}, &redis.GeoLocation{
		Name:      "四川大剧院",
		Longitude: 104.074378,
		Latitude:  30.664804,
	}, &redis.GeoLocation{
		Name:      "新华文轩",
		Longitude: 104.070084,
		Latitude:  30.664649,
	}, &redis.GeoLocation{
		Name:      "手工茶",
		Longitude: 104.072402,
		Latitude:  30.664121,
	}, &redis.GeoLocation{
		Name:      "宽窄巷子",
		Longitude: 104.059826,
		Latitude:  30.669883,
	}, &redis.GeoLocation{
		Name:      "奶茶",
		Longitude: 104.06085,
		Latitude:  30.670054,
	}, &redis.GeoLocation{
		Name:      "钓鱼台",
		Longitude: 104.058424,
		Latitude:  30.670737,
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("[GeoAdd]", res)
	resPos, err := GlobalClient.GeoPos("geo_hash_test", "天府广场", "宽窄巷子").Result()
	if err != nil {
		panic(err)
	}
	for _, v := range resPos {
		fmt.Println("[GeoPos]", "Longitude : ", v.Longitude, "Latitude : ", v.Latitude)
	}
	//其中 unit 参数是距离单位
	//m 表示单位为米。
	//km 表示单位为千米。
	//mi 表示单位为英里。
	//ft 表示单位为英尺。
	resDist, err := GlobalClient.GeoDist("geo_hash_test", "天府广场", "宽窄巷子", "m").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("[GeoDist]", resDist)

	resRadiu, err := GlobalClient.GeoRadius("geo_hash_test", 104.072833, 30.663422, &redis.GeoRadiusQuery{
		Radius:      800,   //radius表示范围距离，
		Unit:        "m",   //距离单位是 m|km|ft|mi
		WithCoord:   true, //传入WITHCOORD参数，则返回结果会带上匹配位置的经纬度
		WithDist:    true, //传入WITHDIST参数，则返回结果会带上匹配位置与给定地理位置的距离。
		WithGeoHash: true, //传入WITHHASH参数，则返回结果会带上匹配位置的hash值。
		Count:       4,     //入COUNT参数，可以返回指定数量的结果。
		Sort:        "ASC", //默认结果是未排序的，传入ASC为从近到远排序，传入DESC为从远到近排序。
	}).Result()
	if err != nil {
		panic(err)
	}
	for _,v:=range resRadiu{
		fmt.Println("[GeoRadiu]", v)
	}

	resRadiusByMember, err := GlobalClient.GeoRadiusByMember("geo_hash_test", "天府广场", &redis.GeoRadiusQuery{
		Radius:      800,
		Unit:        "m",
		WithCoord:   true,
		WithDist:    true,
		WithGeoHash: true,
		Count:       4,
		Sort:        "ASC",
	}).Result()
	if err != nil {
		panic(err)
	}

	for _,v:=range resRadiusByMember{
		fmt.Println("[GeoRadiusByMember]", v)
	}
	resHash, err := GlobalClient.GeoHash("geo_hash_test", "天府广场").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("[GeoHash]", resHash)
}
