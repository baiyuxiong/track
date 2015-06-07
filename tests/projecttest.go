package tests

import (
)

func (t *AppTest) StartTestProject() {
	token = ""
	t.GetToken()
	t.ClearProjectTable()
	t.ClearCompanyUsersTable()
}

