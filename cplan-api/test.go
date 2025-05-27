package main

import (
	"fmt"

	"net/http"

	"io/ioutil"
)

func TestMain() {

	url := "https://dev-g4l3aoxdjakvoo55.us.auth0.com/userinfo"

	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {

		fmt.Println(err)

		return

	}

	req.Header.Add("Accept", "application/json")

	req.Header.Add("access_token", "eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIiwiaXNzIjoiaHR0cHM6Ly9kZXYtZzRsM2FveGRqYWt2b281NS51cy5hdXRoMC5jb20vIn0..euP4DilkFSvb7Fsj.KXyRhqckumoF92-ansrGw5lK2bM_k_BE6Fi5gmdrwTH18ALtCLqcNiuviUTzB5Xiv_QGgZJ7GU4xIMxmvchLHDRbTTiKIg1unj1Y6JDrzEzqExjG_x2GOP0AMEYgOBjZzztX_TTV9cJyeKiIhJ_IMlbM-x4nny1lQvMOY9NyEYbL-rwEVufnJsy41Zv7fi7fhjrJ6TJ29IAgZnAohlJ0iAaA6wcAaN_KQ5lA5jRiI6u1UUl8qAfaL8IeC9iesQ-sCQ1CkrGUGwXFcmJv-PDP-qpwggo285DWUSxgU6lybeD-q-ZImd35cmTMI1C6MXHPx3xIsqTBg_zqIly77BH5qD3ZoLHN63Cg9ONWiPe2c3Bo.GKmspH1UE6IEZcE8fRxumA")

	res, err := client.Do(req)

	if err != nil {

		fmt.Println(err)

		return

	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println(string(body))

}
