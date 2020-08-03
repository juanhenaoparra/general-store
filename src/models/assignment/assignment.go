package assignment

// AssignmentsRepo to save all asignments returned of the saving process
type AssignmentsRepo struct {
	Date         map[string]string
	Buyers       map[string]string
	Products     map[string]string
	Ips          map[string]string
	Devices      map[string]string
	Transactions map[string]string
}
