package t38c

import "strconv"

// Scripting struct
type Scripting struct {
	client *Client
}

func makeScriptArgs(scriptOrSHA string, keys []string, args []string) []string {
	cmdArgs := []string{
		scriptOrSHA,
		strconv.Itoa(len(keys)),
	}
	cmdArgs = append(cmdArgs, keys...)
	cmdArgs = append(cmdArgs, args...)
	return cmdArgs
}

// Eval evaluates a Lua script
func (sc *Scripting) Eval(script string, keys []string, args []string) ([]byte, error) {
	return sc.client.Execute("EVAL", makeScriptArgs(script, keys, args)...)
}

// EvalNA evaluates a Lua script in a non-atomic fashion.
// The command uses None atomicity level and is otherwise identical to EVAL.
func (sc *Scripting) EvalNA(script string, keys []string, args []string) ([]byte, error) {
	return sc.client.Execute("EVALNA", makeScriptArgs(script, keys, args)...)
}

// EvalNASHA evaluates, in a non-atomic fashion, a Lua script cached on the server by its SHA1 digest.
// Scripts are cached using the SCRIPT LOAD command. The command is otherwise identical to EVALNA.
func (sc *Scripting) EvalNASHA(sha string, keys []string, args []string) ([]byte, error) {
	return sc.client.Execute("EVALNASHA", makeScriptArgs(sha, keys, args)...)
}

// EvalRO evaluates a read-only Lua script.
// The command uses Read-only atomicity level and is otherwise identical to EVAL.
func (sc *Scripting) EvalRO(script string, keys []string, args []string) ([]byte, error) {
	return sc.client.Execute("EVALRO", makeScriptArgs(script, keys, args)...)
}

// EvalROSHA evaluates a read-only Lua script cached on the server by its SHA1 digest.
// Scripts are cached using the SCRIPT LOAD command. The command is otherwise identical to EVALRO.
func (sc *Scripting) EvalROSHA(sha string, keys []string, args []string) ([]byte, error) {
	return sc.client.Execute("EVALROSHA", makeScriptArgs(sha, keys, args)...)
}

// EvalSHA evaluates a Lua script cached on the server by its SHA1 digest.
// Scripts are cached using the SCRIPT LOAD command. The command is otherwise identical to EVAL.
func (sc *Scripting) EvalSHA(sha string, keys []string, args []string) ([]byte, error) {
	return sc.client.Execute("EVALSHA", makeScriptArgs(sha, keys, args)...)
}

// ScriptExists returns information about the existence of the scripts in server cache.
// This command takes one or more SHA1 digests and returns a list of one/zero integer values
// to indicate whether or not each SHA1 exists in the server script cache.
// Scripts are cached using the SCRIPT LOAD command.
func (sc *Scripting) ScriptExists(shas ...string) ([]int, error) {
	var resp struct {
		Result []int
	}
	err := sc.client.jExecute(&resp, "SCRIPT", append([]string{"EXISTS"}, shas...)...)
	return resp.Result, err
}

// ScriptFlush flushes the server cache of Lua scripts.
func (sc *Scripting) ScriptFlush() error {
	return sc.client.jExecute(nil, "SCRIPT", "FLUSH")
}

// ScriptLoad loads the compiled version of a script into the server cache, without executing.
// If the parsing and compilation is successful, the command returns the string value
// of the SHA1 digest of the script. That value can be used for EVALSHA and similar commands
// that execute scripts based on the SHA1 digest.
// The script will stay in cache until either the tile38 is restarted or SCRIPT FLUSH is called.
// If either parsing or compilation fails, the command will return the error response
// with the detailed traceback of the Lua failure.
func (sc *Scripting) ScriptLoad(script string) error {
	return sc.client.jExecute(nil, "SCRIPT", append([]string{"LOAD"}, script)...)
}
