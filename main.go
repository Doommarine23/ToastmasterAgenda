package main

import (
	"fmt"
	"github.com/nguyenthenguyen/docx"
)

func main() {
	fmt.Println("Hello World")

	r, err := docx.ReadDocxFile("./Agenda.docx")
	if err != nil {
		panic(err)
	}

	roleDate := AgendaDayMonthYear()
	roles := GetRoles(roleDate)

	docx1 := r.Editable()
	date := "./" + AgendaDate() + ".docx"
	docx1.Replace("<<Date>>", date, -1)
	docx1.Replace("<<President>>", "Shelly Bowe", -1)
	docx1.Replace("<<VPE>>", "Ann Addicks", -1)
	docx1.Replace("<<Toastmaster>>", roles.toastmaster, -1)
	docx1.Replace("<<GE>>", roles.ge, -1)
	docx1.Replace("<<Timer>>", roles.timer, -1)
	docx1.Replace("<<Ah-Counter>>", roles.ahCounter, -1)
	docx1.Replace("<<Grammarian>>", roles.grammarian, -1)
	docx1.Replace("<<Evaluator1>>", roles.eval1, -1)
	docx1.Replace("<<Speaker1>>", roles.speaker1, -1)
	docx1.Replace("<<Evaluator2>>", roles.eval2, -1)
	docx1.Replace("<<Speaker2>>", roles.speaker2, -1)
	docx1.Replace("<<Evaluator3>>", roles.eval3, -1)
	docx1.Replace("<<Speaker3>>", roles.speaker3, -1)
	docx1.Replace("<<Evaluator4>>", roles.eval4, -1)
	docx1.Replace("<<Speaker4>>", roles.speaker4, -1)
	docx1.Replace("<<TableTopicsMaster>>", roles.tableTopicsMaster, -1)
	docx1.WriteToFile(date)

	r.Close()
}
