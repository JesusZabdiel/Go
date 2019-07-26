package main

import(
	"fmt"
)


type User interface{ //in interfaces we only define methos (name and type)
	Permissions() int
	Name() string
	 
}
type Admin struct{
	name string

}

type Editor struct{
	name string
}

func (this Editor) Permissions() int{
	return 3
}

func (this Editor)  Name()string{
	return this.name
}

func (this Admin) Permissions() int{
	return 5
}

func (this Admin)  Name()string{
	return this.name
}

func auth(user User) string{
	permises := user.Permissions()

	if permises == 5{
		return user.Name() + " Tiene permisos de Adminsitrador"
	}else if permises < 5{
			return user.Name() + " Tiene permisos de editor"
		
	}
	return ""
}

func main(){


	usuarios := []User{Admin{"Zabdiel"}, Editor{"Jesus"}}

	for _, usuario := range usuarios{
		fmt.Println(auth(usuario))
	}
	

}