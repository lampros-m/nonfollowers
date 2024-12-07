package main

type ProfileData struct {
	Name string `json:"value"`
	URL  string `json:"href"`
}

type Profile struct {
	StringListData []ProfileData `json:"string_list_data"`
}

type FollowingsRoot struct {
	Followings Followings `json:"relationships_following"`
}

type RequestRoot struct {
	Requests Requests `json:"relationships_follow_requests_sent"`
}

type Followers []Profile

type Followings []Profile

type Requests []Profile

type ProfileURLMap map[string]string
