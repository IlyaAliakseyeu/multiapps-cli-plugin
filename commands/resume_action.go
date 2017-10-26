package commands

import (
	mtaclient "github.com/SAP/cf-mta-plugin/clients/mtaclient"
	"github.com/SAP/cf-mta-plugin/ui"
	"github.com/cloudfoundry/cli/cf/terminal"
)

// ResumeAction retries the process with the specified id
type ResumeAction struct{}

// Execute executes resume action on process with the specified id
func (a *ResumeAction) Execute(operationID, commandName string, mtaClient mtaclient.MtaClientOperations) ExecutionStatus {

	// TODO: Ensure session is not expired

	ui.Say("Resuming multi-target app operation with id %s...", terminal.EntityNameColor(operationID))
	responseHeader, err := mtaClient.ExecuteAction(operationID, "resume")
	if err != nil {
		ui.Failed("Could not resume multi-target app operation: %s", err)
		return Failure
	}
	ui.Ok()

	return NewExecutionMonitor(commandName, responseHeader.Location.String(), mtaClient).Monitor()
}
