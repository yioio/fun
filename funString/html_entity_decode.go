package funString

import "html"

func Html_entity_decode(str string)  string{

	return html.UnescapeString(str)
	
}