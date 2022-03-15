package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

type UserInfo struct {
	Card struct {
		Attentions []int `json:"attentions"`
	} `json:"card"`
	Face     string `json:"face"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}
type UserComponentInfo struct {
}

func LoadFile(filename string) ([]byte, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func GetAttentionList(uid string) ([]int, error) {
	var userInfo UserInfo
	resp, err := http.Get("https://account.bilibili.com/api/member/getCardByMid?mid=" + uid)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println(err)
		return nil, err
	}
	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = json.Unmarshal(bodyData, &userInfo)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return userInfo.Card.Attentions, nil
}

type Vtb struct {
}

func GetConcernedVtbs() ([]Vtb, error) {
	var Vtbs []Vtb
	var uidVtbMap map[string]Vtb
	data, err := LoadFile("vtbs.json")
	if err != nil {
		fmt.Println(err)
		return Vtbs, err
	}
	json.Unmarshal(data, &uidVtbMap)
	attentionList, err := GetAttentionList("13539525")
	if err != nil {
		fmt.Println(err)
		return Vtbs, err
	}
	for _, attentionUID := range attentionList {
		if val, ok := uidVtbMap[strconv.Itoa(attentionUID)]; ok {
			Vtbs = append(Vtbs, val)
		}
	}
	return Vtbs, nil
}

func main() {
	vtbs, err := GetConcernedVtbs()
	if err != nil {
		return
	}
	fmt.Println(vtbs)
}

func convertInterfaceToIntSlice(i interface{}) ([]int, error) {
	var IntSlice []int
	k := reflect.TypeOf(i).Kind()
	fmt.Println(k)
	value := reflect.ValueOf(i)
	fmt.Println(k)
	if k != reflect.Slice {
		err := errors.New("kind error interface")
		return IntSlice, err
	}
	for i := 0; i < value.Len(); i++ {
		element := value.Index(i)
		if element.Kind() != reflect.Int {
			err := errors.New("kind error value")
			return IntSlice, err
		}
		IntSlice = append(IntSlice, int(element.Int()))
	}
	return IntSlice, nil
}
