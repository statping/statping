package source

import "github.com/GeertJohan/go.rice"

var (
	SqlBox  *rice.Box
	CssBox  *rice.Box
	ScssBox *rice.Box
	JsBox   *rice.Box
	TmplBox *rice.Box
)

func Assets() {
	SqlBox = rice.MustFindBox("sql")
	CssBox = rice.MustFindBox("css")
	ScssBox = rice.MustFindBox("scss")
	JsBox = rice.MustFindBox("js")
	TmplBox = rice.MustFindBox("tmpl")
}
