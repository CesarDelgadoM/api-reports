package branch

type Branch struct {
	Name      string    `json:"name"`
	Manager   string    `json:"manager"`
	City      string    `json:"city"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Score     uint8     `json:"score"`
	Employees Employees `json:"employees"`
	Financial Financial `json:"financial"`
	Menu      Menu      `json:"menu"`
}

type Employees struct {
	Admins       []Employee `json:"admins"`
	Waiters      []Employee `json:"waiters"`
	Chefs        []Employee `json:"chefs"`
	TotalAdmins  uint8      `json:"total_admins"`
	TotalWaiters uint8      `json:"tatal_waiters"`
	TotalChefs   uint8      `json:"total_chefs"`
}

type Employee struct {
	Name  string `json:"name"`
	Years uint8  `years:"name"`
	Sales uint8  `json:"sales"`
}

type Financial struct {
	Sales    uint32 `json:"sales"`
	Expenses uint32 `json:"expenses"`
}

type Menu struct {
	EntreePlates []string `json:"entree_plates"`
	MainCourse   []string `json:"main_course"`
	Drinks       []string `json:"drinks"`
	Desserts     []string `json:"desserts"`
}

type Request struct {
	UserId  uint   `json:"userid"`
	Name    string `json:"name"`
	Manager string `json:"manager"`
	Branch  Branch
}
