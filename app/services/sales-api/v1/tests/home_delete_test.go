package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/ardanlabs/service/business/web/v1/response"
	"github.com/google/go-cmp/cmp"
)

func homeDelete200(t *testing.T, app appTest, sd seedData) []tableData {
	table := []tableData{
		{
			name:       "asuser",
			url:        fmt.Sprintf("/v1/homes/%s", sd.users[0].homes[0].ID),
			token:      sd.users[0].token,
			method:     http.MethodDelete,
			statusCode: http.StatusNoContent,
		},
		{
			name:       "asadmin",
			url:        fmt.Sprintf("/v1/homes/%s", sd.admins[0].homes[0].ID),
			token:      sd.admins[0].token,
			method:     http.MethodDelete,
			statusCode: http.StatusNoContent,
		},
	}

	return table
}

func homeDelete401(t *testing.T, app appTest, sd seedData) []tableData {
	table := []tableData{
		{
			name:       "emptytoken",
			url:        fmt.Sprintf("/v1/homes/%s", sd.users[0].homes[1].ID),
			token:      "",
			method:     http.MethodDelete,
			statusCode: http.StatusUnauthorized,
			resp:       &response.ErrorDocument{},
			expResp:    &response.ErrorDocument{Error: "Unauthorized"},
			cmpFunc: func(x interface{}, y interface{}) string {
				return cmp.Diff(x, y)
			},
		},
		{
			name:       "badsig",
			url:        fmt.Sprintf("/v1/homes/%s", sd.users[0].homes[1].ID),
			token:      sd.users[0].token + "A",
			method:     http.MethodDelete,
			statusCode: http.StatusUnauthorized,
			resp:       &response.ErrorDocument{},
			expResp:    &response.ErrorDocument{Error: "Unauthorized"},
			cmpFunc: func(x interface{}, y interface{}) string {
				return cmp.Diff(x, y)
			},
		},
		{
			name:       "wronguser",
			url:        fmt.Sprintf("/v1/homes/%s", sd.users[0].homes[1].ID),
			token:      app.userToken,
			method:     http.MethodDelete,
			statusCode: http.StatusUnauthorized,
			resp:       &response.ErrorDocument{},
			expResp:    &response.ErrorDocument{Error: "Unauthorized"},
			cmpFunc: func(x interface{}, y interface{}) string {
				return cmp.Diff(x, y)
			},
		},
	}

	return table
}
