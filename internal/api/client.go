package api

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/tagmanager/v2"
)

type ClientOptions struct {
	CredentialFile string
	AccountId      string
	ContainerId    string
	RetryLimit     int
}

type Client struct {
	*tagmanager.Service

	Options *ClientOptions
}

func NewClient(opts *ClientOptions) (*Client, error) {
	var ctx = context.Background()

	srv, err := tagmanager.NewService(ctx, option.WithCredentialsFile(opts.CredentialFile))
	if err != nil {
		return nil, err
	}

	return &Client{Service: srv, Options: opts}, nil
}

func (c *Client) containerPath() string {
	opts := c.Options
	return "accounts/" + opts.AccountId + "/containers/" + opts.ContainerId
}

var ErrNotExist = errors.New("not exist")

func (c *Client) CreateWorkspace(ws *tagmanager.Workspace) (*tagmanager.Workspace, error) {
	return getResponseWithRetry(c.Accounts.Containers.Workspaces.Create(c.containerPath(), ws).Do, c.Options.RetryLimit)
}

func (c *Client) ListWorkspaces() ([]*tagmanager.Workspace, error) {
	resp, err := getResponseWithRetry(c.Accounts.Containers.Workspaces.List(c.containerPath()).Do, c.Options.RetryLimit)
	if err != nil {
		return nil, err
	} else {
		return resp.Workspace, nil
	}
}

func (c *Client) Workspace(id string) (*tagmanager.Workspace, error) {
	ws, err := getResponseWithRetry(c.Accounts.Containers.Workspaces.Get(c.containerPath()+"/workspaces/"+id).Do, c.Options.RetryLimit)
	if errTyped, ok := err.(*googleapi.Error); ok && errTyped.Code == 404 {
		return nil, ErrNotExist
	} else {
		return ws, err
	}
}

func (c *Client) UpdateWorkspaces(id string, ws *tagmanager.Workspace) (*tagmanager.Workspace, error) {
	return getResponseWithRetry(c.Accounts.Containers.Workspaces.Update(c.containerPath()+"/workspaces/"+id, ws).Do, c.Options.RetryLimit)
}

func (c *Client) DeleteWorkspace(id string) error {
	return executeWithRetry(c.Accounts.Containers.Workspaces.Delete(c.containerPath()+"/workspaces/"+id).Do, c.Options.RetryLimit)
}

func (c *Client) workspacePath(id string) string {
	return c.containerPath() + "/workspaces/" + id
}

func (c *Client) CreateTag(workspaceId string, tag *tagmanager.Tag) (*tagmanager.Tag, error) {
	return getResponseWithRetry(c.Accounts.Containers.Workspaces.Tags.Create(c.workspacePath(workspaceId), tag).Do, c.Options.RetryLimit)
}

func (c *Client) ListTags(workspaceId string) ([]*tagmanager.Tag, error) {
	resp, err := getResponseWithRetry(c.Accounts.Containers.Workspaces.Tags.List(c.workspacePath(workspaceId)).Do, c.Options.RetryLimit)
	if err != nil {
		return nil, err
	} else {
		return resp.Tag, nil
	}
}

func (c *Client) Tag(workspaceId string, tagId string) (*tagmanager.Tag, error) {
	tag, err := getResponseWithRetry(c.Accounts.Containers.Workspaces.Tags.Get(c.workspacePath(workspaceId)+"/tags/"+tagId).Do, c.Options.RetryLimit)

	if errTyped, ok := err.(*googleapi.Error); ok && errTyped.Code == 404 {
		return nil, ErrNotExist
	} else {
		return tag, err
	}
}

func (c *Client) UpdateTag(workspaceId string, tagId string, tag *tagmanager.Tag) (*tagmanager.Tag, error) {
	return getResponseWithRetry(c.Accounts.Containers.Workspaces.Tags.Update(c.workspacePath(workspaceId)+"/tags/"+tagId, tag).Do, c.Options.RetryLimit)
}

func (c *Client) DeleteTag(workspaceId string, tagId string) error {
	return executeWithRetry(c.Accounts.Containers.Workspaces.Tags.Delete(c.workspacePath(workspaceId)+"/tags/"+tagId).Do, c.Options.RetryLimit)
}

func (c *Client) CreateVariable(workspaceId string, variable *tagmanager.Variable) (*tagmanager.Variable, error) {
	return getResponseWithRetry(c.Accounts.Containers.Workspaces.Variables.Create(c.workspacePath(workspaceId), variable).Do, c.Options.RetryLimit)
}

func (c *Client) ListVariables(workspaceId string) ([]*tagmanager.Variable, error) {
	resp, err := getResponseWithRetry(c.Accounts.Containers.Workspaces.Variables.List(c.workspacePath(workspaceId)).Do, c.Options.RetryLimit)
	if err != nil {
		return nil, err
	} else {
		return resp.Variable, nil
	}
}

