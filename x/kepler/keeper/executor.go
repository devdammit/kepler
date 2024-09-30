package keeper

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/devdammit/kepler/x/kepler/types"
	"github.com/expr-lang/expr"
)

type Automation struct {
	Counter    uint64     `expr:"Counter"`
	ExecutedAt *time.Time `expr:"ExecutedAt"`
}

type Global struct {
	Time time.Time `expr:"Time"`
}

type Env struct {
	Automation Automation `expr:"automation"`
	Global     Global     `expr:"global"`
}

type ExecuteEnv struct {
	Automation Automation `expr:"automation"`
	Global     Global     `expr:"global"`
}

func (ExecuteEnv) SendWebhook(url string, payload string) bool {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return false
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func (Env) Format(t time.Time) string { // Methods defined on the struct become functions.
	return t.Format(time.RFC822)
}

func (k Keeper) ExecuteW3Func(ctx sdk.Context, id uint64) error {
	w3Func, err := k.GetW3Func(ctx, id)
	if err != nil {
		return err
	}

	isRunnable, err := k.w3funcIsRunnable(ctx, *w3Func)
	if err != nil {
		return err
	}

	fmt.Println("isRunnable: ", isRunnable)
	fmt.Println("blockTime: ", ctx.BlockTime().UTC())

	if isRunnable {
		fmt.Println("running w3Func: ", w3Func)
		program, err := expr.Compile(w3Func.Code, expr.Env(ExecuteEnv{}))
		if err != nil {
			return err
		}

		_, err = expr.Run(program, ExecuteEnv{})

		w3Func.ExecutedAt = ctx.BlockTime().UTC().String()
		w3Func.Counter++

		k.MarkW3FuncExecuted(ctx, id, w3Func.ExecutedAt)

		return nil
	}

	return nil
}

func (k Keeper) executeW3Func(ctx sdk.Context, w3Func types.W3Func) error {
	program, err := expr.Compile(w3Func.Code, expr.Env(ExecuteEnv{}))
	if err != nil {
		return err
	}

	env := ExecuteEnv{
		Automation: Automation{
			Counter: w3Func.Counter,
		},
		Global: Global{
			Time: ctx.BlockTime(),
		},
	}

	if w3Func.ExecutedAt != "" {
		executedAt, err := time.Parse(time.RFC3339, w3Func.ExecutedAt)
		if err != nil {
			return err
		}
		env.Automation.ExecutedAt = &executedAt
	}

	_, err = expr.Run(program, env)

	return err
}

func (k Keeper) w3funcIsRunnable(ctx sdk.Context, w3Func types.W3Func) (bool, error) {
	env := Env{
		Automation: Automation{
			Counter: w3Func.Counter,
		},
		Global: Global{
			Time: ctx.BlockTime(),
		},
	}

	if w3Func.ExecutedAt != "" {
		executedAt, err := time.Parse(time.RFC3339, w3Func.ExecutedAt)
		if err != nil {
			return false, err
		}
		env.Automation.ExecutedAt = &executedAt
	}

	program, err := expr.Compile(w3Func.Cond, expr.Env(Env{}))
	if err != nil {
		return false, err
	}

	output, err := expr.Run(program, env)
	if err != nil {
		return false, err
	}

	return output != nil && output.(bool), nil
}
