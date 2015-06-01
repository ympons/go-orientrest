package orientrest

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/jmcvetta/napping"
)

const (
	DB_TYPE_GRAPH       = "graph"
	DB_TYPE_DOCUMENT    = "document"
	STORAGE_TYPE_LOCAL  = "local"
	STORAGE_TYPE_PLOCAL = "plocal"
	STORAGE_TYPE_MEMORY = "memory"
)

// A ODatabase is a REST client connected to a OrientDB Server
type ODatabase struct {
	Session *napping.Session
	URL     string
	Name    string
	Classes []OClass `json:"classes"`
}


// For specifying parameters to initialize the client
type Options struct {
	DbName string
	DbUser string
	DbPass string
	Conn   bool
}

// Initialize the client for the OrientDB Server
func OrientDB(uri string, params ...Options) (*ODatabase, error) {
	h := http.Header{}
	h.Add("User-Agent", "gorientrest")
	h.Add("Accept-Encoding", "gzip,deflate")
	db := &ODatabase{
		Session: &napping.Session{Header: &h},
	}

	if uri == "" {
		uri = "http://127.0.0.1:2480"
	}

	if !strings.HasSuffix(uri, "/") {
		uri += "/"
	}

	pURL, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	db.URL = pURL.String()

	if db.configure(params) {
		return db.Connect()
	}

	return db, nil
}

// Internal method to setup parameters 
func (d *ODatabase) configure(params []Options) bool {
	if params != nil && len(params) > 0 {
		d.Close()
		p := params[0]
		d.Name = p.DbName
		d.Session.Userinfo = url.UserPassword(p.DbUser, p.DbPass)
		return p.Conn
	}
	return false
}

// 
func (d *ODatabase) Configure(params ...Options) {
	d.configure(params)
}

func (d *ODatabase) Connect(params ...Options) (*ODatabase, error) {
	d.configure(params)

	pUrl := fmt.Sprintf("%sconnect/%s", d.URL, d.Name)
	var r string
	if resp, err := d.Session.Get(pUrl, nil, &r, nil); err != nil {
		return nil, fmt.Errorf("[Connect]. ERROR: %v", err)
	} else if r := resp.Status(); r == 401 {
		return nil, fmt.Errorf("[Connect]. ERROR: %s", resp.RawText())
	} else if r != 204 {
		return nil, fmt.Errorf("[Connect]. ERROR: Status %d trying to connect to %s", r, pUrl)
	}
	return d, nil
}

func (d *ODatabase) Close() {
	pUrl := fmt.Sprintf("%sdisconnect", d.URL)
	if _, err := d.Session.Get(pUrl, nil, nil, nil); err != nil {
		log.Printf("[Close]. ERROR: %v", err)
	}
	d.Name = "" // Maybe "d = nil" or d = &ODatabase{} is better
}

func (d *ODatabase) DbCreate(dbname, dbstoretype string, dbtype ...string) error {
	contain := func(v string, l []string) bool {
		for _, s := range l {
			if s == v {
				return true
			}
		}
		return false
	}
	if !contain(dbstoretype, []string{STORAGE_TYPE_LOCAL, STORAGE_TYPE_MEMORY, STORAGE_TYPE_PLOCAL}) {
		return fmt.Errorf("[DbCreate]. ERROR: Invalid dbstoretype: %s", dbstoretype)
	}

	var dtype = DB_TYPE_DOCUMENT
	if dbtype != nil && len(dbtype) > 0 {
		dtype = dbtype[0]
	}
	if !contain(dtype, []string{DB_TYPE_GRAPH, DB_TYPE_DOCUMENT}) {
		return fmt.Errorf("[DbCreate]. ERROR: Invalid dbtype: %s", dtype)
	}

	pUrl := fmt.Sprintf("%sdatabase/%s/%s/%s", d.URL, dbname, dbstoretype, dtype)

	if resp, err := d.Session.Post(pUrl, nil, &d, nil); err != nil {
		return err
	} else {
		log.Printf("[DbCreate]. CODE: %d URI: %s", resp.Status(), pUrl)
	}
	return nil
}

func (d *ODatabase) DbDrop(dbname string) error {
	pUrl := fmt.Sprintf("%sdatabase/%s", d.URL, dbname)
	if resp, err := d.Session.Delete(pUrl, nil, nil); err != nil {
		return err
	} else {
		log.Printf("DbDrop: %d", resp.Status())
	}
	return nil
}

func (d *ODatabase) DbInfo(dbname string) ([]OClass, error) {
	pUrl := fmt.Sprintf("%sdatabase/%s", d.URL, dbname)
	if resp, err := d.Session.Get(pUrl, nil, &d, nil); err != nil {
		return nil, err
	} else {
		log.Printf("DbInfo: %d", resp.Status())
	}
	return d.Classes, nil
}

func (d *ODatabase) DbExport(dbname string) ([]byte, error) {
	return nil, nil //TODO
}

func (d *ODatabase) DbImport(dbname string, file []byte) error {
	return nil //TODO
}

func (d *ODatabase) DbList() (*ODbList, error) {
	pUrl := fmt.Sprintf("%slistDatabases", d.URL)
	var r *ODbList
	if resp, err := d.Session.Get(pUrl, nil, &r, nil); err != nil {
		return nil, err
	} else {
		log.Printf("DbList: %d", resp.Status())
	}
	return r, nil
}

func (d *ODatabase) DbAvailableLangs(dbname string) (*ODbLang, error) {
	pUrl := fmt.Sprintf("%ssupportedLanguages/%s", d.URL, dbname)
	var r *ODbLang
	if resp, err := d.Session.Get(pUrl, nil, &r, nil); err != nil {
		return nil, err
	} else {
		log.Printf("DbAvailableLangs: %d", resp.Status())
	}
	return r, nil
}
