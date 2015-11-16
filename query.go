package orientrest
/*
import (
	"fmt"
	"log"
	"net/url"
)

func (d *ODatabase) Query(query string, params ...interface{}) (*OQueryResult, error) {
	defer func() {
		if s := recover(); s != nil {
			log.Printf("Command: %+v", s)
		}
	}()

	lang, limit, fetch := "sql", 20, "*:0"
	if params != nil {
		if l := len(params); l > 0 && l < 3 {
			if l == 2 {
				fetch = params[1].(string)
			}
			limit = params[0].(int)
		} else {
			return nil, fmt.Errorf("Many parameters in Command function: %d", l)
		}
	}
	r := OQueryResult{}
	pUrl := fmt.Sprintf("%squery/%s/%s/%s/%d/%s", d.URL, d.Name, lang, url.QueryEscape(query), limit, fetch)
	if resp, err := d.Session.Get(pUrl, nil, &r, nil); err != nil {
		return nil, err
	} else {
		log.Printf("[Query]. Code: %d Uri: %s", resp.Status(), pUrl)
	}
	return &r, nil
}
*/
