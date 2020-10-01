package lms

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"strings"
	"time"
)

func FetchCourses(token string) map[string]string{
	m := make(map[string]string)
	url := "https://lms2.razi.ac.ir/ViewProfile.aspx"
	req,_ := http.NewRequest("GET",url,nil)
	cookie := http.Cookie{Name: ".ASPXAUTH", Value: token}
	req.AddCookie(&cookie)
	client := &http.Client{Timeout: time.Second * 10}
	resp,err := client.Do(req)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}else {
		HtmlResp := resp
		doc,_ := goquery.NewDocumentFromReader(HtmlResp.Body)
		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			// Get all courses
			rel,_ := s.Attr("rel")
			target,_ := s.Attr("target")
			if rel == "noopener noreferrer" && target == "_blank" {
				link, _ := s.Attr("href")
				name := s.Text()
				bi := strings.Contains(link, "/Lesson/")
				golestan := strings.Contains(name,"گروه")
				if bi && golestan {
					m[name] = "https://lms2.razi.ac.ir" + link
				}
			}
		})
	}
	return m
}


func CourseDetail(url string,token string){
	req,_ := http.NewRequest("GET",url,nil)
	cookie := http.Cookie{Name: ".ASPXAUTH", Value: token}
	req.AddCookie(&cookie)
	client := &http.Client{Timeout: time.Second * 10}
	resp,err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}else{
		HtmlResp := resp
		doc,_ := goquery.NewDocumentFromReader(HtmlResp.Body)
		doc.Find("div.grid-parent").Find("[class=\"cell cellb meetingdetial\"]").Find("span").Each(func(i int, s *goquery.Selection){
			abc := s.Text()
			abc = strings.Replace(abc,":","",1)
			fmt.Println(abc)
		})

	}

}