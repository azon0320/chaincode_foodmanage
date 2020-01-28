package models

/*
 * TODO WASTED in this project
 */
type Permissions struct {
	PermissionTable []string
}

func (p *Permissions) UpdatePermissions(table []string) {
	p.PermissionTable = table
}

func (p *Permissions) hasPermission(sub string) bool {
	for _, value := range p.PermissionTable {
		if value == sub {
			return true
		}
	}
	return false
}

func NewPermissions(operatorType byte) *Permissions {
	var defaultPermissions map[byte][]string = make(map[byte][]string)
	//TODO Add preset permissions
	defaultPermissions[OperatorSeller] = []string{}
	defaultPermissions[OperatorBuyer] = []string{}
	defaultPermissions[OperatorTransporter] = []string{}
	var obj *Permissions = &Permissions{}
	value, ok := defaultPermissions[operatorType]
	if ok {
		obj.UpdatePermissions(value)
	} else {
		obj.UpdatePermissions([]string{})
	}
	return obj
}
