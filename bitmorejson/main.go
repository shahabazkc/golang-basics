package main

import (
	"encoding/json"
	"fmt"
)

type programmingLanguages struct {
	Name string `json:"languagename"`
	Platform string `json:"languageplatform"`
	Tags []string `json:"tags,omitempty"`
}

func main(){ 
	fmt.Println("Welcome to json");
	// EncodeJson()
	decodeJson()
}

func decodeJson(){
	jsonDataFromWeb := []byte(`
	{
		"languagename": "GoLang",
		"languageplatform": "go",
		"tags": ["web","backend"]
	}
	`)

	var lang programmingLanguages
	checkValid:=json.Valid(jsonDataFromWeb)

	if checkValid{
		fmt.Println("Valid Json")
		err := json.Unmarshal(jsonDataFromWeb,&lang)
		checkNilErr(err)
		fmt.Printf("%#v\n",lang)
	}else {
		fmt.Println("Invalid Json")
	}

	var data map[string]interface{}
	json.Unmarshal(jsonDataFromWeb,&data)
	fmt.Printf("%#v\n",data);

	for k,v := range data{
		fmt.Printf("Key: %v  Value: %v\n",k,v);
	}

}

func EncodeJson() {
	languages:= []programmingLanguages{
		{"Javascript","JS",[]string{"web","js"}},
		{"Python","py",[]string{"backend","machine learning"}},
		{"GoLang","go",[]string{"web","backend"}},
	}
	
	finalJson,err := json.MarshalIndent(languages,"","\t")
	checkNilErr(err);
	fmt.Printf("%s\n",finalJson)
}

func checkNilErr(err error){
	if err != nil{
		fmt.Println("Error: ", err)
		panic(err)
	}
}