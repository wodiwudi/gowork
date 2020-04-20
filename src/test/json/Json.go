package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age      int
	Birthday string
	Sex      string
	Email    string
	Phone    string
}

/*map转json*/
func testMap() ([]byte, error) {
	var mmp map[string]interface{}
	mmp = make(map[string]interface{})

	mmp["username"] = "Murphy"
	mmp["age"] = 19
	mmp["sex"] = "man"

	data, err := json.Marshal(mmp)
	if err != nil {
		fmt.Println("json marshal failed,err:", err)
		return make([]byte, 0), err
	}
	fmt.Printf("%s\n", string(data))
	return data, nil
}

/*json转map*/
func jsonToMap() map[string]interface{} {
	var receive map[string]interface{}
	jsonData, err := testMap()
	if err != nil {
		panic(err)
	}
	err2 := json.Unmarshal(jsonData, &receive)
	if err2 != nil {
		panic(err)
	}
	return receive
}

/*结构体转json*/
func testStruct() ([]byte, error) {
	user1 := &User{
		UserName: "user1",
		NickName: "Murphy",
		Age:      18,
		Birthday: "2008/8/8",
		Sex:      "男",
		Email:    "123456@qq.com",
		Phone:    "15600000000",
	}

	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return make([]byte, 0), err
	}

	fmt.Printf("%s\n", data)
	return data, nil

}

func jsonToStruct() User {
	data, err := testStruct()
	if err != nil {
		panic(err)
	}
	var user2 User
	err2 := json.Unmarshal(data, &user2)
	if err2 != nil {
		panic(err)

	}
	return user2
}

/*int转json*/

func testInt() {
	var age = 100
	data, err := json.Marshal(age)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
}

/*slice转json*/
func testSlice() {
	var m map[string]interface{}
	var s []map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "user1"
	m["age"] = 18
	m["sex"] = "man"

	s = append(s, m)

	m = make(map[string]interface{})
	m["username"] = "user2"
	m["age"] = 29
	m["sex"] = "female"
	s = append(s, m)

	data, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
}
func main() {
	//var user2 = jsonToStruct()
	//fmt.Printf("%T,%v\n", user2, user2)
	var user3 = jsonToMap()
	fmt.Printf("%T,%v\n", user3, user3)
	//testSlice()
	//testInt()
	//testMap()
	//testStruct()
	fmt.Println("----")
}
