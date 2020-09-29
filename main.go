package main

import (
	"Project/lms"
	"fmt"
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
	var username,password string
	fmt.Printf("Enter Username : ");fmt.Scanln(&username)
	fmt.Printf("Enter Password : ");fmt.Scanln(&password)
	token := lms.LoginLms(username,password)
	courses := lms.FetchCourses(token)
	for name,link := range courses {
		fmt.Printf("Link : %v\nName : %v\n",link,name)
	}
}
