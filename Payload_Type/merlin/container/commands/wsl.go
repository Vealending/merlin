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
	"encoding/base64"
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
		Choices:                                 []string{"list", "exec", "import", "import-local", "unregister", "terminate"},
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
			{
				ParameterIsRequired:   true,
				GroupName:             "Import",
				UIModalPosition:       0,
				AdditionalInformation: nil,
			},
			{
				ParameterIsRequired:   true,
				GroupName:             "Import Local",
				UIModalPosition:       0,
				AdditionalInformation: nil,
			},
			{
				ParameterIsRequired:   true,
				GroupName:             "Unregister",
				UIModalPosition:       0,
				AdditionalInformation: nil,
			},
			{
				ParameterIsRequired:   true,
				GroupName:             "Terminate",
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
		Description:                             "The name of the WSL distribution (e.g., Ubuntu, RedOps)",
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
			{
				ParameterIsRequired:   true,
				GroupName:             "Import",
				UIModalPosition:       1,
				AdditionalInformation: nil,
			},
			{
				ParameterIsRequired:   true,
				GroupName:             "Import Local",
				UIModalPosition:       1,
				AdditionalInformation: nil,
			},
			{
				ParameterIsRequired:   true,
				GroupName:             "Unregister",
				UIModalPosition:       1,
				AdditionalInformation: nil,
			},
			{
				ParameterIsRequired:   true,
				GroupName:             "Terminate",
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

	file := structs.CommandParameter{
		Name:                                    "file",
		ModalDisplayName:                        "Tarball",
		CLIName:                                 "file",
		ParameterType:                           structs.COMMAND_PARAMETER_TYPE_FILE,
		Description:                             "The rootfs tarball to import (tar.gz or tar.zst). Streamed to WSL via pipe — never touches disk.",
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
				GroupName:             "Import",
				UIModalPosition:       2,
				AdditionalInformation: nil,
			},
		},
	}

	path := structs.CommandParameter{
		Name:                                    "path",
		ModalDisplayName:                        "Tarball Path",
		CLIName:                                 "path",
		ParameterType:                           structs.COMMAND_PARAMETER_TYPE_STRING,
		Description:                             "Full path to a rootfs tarball already on the target (e.g., C:\\Temp\\redops.tar.zst)",
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
				GroupName:             "Import Local",
				UIModalPosition:       2,
				AdditionalInformation: nil,
			},
		},
	}

	targetDir := structs.CommandParameter{
		Name:                                    "target_dir",
		ModalDisplayName:                        "Install Directory",
		CLIName:                                 "target_dir",
		ParameterType:                           structs.COMMAND_PARAMETER_TYPE_STRING,
		Description:                             "Optional install directory for the distro VHD. Defaults to WSL's standard location if empty.",
		Choices:                                 nil,
		DefaultValue:                            "",
		SupportedAgents:                         nil,
		SupportedAgentBuildParameters:           nil,
		ChoicesAreAllCommands:                   false,
		ChoicesAreLoadedCommands:                false,
		FilterCommandChoicesByCommandAttributes: nil,
		DynamicQueryFunction:                    nil,
		ParameterGroupInformation: []structs.ParameterGroupInfo{
			{
				ParameterIsRequired:   false,
				GroupName:             "Import",
				UIModalPosition:       3,
				AdditionalInformation: nil,
			},
			{
				ParameterIsRequired:   false,
				GroupName:             "Import Local",
				UIModalPosition:       3,
				AdditionalInformation: nil,
			},
		},
	}

	params := []structs.CommandParameter{action, distro, command, file, path, targetDir}
	cmd := structs.Command{
		Name:                           "wsl",
		NeedsAdminPermissions:          false,
		HelpString:                     "wsl list | wsl exec <distro> <command> | wsl import <distro> <tarball> | wsl import-local <distro> <path> | wsl unregister <distro> | wsl terminate <distro>",
		Description:                    "Interact with Windows Subsystem for Linux distributions via direct COM interface calls without spawning wsl.exe. Supports listing, exec, importing, unregistering, and terminating distributions.",
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
	var disp string

	switch action {
	case "list":
		disp = "list"

	case "exec":
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
		disp = fmt.Sprintf("exec %s %s", distroName, cmd)

	case "import":
		distroName, err := task.Args.GetStringArg("distro")
		if err != nil {
			resp.Error = fmt.Sprintf("%s: %s", pkg, err)
			resp.Success = false
			return
		}
		// Get uploaded file contents
		fileID, err := task.Args.GetStringArg("file")
		if err != nil {
			resp.Error = fmt.Sprintf("%s: %s", pkg, err)
			resp.Success = false
			return
		}
		data, err := GetFileContents(fileID)
		if err != nil {
			resp.Error = fmt.Sprintf("%s: failed to retrieve uploaded file: %s", pkg, err)
			resp.Success = false
			return
		}
		targetDir, _ := task.Args.GetStringArg("target_dir")
		args = append(args, distroName, base64.StdEncoding.EncodeToString(data), targetDir)
		disp = fmt.Sprintf("import %s (%d bytes)", distroName, len(data))

	case "import-local":
		distroName, err := task.Args.GetStringArg("distro")
		if err != nil {
			resp.Error = fmt.Sprintf("%s: %s", pkg, err)
			resp.Success = false
			return
		}
		filePath, err := task.Args.GetStringArg("path")
		if err != nil {
			resp.Error = fmt.Sprintf("%s: %s", pkg, err)
			resp.Success = false
			return
		}
		targetDir, _ := task.Args.GetStringArg("target_dir")
		args = append(args, distroName, filePath, targetDir)
		disp = fmt.Sprintf("import-local %s %s", distroName, filePath)

	case "unregister":
		distroName, err := task.Args.GetStringArg("distro")
		if err != nil {
			resp.Error = fmt.Sprintf("%s: %s", pkg, err)
			resp.Success = false
			return
		}
		args = append(args, distroName)
		disp = fmt.Sprintf("unregister %s", distroName)

	case "terminate":
		distroName, err := task.Args.GetStringArg("distro")
		if err != nil {
			resp.Error = fmt.Sprintf("%s: %s", pkg, err)
			resp.Success = false
			return
		}
		args = append(args, distroName)
		disp = fmt.Sprintf("terminate %s", distroName)

	default:
		resp.Error = fmt.Sprintf("%s: unknown action: %s", pkg, action)
		resp.Success = false
		return
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
	resp.DisplayParams = &disp
	resp.Success = true

	return
}
