package main

import (
	"Project/lms"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	asciiArt :=
		`
    ___       ___       ___       ___       ___       ___       ___   
   /\  \     /\  \     /\  \     /\  \     /\__\     /\__\     /\  \  
  /::\  \   /::\  \   _\:\  \   _\:\  \   /:/  /    /::L_L_   /::\  \ 
 /::\:\__\ /::\:\__\ /::::\__\ /\/::\__\ /:/__/    /:/L:\__\ /\:\:\__\
 \;:::/  / \/\::/  / \::;;/__/ \::/\/__/ \:\  \    \/_/:/  / \:\:\/__/
  |:\/__/    /:/  /   \:\__\    \:\__\    \:\__\     /:/  /   \::/  / 
   \|__|     \/__/     \/__/     \/__/     \/__/     \/__/     \/__/  

`
	fmt.Println(asciiArt)
	var courses = map[string]string{}
	var username,password,token string
	var selectarr int
	var courselist []string
	fmt.Println("# Looking for token.ini file...\n")
	_,exists := os.Stat("token.ini")
	if os.IsNotExist(exists){
		fmt.Println("# token.ini not found!")
		fmt.Printf("Enter Username : ");fmt.Scanln(&username)
		fmt.Printf("Enter Password : ");fmt.Scanln(&password)
		token = lms.LoginLms(username,password)
		os.Create("token.ini")
		ioutil.WriteFile("token.ini",[]byte(token),0644)
	}else {
		fmt.Println("# token.ini found!\n\n")
		readtoken,_ := ioutil.ReadFile("token.ini")
		token = string(readtoken)
	}
	courses = lms.FetchCourses(token)
	count := 0
	for name,link := range courses {
		count = count + 1
		//fmt.Printf("(%v) Link : %v\nName : %v\n\n",strconv.Itoa(count),link,name)
		link += "STRING_VAL"
		link = strings.Replace(link,"/Lesson/","/VirtualAdmin/",1)
		laststring := strings.Split(link,"/")
		removestr := "/" + laststring[len(laststring) - 1]
		link = strings.Replace(link,removestr,"",1)
		fmt.Printf("(%v) Link : %v\nName : %v\n\n",strconv.Itoa(count),link,name)
		courselist = append(courselist,link)
	}

	fmt.Printf("Select : ");fmt.Scanln(&selectarr)
	lms.CourseDetail(courselist[selectarr - 1],token)
}
