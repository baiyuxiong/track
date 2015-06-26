package tests

import (
	"net/url"
)


func (t *AppTest) StartTestUser() {
	data := make(url.Values)
	data["name"] = []string{profileName}
	data["phone"] = []string{"12345678"}
	data["avatar"] = []string{profileAvatar}
	data["avatar_thumb1"] = []string{profileAvatar}
	data["avatar_thumb2"] = []string{profileAvatar}
	t.PostForm(t.GenUrl("/user/editProfile", token), data)
	t.AssertContains("200")

	t.Get(t.GenUrl("/user/me",token))
	t.AssertContains(profileName)
}