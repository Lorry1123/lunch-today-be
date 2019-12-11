package restaurants

import (
	"io/ioutil"
	"encoding/json"
	"math/rand"
	"time"
)

// Restaurant 餐馆结构
type Restaurant struct {
	Name string	`json:"name"`
	Good string	`json:"good"`
	Bad string	`json:"bad"`
}

// Result 随机获取的餐馆结果
type Result struct {
	Goods []Restaurant	`json:"goods"`
	Bads []Restaurant	`json:"bads"`
}

type configs []Restaurant

func getDifferntRandomNumbers(r *rand.Rand, size int) []int {
	ret := map[int]int{}
	for len(ret) < 4 {
		n := r.Intn(size)
		if ret[n] == 1 {
			continue
		}
		ret[n] = len(ret) + 1
	}

	retList := make([]int, 4)
	for k, v := range ret {
		retList[v - 1] = k
	}

	return retList
}

// GetRandomRestuarants 获取随机餐馆
func GetRandomRestuarants() Result {
	data := readConfig()
	// read configs

	r := rand.New(rand.NewSource(time.Now().Unix()))
	// init random by ts -> returns defferent data for every query

	indexes := getDifferntRandomNumbers(r, len(data))
	return Result{[]Restaurant{data[indexes[0]], data[indexes[1]]}, []Restaurant{data[indexes[2]], data[indexes[3]]}}
}

// GetRestuarantsByDate 获取每日餐馆
func GetRestuarantsByDate() Result {
	data := readConfig()
	// read configs

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	r := rand.New(rand.NewSource(today.Unix()))
	// init random by date -> returns the same data in one day

	indexes := getDifferntRandomNumbers(r, len(data))
	return Result{[]Restaurant{data[indexes[0]], data[indexes[1]]}, []Restaurant{data[indexes[2]], data[indexes[3]]}}
}

func readConfig() []Restaurant {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return []Restaurant{}
	}

	res := &configs{}
	json.Unmarshal([]byte(file), &res)
	ret := []Restaurant{}
	for _, v := range *res {
		ret = append(ret, v)
	}
	return ret
}