package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hey)
	http.ListenAndServe(":8081", nil)
}

const tpl = `
<html>
	<head>
		<title>Hey</title>
	</head>
	<body>
		<form action="/" method="post" >
			Username: <input type="text" name="uname"><br/>
			Password: <input type="password" name="pwd"><br/>
			<button type="submit" value="提交">提交</button>
		</form>
	</body>
</html>
`

func Hey(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t := template.New("hey")
		t.Parse(tpl)
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		//解析并处理表单
		/*
			//第一种
			r.ParseForm()
			fmt.Println("用户名：", r.Form["uname"][0], "密码：", r.Form["pwd"])
		*/

		//第二种
		fmt.Println(r.FormValue("uname"), ":", r.FormValue("pwd"))
	}
}
