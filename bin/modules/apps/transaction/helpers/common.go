package helpers

var Status = struct {
	Pending    string
	Rejected   string
	Canceled   string
	Transfered string
	Checked    string
	Finalized  string
	Used       string
}{
	"pending",
	"rejected",
	"canceled",
	"transfered",
	"checked",
	"finalized",
	"used",
}

var StatusMap = map[string]int{
	"pending": 0,
	"transfered": 1,
	"checked": 2,
	"finalized": 3,
	"used": 4,
	"rejected": 99,
	"canceled": 99,
}