func (c *Client) Variable(workspaceId string, variableId string) (*tagmanager.Variable, error) {
	variable, err := getResponseWithRetry(c.Accounts.Containers.Workspaces.Variables.Get(c.workspacePath(workspaceId)+"/variables/"+variableId).Do, c.Options.RetryLimit)

	if errTyped, ok := err.(*googleapi.Error); ok && errTyped.Code == 404 {
		return nil, ErrNotExist
	} else {
		return variable, err
	}
}

func (c *Client) UpdateVariable(workspaceId string, variableId string, variable *tagmanager.Variable) (*tagmanager.Variable, error) {
	return getResponseWithRetry(c.Accounts.Containers.Workspaces.Variables.Update(c.workspacePath(workspaceId)+"/variables/"+variableId, variable).Do, c.Options.RetryLimit)
}

func (c *Client) DeleteVariable(workspaceId string, variableId string) error {
	return executeWithRetry(c.Accounts.Containers.Workspaces.Variables.Delete(c.workspacePath(workspaceId)+"/variables/"+variableId).Do, c.Options.RetryLimit)
}

func (c *Client) CreateTrigger(workspaceId string, trigger *tagmanager.Trigger) (*tagmanager.Trigger, error) {
	return getResponseWithRetry(c.Accounts.Containers.Workspaces.Triggers.Create(c.workspacePath(workspaceId), trigger).Do, c.Options.RetryLimit)
}

func (c *Client) ListTriggers(workspaceId string) ([]*tagmanager.Trigger, error) {
	resp, err := getResponseWithRetry(c.Accounts.Containers.Workspaces.Triggers.List(c.workspacePath(workspaceId)).Do, c.Options.RetryLimit)
	if err != nil {
		return nil, err
	} else {
		return resp.Trigger, nil
	}
}

func (c *Client) Trigger(workspaceId string, triggerId string) (*tagmanager.Trigger, error) {
	trigger, err := getResponseWithRetry(c.Accounts.Containers.Workspaces.Triggers.Get(c.workspacePath(workspaceId)+"/triggers/"+triggerId).Do, c.Options.RetryLimit)

	if errTyped, ok := err.(*googleapi.Error); ok && errTyped.Code == 404 {
		return nil, ErrNotExist
	} else {
		return trigger, err
	}
}

func (c *Client) UpdateTrigger(workspaceId string, triggerId string, trigger *tagmanager.Trigger) (*tagmanager.Trigger, error) {
	return getResponseWithRetry(c.Accounts.Containers.Workspaces.Triggers.Update(c.workspacePath(workspaceId)+"/triggers/"+triggerId, trigger).Do, c.Options.RetryLimit)
}

func (c *Client) DeleteTrigger(workspaceId string, triggerId string) error {
	return executeWithRetry(c.Accounts.Containers.Workspaces.Triggers.Delete(c.workspacePath(workspaceId)+"/triggers/"+triggerId).Do, c.Options.RetryLimit)
}

// When rate limited, we wait {backoffFactor} seconds for the first retry, {2 * backoffFactor} for the second retry, {3 * backoffFactor} for the third, etc.
var backoffFactor = time.Duration(20)

func executeWithRetry(query func(opts ...googleapi.CallOption) error, retryLimit int) error {
	retryCount := 0

	for {
		err := query()
		if errTyped, ok := err.(*googleapi.Error); ok && errTyped.Code == 429 {
			if retryCount < retryLimit {
				retryCount++
				backoffDuration := backoffFactor * time.Second * time.Duration(retryCount)
				fmt.Printf("Rate limited. Retrying in %s...\n", backoffDuration)
				time.Sleep(backoffDuration)
				continue
			} else {
				return fmt.Errorf("Still rate limited after %d retries", retryLimit)
			}
		} else if err != nil {
			return err
		} else {
			return nil
		}
	}
}

func getResponseWithRetry[R any](query func(opts ...googleapi.CallOption) (*R, error), retryLimit int) (*R, error) {
	retryCount := 0

	for {
		resp, err := query()
		if errTyped, ok := err.(*googleapi.Error); ok && errTyped.Code == 429 {
			if retryCount < retryLimit {
				retryCount++
				backoffDuration := backoffFactor * time.Second * time.Duration(retryCount)
				fmt.Printf("Rate limited. Retrying in %s...\n", backoffDuration)
				time.Sleep(backoffDuration)
				continue
			} else {
				return nil, fmt.Errorf("Still rate limited exceeded after %d retries", retryLimit)
			}
		} else if err != nil {
			return nil, err
		} else {
			return resp, nil
		}
	}
}
