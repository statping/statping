package plugins

import "net/http"

func init() {
	AddRoute("save_slack", "POST", SaveIt)
	AddRoute("create_slack", "GET", Create)
}

func SaveIt(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/plugins", http.StatusSeeOther)
}

func Create(w http.ResponseWriter, r *http.Request) {
	CreateTable()
}

func CreateTable() {
	sql := "CREATE TABLE slack (enabled BOOLEAN, api_key text, api_key text, channel text);"
	db.QueryRow(sql).Scan()
}

func DropTable() {
	sql := "DROP TABLE slack;"
	db.QueryRow(sql).Scan()
}

func UpdateDatabase() {

}
