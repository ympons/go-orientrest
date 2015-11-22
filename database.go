package orientrest

import (
	"bytes"
	"fmt"
	"mime/multipart"
)

type DatabaseType string

const (
	DB_TYPE_GRAPH       DatabaseType = "graph"
	DB_TYPE_DOCUMENT    DatabaseType = "document"
)

type StoreType string

const (
	STORAGE_TYPE_PLOCAL StoreType = "plocal"
	STORAGE_TYPE_MEMORY StoreType = "memory"
)

type Admin struct {
	client *Client
}

func (a *Admin) DbCreate(name string, dbType DatabaseType, storeType StoreType) (*ODatabase, error) {
	u := fmt.Sprintf("database/%s/%s/%s", name, storeType, dbType)

	req, err := a.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}

	var r *ODatabase
	err = a.client.Do(req, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (a *Admin) DbDrop(name string) error {
	u := fmt.Sprintf("database/%s", name)

	req, err := a.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	return a.client.Do(req, nil)
}

func (a *Admin) DbInfo(name string) (*ODatabase, error) {
	u := fmt.Sprintf("database/%s", name)

	req, err := a.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	var r *ODatabase
	err = a.client.Do(req, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *Admin) DbExport(name string) (interface{}, error) {
	u := fmt.Sprintf("export/%s", name)

	req, err := a.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var r interface{}
	err = a.client.Do(req, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// There is an issue with importing DB using Orient REST Api: https://github.com/orientechnologies/orientdb/issues/3431
func (a *Admin) DbImport(name string, file []byte) (interface{}, error) {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	part, err := w.CreateFormFile("filename", "dbimport")
	if err != nil {
		return nil, err
	}
	part.Write(file)
	w.Close()


	u := fmt.Sprintf("import/%s", name)
	req, err := a.client.NewUploadRequest(u, buf, w.FormDataContentType(), int64(buf.Len()))
	if err != nil {
		return nil, err
	}

	var r interface{}
	err = a.client.Do(req, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *Admin) DbList() (*ODbList, error) {
	req, err := a.client.NewRequest("GET", "listDatabases", nil)
	if err != nil {
		return nil, err
	}

	var r *ODbList
	err = a.client.Do(req, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *Admin) DbAvailableLangs(name string) (*ODbLang, error) {
	u := fmt.Sprintf("supportedLanguages/%s", name)

	req, err := a.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var r *ODbLang
	err = a.client.Do(req, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *Admin) Close() error {
	return nil
}

type Database struct {
	name string
	client *Client
}

func (d *Database) connect() (*Database, error) {
	u := fmt.Sprintf("connect/%s", d.name)

	req, err := d.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	err = d.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Database) Close() error {
	req, err := d.client.NewRequest("GET", "disconnect", nil)
	if err != nil {
		return err
	}

	err = d.client.Do(req, nil)
	if err != nil {
		return err
	}

	d.name = "" // Maybe "d = nil" or d = &ODatabase{} is better

	return nil
}
