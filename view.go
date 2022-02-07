package models

type View struct {
	Todos []Todo
}


slicee:=[]string{"Netflix","Metaversr","Amzone","another life"}

for index,err:=range slicee{
	
	if err !=nil{
		panic("something went wrong")
	}
	fmt.Println("The data of the slice is as follow")
	slicee=append(slicee[:index+1],slicee[index+1:]...)
}

fmt.Println(slicee)
}
