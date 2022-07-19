package notion

import (
	"github.com/enrico-laboratory/music-projects-database-notionapi/internal/models/unparsedmodels"
	"net/http"
)

const (
	objectPages            = "pages"
	objectDatabases        = "databases"
	MusicProjectDatabaseId = "a290b3f737f242b697959acb9f18283e"
	RepertoireDatabaseId   = "aaff3c7cc9fd4c49a9d0cead3d51f75b"
	CastDatabaseId         = "065066b1d60a478a9d8a2725b2fee660"
	//LocationDatabaseId     = "264aecbf12ec4d439a306d3360a70001"
	TaskDatabaseId = "b25b103631b34fc2a8804eec0b2b3813"
)

func (c *Client) GetMusicProjects(body string, musicProject *unparsedmodels.MusicProjects) error {
	resp, err := c.request(MusicProjectDatabaseId, http.MethodPost, objectDatabases, body)
	if err != nil {
		return err
	}
	if err = readJSON(resp, &musicProject); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetMusicProjectsByStatus(status string, musicProject *unparsedmodels.MusicProjects) error {
	body := `{
    	"filter": {
        	"property": "Status",
        	"select": {
            	"equals": "` + status + `"
        	}
    	}
	}`
	resp, err := c.request(MusicProjectDatabaseId, http.MethodPost, objectDatabases, body)
	if err != nil {
		return err
	}
	if err = readJSON(resp, &musicProject); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetRepertoireByProjectId(projectId string, repertoire *unparsedmodels.Repertoire) error {
	body := `{
    	"filter": {
        	"property": "Music Project",
        	"relation": {
            	"contains": "` + projectId + `"
        	}
    	}
	}`
	resp, err := c.request(RepertoireDatabaseId, http.MethodPost, objectDatabases, body)
	if err != nil {
		return err
	}

	if err = readJSON(resp, &repertoire); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetCastByProjectId(projectId string, cast *unparsedmodels.Cast) error {
	body := `{
    	"filter": {
        	"property": "Music Project",
        	"relation": {
            	"contains": "` + projectId + `"
        	}
    	}
	}`
	resp, err := c.request(CastDatabaseId, http.MethodPost, objectDatabases, body)
	if err != nil {
		return err
	}
	if err = readJSON(resp, &cast); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetScheduleByProjectId(projectId string, schedule *unparsedmodels.Schedule) error {
	body := `{
	 "filter": {
	   "and": [
	     {
	       "property": "Music Project",
	       "relation": {
	         "contains": "` + projectId + `"
	       }
	     },
	     {
	       "or": [
	         {
	           "property": "Type",
	           "select": {
	             "equals": "Rehearsal"
	           }
	         },
	         {
	           "property": "Type",
	           "select": {
	             "equals": "Concert"
	           }
	         }
	       ]
	     }
	   ]
	 }
	}`
	resp, err := c.request(TaskDatabaseId, http.MethodPost, objectDatabases, body)
	if err != nil {
		return err
	}
	if err = readJSON(resp, &schedule); err != nil {
		return err
	}
	return nil
}

// GetMusicProjectByID SINGLE PAGE
func (c *Client) GetMusicProjectByID(id string, musicProject *unparsedmodels.MusicProject) error {
	resp, err := c.request(id, http.MethodGet, objectPages, "")
	if err != nil {
		return err
	}
	if err = readJSON(resp, &musicProject); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetLocationById(locationID string, location *unparsedmodels.Location) error {
	resp, err := c.request(locationID, http.MethodGet, objectPages, "")
	if err != nil {
		return err
	}
	if err = readJSON(resp, &location); err != nil {
		return err
	}
	return nil
}
