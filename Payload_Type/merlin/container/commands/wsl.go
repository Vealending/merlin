/*
Merlin is a post-exploitation command and control framework.

This file is part of Merlin.
Copyright (C) 2024  Russel Van Tuyl

Merlin is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as
published by the Free Software Foundation, either version 3 of the License, or any later version.

Merlin is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with Merlin.  If not, see <http://www.gnu.org/licenses/>.
*/

package commands

import (
	// Standard
	"fmt"

	// Mythic
	structs "github.com/MythicMeta/MythicContainer/agent_structs"

	// Merlin Message
	"github.com/Ne0nd0g/merlin-message/jobs"
)

// wsl creates and returns a Mythic Command structure that is registered with the Mythic server
func wsl() structs.Command {
	attr := structs.CommandAttribute{
		SupportedOS: []string{structs.SUPPORTED_OS_WINDOWS},
	}

	action := structs.CommandParameter{
		Name:                                    "action",
		ModalDisplayName:                        "Action",
		CLIName:                                 "action",
		ParameterType:                           structs.COMMAND_PARAMETER_TYPE_CHOOSE_ONE,
		Description:                             "The WSL action to perform",
		Choices:                                 []string{"list", "exec"},
		DefaultValue:                            nil,
		SupportedAgents:                         nil,
		SupportedAgentBuildParameters:           nil,
		ChoicesAreAllCommands:                   false,
		ChoicesAreLoadedCommands:                false,
		FilterCommandChoicesByCommandAttributes: nil,
		DynamicQueryFunction:                    nil,
		ParameterGroupInformation: []structs.ParameterGroupInfo{
			{
				ParameterIsRequired:   true,
				GroupName:             "List",
				UIModalPosition:       0,
				AdditionalInformation: nil,
			},
			{
				ParameterIsRequired:   true,
				GroupName:             "Exec",
				UIModalPosition:       0,
				AdditionalInformation: nil,
			},
		},
	}

	distro := structs.CommandParameter{
		Name:                                    "distro",
		ModalDisplayName:                        "Distribution Name",
		CLIName:                                 "distro",
		ParameterType:                           structs.COMMAND_PARAMETER_TYPE_STRING,
		Description:                             "The name of the WSL distribution to execute the command in (e.g., Ubuntu)",
		Choices:                                 nil,
		DefaultValue:                            nil,
		SupportedAgents:                         nil,
		SupportedAgentBuildParameters:           nil,
		ChoicesAreAllCommands:                   false,
		ChoicesAreLoadedCommands:                false,
		FilterCommandChoicesByCommandAttributes: nil,
		DynamicQueryFunction:                    nil,
		ParameterGroupInformation: []structs.ParameterGroupInfo{
			{
				ParameterIsRequired:   true,
				GroupName:             "Exec",
				UIModalPosition:       1,
				AdditionalInformation: nil,
			},
		},
	}

	command := structs.CommandParameter{
		Name:                                    "command",
		ModalDisplayName:                        "Command",
		CLIName:                                 "command",
		ParameterType:                           structs.COMMAND_PARAMETER_TYPE_STRING,
		Description:                             "The command to execute inside the WSL distribution",
		Choices:                                 nil,
		DefaultValue:                            nil,
		SupportedAgents:                         nil,
		SupportedAgentBuildParameters:           nil,
		ChoicesAreAllCommands:                   false,
		ChoicesAreLoadedCommands:                false,
		FilterCommandChoicesByCommandAttributes: nil,
		DynamicQueryFunction:                    nil,
		ParameterGroupInformation: []structs.ParameterGroupInfo{
			{
				ParameterIsRequired:   true,
				GroupName:             "Exec",
				UIModalPosition:       2,
				AdditionalInformation: nil,
			},
		},
	}

	params := []structs.CommandParameter{action, distro, command}
	cmd := structs.Command{
		Name:                           "wsl",
		NeedsAdminPermissions:          false,
		HelpString:                     "wsl list | wsl exec <distro> <command>",
		Description:                    "Interact with Windows Subsystem for Linux distributions via direct COM interface calls without spawning wsl.exe",
		Version:                        0,
		SupportedUIFeatures:            nil,
		Author:                         "@Vealending",
		MitreAttackMappings:            []string{"T1059", "T1106"},
		ScriptOnlyCommand:              false,
		CommandAttributes:              attr,
		CommandParameters:              params,
		AssociatedBrowserScript:        nil,
		TaskFunctionOPSECPre:           nil,
		TaskFunctionCreateTasking:      wslCreateTask,
		TaskFunctionProcessResponse:    nil,
		TaskFunctionOPSECPost:          nil,
		TaskFunctionParseArgString:     taskFunctionParseArgString,
		TaskFunctionParseArgDictionary: taskFunctionParseArgDictionary,
		TaskCompletionFunctions:        nil,
	}

	return cmd
}

// wslCreateTask takes a Mythic Task and converts it into a Merlin Job that is encoded into JSON and subsequently sent to the Merlin Agent
func wslCreateTask(task *structs.PTTaskMessageAllData) (resp structs.PTTaskCreateTaskingMessageResponse) {
	pkg := "mythic/container/commands/wsl/wslCreateTask()"
	resp.TaskID = task.Task.ID

	action, err := task.Args.GetStringArg("action")
	if err != nil {
		resp.Error = fmt.Sprintf("%s: %s", pkg, err)
		resp.Success = false
		return
	}

	args := []string{action}

	if action == "exec" {
		distroName, err := task.Args.GetStringArg("distro")
		if err != nil {
			resp.Error = fmt.Sprintf("%s: %s", pkg, err)
			resp.Success = false
			return
		}

		cmd, err := task.Args.GetStringArg("command")
		if err != nil {
			resp.Error = fmt.Sprintf("%s: %s", pkg, err)
			resp.Success = false
			return
		}

		args = append(args, distroName, cmd)
	}

	job := jobs.Command{
		Command: task.Task.CommandName,
		Args:    args,
	}

	mythicJob, err := ConvertMerlinJobToMythicTask(job, jobs.MODULE)
	if err != nil {
		resp.Error = fmt.Sprintf("%s: %s", pkg, err)
		resp.Success = false
		return
	}

	task.Args.SetManualArgs(mythicJob)

	var disp string
	if action == "list" {
		disp = "list"
	} else {
		disp = fmt.Sprintf("exec %s %s", args[1], args[2])
	}
	resp.DisplayParams = &disp
	resp.Success = true

	return
}
