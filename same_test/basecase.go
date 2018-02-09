package same_test

type ICase interface {
	Test(datamap map[string]string, url string) interface{}
	Mockdata(cases []interface{}) map[string]string
	GetTestCases() []interface{}
}
