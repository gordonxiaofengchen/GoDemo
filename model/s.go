package model

type Student struct{
	Model
	Name string `json:"name"`
	Age int `json:"age"`
	BirthDay Time `json:"birthDay"`
}

func (Student) TableName() string{
	return "t_student"
}

type Staff struct{
	Model
	StaffNumber int `json:"staffNumber"`
	Name string `json:"name"`
	Password string `json:"password"`
	JoinDate Time `json:"joinDate"`
	Remarks string `json:"remarks"`
	Status string `json:"status"`
	Weight float32 `json:"weight"`
}

func (Staff) TableName() string{
	return "t_staff"
}