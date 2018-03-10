package main
import (
	"fmt"
	"io/ioutil"
	"encoding/xml"
	"os"
	"io"
)

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users []User `xml:"user"`
}

type User struct {
	XMLName xml.Name `xml:"user"`
	Type string `xml:"type,attr"`
	Name string `xml:"name"`
	Social Social `xml:"social"`
}

type Social struct {
	XMLName xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
}

func main() {
	xmlFile,err := os.Open("test.xml")
	if(err != nil) {
		fmt.Println(err)
	}

	fmt.Println("Success")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var users Users
	xml.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User type:"+users.Users[i].Type)
		fmt.Println("User name:"+users.Users[i].Name)
		fmt.Println("Facebook:"+users.Users[i].Social.Facebook)
	}

	newFileName := "newFile.xml"
	newFile,_:=os.Create(newFileName)
	xmlWriter := io.Writer(newFile)
	enc := xml.NewEncoder(xmlWriter)
	enc.Encode(users)

}