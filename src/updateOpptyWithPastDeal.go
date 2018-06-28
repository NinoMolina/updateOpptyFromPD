package main

import (
	"./util"
	"./model"
	"./walleClient"
	"fmt"
	"io"
	"log"
	"bytes"
	"strings"
)

const (
	Home_File_Dir = "/Users/Nino/Projects/pds/raiders-scripts/golang"
	Input_File = Home_File_Dir + "/input.csv"
	Output_Ok_File = Home_File_Dir + "/output_ok.csv"
	Output_Fail_File = Home_File_Dir + "/output_fail.csv"
)


func main() {
	util.CreateFile(Output_Ok_File, true)
	util.CreateFile(Output_Fail_File, true)

	csvr, file := util.ReadCsvFile(Input_File)
	defer util.CloseCsvFile(file)
	for {

		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				util.CheckErr(err, "")
			}
		}

		if len(row) < 2 {
			log.Println(fmt.Sprintf("missing opptyId or pastDealId: %s",row))
			util.AppendStringToFile(Output_Fail_File, fmt.Sprintf("%s,%s,s%\n", nil, nil, fmt.Sprintf("s%: s%","missing opptyId or pastDealId", row)))
			continue
		}

		opptyId := row[0]
		pastDealId := row[1]

		processOppty(opptyId, pastDealId)

	}
}

func processOppty(opptyId string, pastDealId string) {
	defer func() {
		var buffer bytes.Buffer
		buffer.WriteString("--> Start DEFER")
		err := recover()
		if err != nil {
			buffer.WriteString(fmt.Sprintf(" --> saving opptyId %s with error", opptyId))
			util.AppendStringToFile(Output_Fail_File, fmt.Sprintf("%s,%s,%s\n", opptyId, pastDealId, err))
		} else {
			buffer.WriteString(fmt.Sprintf(" --> no error for opptyId %s", opptyId))
		}
		buffer.WriteString(" --> End DEFER")
		log.Println(buffer.String())
	}()

	log.Println(fmt.Sprintf("Oppty Id: %s - Past Deal Id: %s ", opptyId, pastDealId))
	pastDealGdoc := walleClient.GetPastDeal(pastDealId)
	opptyToEditGdoc := walleClient.GetOpptyToEdit(opptyId)

	if opptyToEditGdoc.Opportunity.CdDealID > 0 {
		util.CheckErr(fmt.Errorf("Oppty to edit with cdDealId: %s", opptyId), "")
	}

	updateOpptyWithPastDeal(opptyToEditGdoc, pastDealGdoc)
	updateWorkflowStatus(opptyToEditGdoc)

	util.AppendStringToFile(Output_Ok_File, fmt.Sprintf("%s,%s\n", opptyId, pastDealId))
}

func updateOpptyWithPastDeal(opptyToEditGdoc model.GDoc, pastDealGdoc model.GDoc) {

	opptyToEditGdoc.Opportunity.Title = pastDealGdoc.Opportunity.Title
	opptyToEditGdoc.Opportunity.ShortTitle = pastDealGdoc.Opportunity.ShortTitle
	opptyToEditGdoc.Opportunity.CouponTitle = pastDealGdoc.Opportunity.CouponTitle
	opptyToEditGdoc.Opportunity.NlTitle = pastDealGdoc.Opportunity.NlTitle
	opptyToEditGdoc.Opportunity.Description = pastDealGdoc.Opportunity.Description

	opptyToEditGdoc.Images = pastDealGdoc.Images

	if !util.IsEmpty(opptyToEditGdoc.Account.Website) && !strings.HasPrefix(opptyToEditGdoc.Account.Website, "http") {
		opptyToEditGdoc.Account.Website = "http://" + opptyToEditGdoc.Account.Website
	}

	// SAVE oppty
	resp, err := walleClient.UpdateGdoc(opptyToEditGdoc)
	util.CheckErr(err, "")
	if resp.StatusCode() != 204 {
		util.CheckErr(fmt.Errorf("ERROR"), fmt.Sprintf("could not update opportunity %s, walleClient status %s", opptyToEditGdoc.Opportunity.SfOpportunityID, resp.Status))
	}

}

func updateWorkflowStatus(opptyToEditGdoc model.GDoc) {
	// SAVE workflow status
	walleClient.UpdateWorkflow(opptyToEditGdoc.Opportunity.ID,2)
	walleClient.UpdateWorkflow(opptyToEditGdoc.Opportunity.ID,3)
}




