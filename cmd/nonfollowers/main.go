package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	FollowersFilePath = "connections/followers_and_following/followers_1.json"
	FollowingFilePath = "connections/followers_and_following/following.json"
)

func main() {
	var err error

	// Followers file
	data, err := os.ReadFile(FollowersFilePath)
	if err != nil {
		panic(err)
	}

	// fmt.Println("followers data len:", len(data))

	// Unmarshal followers file data
	var followers Followers

	err = json.Unmarshal(data, &followers)
	if err != nil {
		panic(err)
	}

	// fmt.Println("followers len:", len(followers))

	// Followers to map
	followerURLMap := make(ProfileURLMap)

	for _, follower := range followers {
		if len(follower.StringListData) != 1 {
			panic("unexpected follower data")
		}

		stringListData := follower.StringListData[0]
		followerURLMap[stringListData.Name] = stringListData.URL
	}

	// fmt.Println("followerURLMap len:", len(followerURLMap))

	// Following file
	data, err = os.ReadFile(FollowingFilePath)
	if err != nil {
		panic(err)
	}

	// fmt.Println("following data len:", len(data))

	// Unmarshal followers file data
	var followingsRoot FollowingsRoot

	err = json.Unmarshal(data, &followingsRoot)
	if err != nil {
		panic(err)
	}

	// fmt.Println("followings len:", len(followingsRoot.Followings))

	// Followings to map
	followingURLMap := make(ProfileURLMap)

	for _, following := range followingsRoot.Followings {
		if len(following.StringListData) != 1 {
			panic("unexpected following data")
		}

		stringListData := following.StringListData[0]
		followingURLMap[stringListData.Name] = stringListData.URL
	}

	// fmt.Println("followingURLMap len:", len(followingURLMap))

	// Compare
	nonFollowBack := make(ProfileURLMap)
	for name, url := range followingURLMap {
		if _, ok := followerURLMap[name]; !ok {
			nonFollowBack[name] = url
		}
	}
	var b []byte
	b, _ = json.MarshalIndent(nonFollowBack, "", "    ")

	fmt.Println("People don't follow back")
	fmt.Println("=========================")
	fmt.Println(string(b))
}
