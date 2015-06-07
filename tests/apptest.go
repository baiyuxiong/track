package tests

import "github.com/revel/revel/testing"

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestAll() {
	t.StartTestAuth()
	t.StartTestCompany()
	t.StartTestProject()
}

func (t *AppTest) After() {
	println("Tear down")
}