package orientrest

import (
	"fmt"
	"net/url"
)

type OCommonSQL interface {
	Limit(int) OCommonSQL
	Lang(string) OCommonSQL
	FetchPlan(string) OCommonSQL
}

type commonSQL struct {
	text string
	params []interface{}
	limit int
	lang string
	fetchPlan string
}

func (c *commonSQL) Limit(limit int) OCommonSQL {
	c.limit = limit
	return c
}

func (c *commonSQL) Lang(lang string) OCommonSQL {
	c.lang = lang
	return c
}

func (c *commonSQL) FetchPlan(fetch string) OCommonSQL {
	c.fetchPlan = fetch
	return c
}

type querySQL struct {
	*commonSQL
}

type commandSQL struct {
	*commonSQL
}

func NewCommandSQL(cmd string, params ...interface{}) commandSQL {
	return commandSQL{&commonSQL{
		text: cmd,
		params: params,
		limit: -1,
		lang: "sql",
		fetchPlan: "*:0",
	}}
}

func NewQuerySQL(query string, params ...interface{}) querySQL {
	return querySQL{&commonSQL{
		text: query,
		params: params,
		limit: -1,
		lang: "sql",
		fetchPlan: "*:0",
	}}
}

func (d *Database) command_(cmd commandSQL) (*OResult, error) {
	u := fmt.Sprintf("command/%s/%s/%s/%d/%s?format=rid,type,version,class,graph",
		d.name,
		cmd.commonSQL.lang,
		url.QueryEscape(cmd.commonSQL.text),
		cmd.commonSQL.limit,
		cmd.commonSQL.fetchPlan)

	var payload interface{}
	if params := cmd.commonSQL.params; params != nil {
		payload = map[string]interface{}{
			"command": cmd.commonSQL.text,
			"parameters": cmd.commonSQL.params,
		}
	}
	req, err := d.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, err
	}

	var r *OResult
	err = d.client.Do(req, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (d *Database) query_(query querySQL) (*OResult, error) {
	u := fmt.Sprintf("query/%s/%s/%s/%d/%s",
		d.name,
		query.commonSQL.lang,
		url.QueryEscape(query.commonSQL.text),
		query.commonSQL.limit,
		query.commonSQL.fetchPlan)

	req, err := d.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var r *OResult
	err = d.client.Do(req, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (d *Database) Command(cmd OCommonSQL) (*OResult, error) {
	switch v := cmd.(type) {
		case querySQL:
			return d.query_(v)
		case commandSQL:
			return d.command_(v)
	}
	return nil, fmt.Errorf("orientrest: Invalid sql command")
}

func (d *Database) CmdInterrupt(cmd string) error {
	u := fmt.Sprintf("dbconnection/%s", d.name)

	req, err := d.client.NewRequest("POST", u, url.QueryEscape(cmd))
	if err != nil {
		return err
	}

	err = d.client.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) CmdGetAll(clazz string) (interface{}, error) {
	return nil, nil
}
