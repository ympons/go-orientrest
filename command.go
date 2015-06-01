package orientrest

import (
	"fmt"
	"log"
)

func (d *ODatabase) CmdInterrupt(cmd string) error {
	pUrl := fmt.Sprintf("%sdbconnection/%s", d.URL, d.Name)
	if resp, err := d.Session.Post(pUrl, cmd, nil, nil); err != nil {
		return err
	} else {
		log.Printf("[CmdInterrupt]. CODE: %d URI: %s", resp.Status(), pUrl)
	}
	return nil
}

// TODO Validate lang against AvailableLangs
func (d *ODatabase) Command(cmd string, params ...interface{}) error {
	defer func() {
		if s := recover(); s != nil {
			log.Printf("[Command]. ERROR: %+v", s)
		}
	}()

	lang, limit := "sql", 20
	if params != nil {
		if l := len(params); l > 0 && l < 3 {
			if l == 2 {
				limit = params[1].(int)
			}
			lang = params[0].(string)
		} else {
			return fmt.Errorf("[Command]. ERROR: Many parameters in Command function: %d", l)
		}
	}
	var r interface{}
	pUrl := fmt.Sprintf("%scommand/%s/%s/%s/%d?format=rid,type,version,class,graph", d.URL, d.Name, lang, cmd, limit)
	if resp, err := d.Session.Post(pUrl, nil, &r, nil); err != nil {
		return err
	} else {
		log.Printf("[Command]. CODE: %d URI: %s", resp.Status(), pUrl)
	}
	return nil
}

func (d *ODatabase) CmdGetAll(clazz string) (interface{}, error) {
	return nil, nil
}
