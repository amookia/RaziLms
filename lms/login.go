package lms

import (
	"github.com/gocolly/colly"
)


func LoginLms(username,password string) string{
	c := colly.NewCollector()
	var viewstate,ret string
	const url string = "https://lms2.razi.ac.ir/LoginPage.aspx?ReturnUrl=%2f"

	// Find and visit all links
	c.OnHTML("input", func(e *colly.HTMLElement) {
		if (e.Attr("name") == "__VIEWSTATE"){
			viewstate = e.Attr("value")
		}
	})


	_ = c.Visit(url)
	_ = c.Post(url, map[string]string{
		"__VIEWSTATE":                   viewstate,
		"ctl00$mainContent$hdfPass":     password,
		"ctl00$mainContent$UserName":    username,
		"ctl00$mainContent$Password":    password,
		"ctl00$mainContent$LoginButton": "ورود",
	})

	cookies := c.Cookies(url)
	for _,name := range cookies{

		//fmt.Println(reflect.TypeOf(name))
		if (name.Name == ".ASPXAUTH"){
			ret = name.Value
		}else{
			ret = "User name or password is wrong"
		}
	}
	return ret
}