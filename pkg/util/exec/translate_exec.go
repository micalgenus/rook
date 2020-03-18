/*
Copyright 2016 The Rook Authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package exec

import (
	"time"
)

// TranslateCommandExecutor is an exec.Executor that translates every command before executing it
// This is useful to run the commands in a job with `kubectl run ...` when running the operator outside
// of Kubernetes and need to run tools that require running inside the cluster.
type TranslateCommandExecutor struct {

	// Executor is probably a exec.CommandExecutor that will run the translated commands
	Executor Executor

	// Translator translates every command before running it
	Translator func(suppressLogOutput bool, command string, arg ...string) (string, []string)
}

// ExecuteCommand starts a process and wait for its completion
func (e *TranslateCommandExecutor) ExecuteCommand(suppressLogOutput bool, command string, arg ...string) error {
	transCommand, transArgs := e.Translator(suppressLogOutput, command, arg...)
	return e.Executor.ExecuteCommand(suppressLogOutput, transCommand, transArgs...)
}

// ExecuteCommandWithOutput starts a process and wait for its completion
func (e *TranslateCommandExecutor) ExecuteCommandWithOutput(suppressLogOutput bool, command string, arg ...string) (string, error) {
	transCommand, transArgs := e.Translator(suppressLogOutput, command, arg...)
	return e.Executor.ExecuteCommandWithOutput(suppressLogOutput, transCommand, transArgs...)
}

// ExecuteCommandWithCombinedOutput starts a process and returns its stdout and stderr combined.
func (e *TranslateCommandExecutor) ExecuteCommandWithCombinedOutput(suppressLogOutput bool, command string, arg ...string) (string, error) {
	transCommand, transArgs := e.Translator(suppressLogOutput, command, arg...)
	return e.Executor.ExecuteCommandWithCombinedOutput(suppressLogOutput, transCommand, transArgs...)
}

// ExecuteCommandWithOutputFile starts a process and saves output to file
func (e *TranslateCommandExecutor) ExecuteCommandWithOutputFile(suppressLogOutput bool, command, outfileArg string, arg ...string) (string, error) {
	transCommand, transArgs := e.Translator(suppressLogOutput, command, arg...)
	return e.Executor.ExecuteCommandWithOutputFile(suppressLogOutput, transCommand, outfileArg, transArgs...)
}

// ExecuteCommandWithOutputFileTimeout is the same as ExecuteCommandWithOutputFile but with a timeout limit.
func (e *TranslateCommandExecutor) ExecuteCommandWithOutputFileTimeout(
	suppressLogOutput bool, timeout time.Duration,
	command, outfileArg string, arg ...string) (string, error) {
	transCommand, transArgs := e.Translator(suppressLogOutput, command, arg...)
	return e.Executor.ExecuteCommandWithOutputFileTimeout(suppressLogOutput, timeout, transCommand, outfileArg, transArgs...)
}

// ExecuteCommandWithTimeout starts a process and wait for its completion with timeout.
func (e *TranslateCommandExecutor) ExecuteCommandWithTimeout(suppressLogOutput bool, timeout time.Duration, command string, arg ...string) (string, error) {
	transCommand, transArgs := e.Translator(suppressLogOutput, command, arg...)
	return e.Executor.ExecuteCommandWithTimeout(suppressLogOutput, timeout, transCommand, transArgs...)
}
