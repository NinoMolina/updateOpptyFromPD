package walleClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"../model"
	"../util"
	"gopkg.in/resty.v1"
)

const (
	Walle_URLBase = "http://localhost:9100"
	//Walle_URLBase = "http://localhost:9100"
	Gdoc_Path_Resource        = "/gdoc?opportunity="
	Update_Gdoc_Path_Resource = "/gdoc?action=images,editorial"
	Pastdeals_Resource        = "/opportunity/{}/past-deal"
	Workflow_Resource         = "/workflows/opportunities/"
	Walle_Token               = "eyJhbGciOiJIUzI1NiJ9.eyJqdGkiOiI2N2QyYzJkN2E1ZGE3ZDBmMDEwMjRlYWRhYzEyODBjOGQ0YmExYjRmIiwiaWF0IjoxNTIzMDI0NDgxLCJzdWIiOiJsbW9saW5hQGdyb3Vwb25sYXRhbS5jb20iLCJpc3MiOiJhY2NvdW50cy5nb29nbGUuY29tIiwic2NvcGUiOiJbd3JpdGVyLCBkZXNpZ25lcl0iLCJleHAiOjE1MjMxMTA4ODF9.LSCq7H54PH3SukRS3otS1xOiTfkiJ3AefEsVO33LfNU"
)

func UpdateGdoc(gdoc model.GDoc) (*resty.Response, error) {
	resp, err := resty.R().
		SetBody(gdoc).
		SetAuthToken(Walle_Token).
		//SetError(&Error{}).       // or SetError(Error{}).
		Put(Walle_URLBase + Update_Gdoc_Path_Resource)
	util.CheckErr(err, "")
	log.Println()
	if resp.StatusCode() != 204 {
		body := resp.Body()
		var bodyString = ""
		if body != nil {
			bodyString = string(resp.Body())
		}
		util.CheckErr(fmt.Errorf("Could not update oppty - HttpStatusCode %s - %s", resp.Status(), bodyString), "")
	}
	return resp, err
}

func UpdateWorkflow(opptyId int, workflowStatus int) (*resty.Response, error) {
	var jsonStr = []byte(`{ "new_state":` + strconv.Itoa(workflowStatus) + ` }`)
	urlStr := Walle_URLBase + Workflow_Resource + strconv.Itoa(opptyId)
	response, err := resty.R().
		SetAuthToken(Walle_Token).
		SetBody(bytes.NewBuffer(jsonStr)).
		Put(urlStr)
	util.CheckErr(err, "")
	if response.StatusCode() != 204 {
		body := response.Body()
		var bodyString = ""
		if body != nil {
			bodyString = string(response.Body())
		}
		util.CheckErr(fmt.Errorf("Could not update workflow with wf_id %d - HttpStatusCode %s - %s", workflowStatus, response.Status(), bodyString), "")
	}
	return response, err
}

func GetPastDeal(pastDealId string) model.GDoc {
	_, err := resty.R().
		SetAuthToken(Walle_Token).
		Get(Walle_URLBase + strings.Replace(Pastdeals_Resource, "{}", pastDealId, -1))

	util.CheckErr(err, fmt.Sprintf("Error al traer el pastDeal %s", err))

	// GET past deal
	pastDealGdocResponse, err := resty.R().
		SetAuthToken(Walle_Token).
		Get(Walle_URLBase + Gdoc_Path_Resource + string(pastDealId))

	if pastDealGdocResponse.StatusCode() != 200 {
		body := pastDealGdocResponse.Body()
		var bodyString = ""
		if body != nil {
			bodyString = string(pastDealGdocResponse.Body())
		}
		util.CheckErr(fmt.Errorf("Could not get past deal - HttpStatusCode %s - %s", pastDealGdocResponse.Status(), bodyString), "")
	}
	var pastDealGdoc model.GDoc
	err = json.Unmarshal(pastDealGdocResponse.Body(), &pastDealGdoc)
	util.CheckErr(err, "unmarshal pastdeal "+pastDealId)

	return pastDealGdoc
}

func GetOpptyToEdit(opptyId string) model.GDoc {
	// GET oppty
	newDealGdocResponse, err := resty.R().
		SetAuthToken(Walle_Token).
		Get(Walle_URLBase + Gdoc_Path_Resource + string(opptyId))

	util.CheckErr(err, fmt.Sprintf("Error al traer oppty %s", err))

	if newDealGdocResponse.StatusCode() != 200 {
		body := newDealGdocResponse.Body()
		var bodyString = ""
		if body != nil {
			bodyString = string(newDealGdocResponse.Body())
		}
		util.CheckErr(fmt.Errorf("Could not get oppty to edit - HttpStatusCode %s - %s", newDealGdocResponse.Status(), bodyString), "")
	}
	var opptyToEditGdoc model.GDoc
	err = json.Unmarshal(newDealGdocResponse.Body(), &opptyToEditGdoc)
	util.CheckErr(err, "unmarshal pastdeal "+opptyId)

	return opptyToEditGdoc
}
