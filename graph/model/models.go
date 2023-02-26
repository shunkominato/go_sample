package model

type Employee struct {
    ID            string `json:"id"`
    Name          string `json:"name"`
    Gender        Gender `json:"gender"`
    Email         string `json:"email"`
    LatestLoginAt string `json:"latestLoginAt"`
    //  扶養家族の人数
    DependentsNum int `json:"dependentsNum"`
    //  管理職かどうか
    IsManager    bool   `json:"isManager"`
    DepartmentID string `json:"department"` // Departmentを削除して、代わりにDepartmentIDを記述
    CompanyID    string `json:"company"`    // Companyを削除して、代わりにCompanyIDを記述
}
func (Employee) IsNode() {}