package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	_ "dao-service/resource-dao-service/routers"
	"github.com/astaxie/beego/orm"
	"model"
)

const (
	user_base_url = "http://localhost:8080/v1/user"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:corex123@tcp(localhost:3306)/PME?charset=utf8")
}

func Test_User_Create(t *testing.T) {
	var res *http.Response
	var err error
	var resBody []byte
	var requestData []byte

	// create user
	var user model.User
	user.Id = 0
	user.Name = "user" + strconv.FormatInt(time.Now().Unix(), 10)
	user.Email = "user@xx.com"
	user.Company = "company"
	user.Role = 1

	// post create resource
	requestData, err = json.Marshal(&user)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	res, err = http.Post(user_base_url+"/", "application/x-www-form-urlencoded", bytes.NewBuffer(requestData))
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()

	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}

func Test_User_GetAll(t *testing.T) {
	var res *http.Response
	var err error
	var resBody []byte

	// get all
	res, err = http.Get(user_base_url + "/")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}

func Test_User_GetById(t *testing.T) {
	var res *http.Response
	var err error
	var resBody []byte

	// get all
	res, err = http.Get(user_base_url + "/")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	var users []*model.User
	json.Unmarshal(([]byte)(response.Result), &users)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if len(users) <= 0 {
		t.Log("error : ", "there is no user to operate!")
		return
	}

	// get user by id
	res, err = http.Get(user_base_url + "/id/" + strconv.FormatInt(users[0].Id, 10))
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	response = model.Response{}
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}

func Test_User_GetAllExcludeOneId(t *testing.T) {
	var res *http.Response
	var err error
	var resBody []byte

	// get all
	res, err = http.Get(user_base_url + "/")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	var users []*model.User
	json.Unmarshal(([]byte)(response.Result), &users)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if len(users) <= 0 {
		t.Log("error : ", "there is no user to operate!")
		return
	}

	// get all users exclude one id
	res, err = http.Get(user_base_url + "/exclude/id/" + strconv.FormatInt(users[0].Id, 10))
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	response = model.Response{}
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}

func Test_User_GetByName(t *testing.T) {
	var res *http.Response
	var err error
	var resBody []byte

	// get all
	res, err = http.Get(user_base_url + "/")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	var users []*model.User
	json.Unmarshal(([]byte)(response.Result), &users)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if len(users) <= 0 {
		t.Log("error : ", "there is no user to operate!")
		return
	}

	// get user name
	res, err = http.Get(user_base_url + "/name/" + users[0].Name)
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	response = model.Response{}
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}

func Test_User_GetByRole(t *testing.T) {
	var res *http.Response
	var err error
	var resBody []byte

	// get user by role
	res, err = http.Get(user_base_url + "/role/1")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}

func Test_User_Update(t *testing.T) {
	var res *http.Response
	var err error
	var resBody []byte
	var requestData []byte

	// get all
	res, err = http.Get(user_base_url + "/")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	var users []*model.User
	json.Unmarshal(([]byte)(response.Result), &users)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if len(users) <= 0 {
		t.Log("error : ", "there is no user to operate!")
		return
	}

	users[0].Name = "user-update"
	users[0].Email = "user-update@xx.com"
	users[0].Company = "company-update"

	// put update job
	requestData, err = json.Marshal(&users[0])
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	// put
	client := http.Client{}
	req, _ := http.NewRequest("PUT", user_base_url, strings.NewReader(string(requestData)))

	res, err = client.Do(req)

	if err != nil {
		// handle error
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)

	if err != nil {
		t.Log("erro : ", err)
	}

	t.Log(string(resBody))

	response = model.Response{}
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}

func Test_User_DeleteById(t *testing.T) {
	var res *http.Response
	var err error
	var resBody []byte

	// get all
	res, err = http.Get(user_base_url + "/")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	var users []*model.User
	json.Unmarshal(([]byte)(response.Result), &users)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if len(users) <= 0 {
		t.Log("error : ", "there is no user to operate!")
		return
	}

	// delete
	client := http.Client{}
	req, _ := http.NewRequest("DELETE", user_base_url+"/id/"+strconv.FormatInt(users[0].Id, 10), nil)

	res, err = client.Do(req)

	if err != nil {
		// handle error
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)

	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	response = model.Response{}
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}
