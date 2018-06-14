package plugin

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	db         *sql.DB
	AllPlugins []Info
)

func CreateSettingsTable(p Info) string {
	var tableValues []string
	for _, v := range p.Form {
		tb := fmt.Sprintf("%v %v", v.SQLValue, v.SQLType)
		tableValues = append(tableValues, tb)
	}
	vals := strings.Join(tableValues, ", ")
	out := fmt.Sprintf("CREATE TABLE settings_%v (%v);", p.Name, vals)
	smtp, _ := db.Prepare(out)
	_, _ = smtp.Exec()
	InitalSettings(p)
	return out
}

func InitalSettings(p Info) {
	var tableValues []string
	var tableInput []string
	for _, v := range p.Form {
		val := fmt.Sprintf("'%v'", "example data")
		tableValues = append(tableValues, v.SQLValue)
		tableInput = append(tableInput, val)
	}
	vals := strings.Join(tableValues, ",")
	ins := strings.Join(tableInput, ",")
	sql := fmt.Sprintf("INSERT INTO settings_%v(%v) VALUES(%v);", p.Name, vals, ins)
	smtp, _ := db.Prepare(sql)
	_, _ = smtp.Exec()

	SelectSettings(p)
}

func (f FormElement) Val() string {
	var v string
	fmt.Println(f.Value)
	b, ok := f.Value.([]byte)
	if ok {
		v = string(b)
	}
	return v
}

func SelectSettings(p Info) []*FormElement {

	var newForm []*FormElement

	sql := fmt.Sprintf("SELECT * FROM settings_%v LIMIT 1", p.Name)
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	count := len(p.Form)
	valuePtrs := make([]interface{}, count)
	values := make([]interface{}, count)

	for rows.Next() {

		for i, _ := range p.Form {
			valuePtrs[i] = &values[i]
		}

		err = rows.Scan(valuePtrs...)
		if err != nil {
			panic(err)
		}

		for i, col := range p.Form {

			var v interface{}

			val := values[i]

			b, ok := val.([]byte)

			if ok {
				v = string(b)
			} else {
				v = val
			}

			ll := &FormElement{
				Name:        col.Name,
				Description: col.Description,
				SQLValue:    col.SQLValue,
				SQLType:     col.SQLType,
				Value:       v,
			}

			newForm = append(newForm, ll)

			fmt.Println(col.SQLValue, v)

			col.Value = v
		}

	}
	return newForm
}

func RunSQL(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args)
	return rows, err
}

func (i Info) Template() *template.Template {
	t := template.New("form")
	temp, _ := t.Parse("hello nworld")
	return temp
}

func DownloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func SetDatabase(database *sql.DB) {
	db = database
}

func (p *PluginInfo) InstallPlugin(w http.ResponseWriter, r *http.Request) {

	//sql := "CREATE TABLE " + p.Name + " (enabled BOOLEAN, api_key text, api_secret text, channel text);"
	//db.QueryRow(p.InstallSQL()).Scan()

	http.Redirect(w, r, "/plugins", http.StatusSeeOther)
}

func (p *PluginInfo) UninstallPlugin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/plugins", http.StatusSeeOther)
}

func (p *PluginInfo) SavePlugin(w http.ResponseWriter, r *http.Request) {
	//values := r.PostForm
	//p.SaveFunc(values)
	http.Redirect(w, r, "/plugins", http.StatusSeeOther)
}
