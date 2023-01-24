// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package ynab

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
)

func TestClient_GET(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s%s", apiEndpoint, "/foo"),
			func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "application/json", req.Header.Get("Accept"))
				assert.Equal(t, "Bearer 6zL9vh8]B9H3BEecwL%Vzh^VwKR3C2CNZ3Bv%=fFxm$z)duY[U+2=3CydZrkQFnA", req.Header.Get("Authorization"))

				res := httpmock.NewStringResponse(http.StatusOK, `{"foo":"bar"}`)
				res.Header.Add("X-Rate-Limit", "36/200")
				return res, nil
			},
		)

		response := struct {
			Foo string `json:"foo"`
		}{}

		c := NewClient("6zL9vh8]B9H3BEecwL%Vzh^VwKR3C2CNZ3Bv%=fFxm$z)duY[U+2=3CydZrkQFnA")
		err := c.(*client).GET("/foo", &response)
		assert.NoError(t, err)
		assert.Equal(t, "bar", response.Foo)
	})

	t.Run("failure with with expected API error", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s%s", apiEndpoint, "/foo"),
			func(req *http.Request) (*http.Response, error) {
				res := httpmock.NewStringResponse(http.StatusBadRequest, `{
  "error": {
    "id": "400",
    "name": "error_name",
    "detail": "Error detail"
  }
}`)
				res.Header.Add("X-Rate-Limit", "36/200")
				return res, nil
			},
		)

		response := struct {
			Foo string `json:"foo"`
		}{}

		c := NewClient("")
		err := c.(*client).GET("/foo", &response)
		expectedErrStr := "api: error id=400 name=error_name detail=Error detail"
		assert.EqualError(t, err, expectedErrStr)
	})

	t.Run("failure with with unexpected API error", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s%s", apiEndpoint, "/foo"),
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(http.StatusInternalServerError, "Internal Server Error"), nil
			},
		)

		response := struct {
			Foo string `json:"foo"`
		}{}

		c := NewClient("")
		err := c.(*client).GET("/foo", &response)
		expectedErrStr := "api: error id=500 name=unknown_api_error detail=Unknown API error"
		assert.EqualError(t, err, expectedErrStr)
	})

	t.Run("silent failure due to invalid response model", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s%s", apiEndpoint, "/foo"),
			func(req *http.Request) (*http.Response, error) {
				res := httpmock.NewStringResponse(http.StatusOK, `{"bar":"foo"}`)
				res.Header.Add("X-Rate-Limit", "36/200")
				return res, nil
			},
		)

		response := struct {
			Foo string `json:"foo"`
		}{}

		c := NewClient("")
		err := c.(*client).GET("/foo", &response)
		assert.NoError(t, err)
		assert.Equal(t, struct {
			Foo string `json:"foo"`
		}{}, response)
	})
}

func TestClient_POST(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s%s", apiEndpoint, "/foo"),
			func(req *http.Request) (*http.Response, error) {
				buf, err := io.ReadAll(req.Body)
				assert.NoError(t, err)
				assert.Equal(t, `{"bar":"foo"}`, string(buf))
				assert.Equal(t, "application/json", req.Header.Get("Accept"))
				assert.Equal(t, "Bearer 6zL9vh8]B9H3BEecwL%Vzh^VwKR3C2CNZ3Bv%=fFxm$z)duY[U+2=3CydZrkQFnA", req.Header.Get("Authorization"))

				res := httpmock.NewStringResponse(http.StatusOK, string(buf))
				res.Header.Add("X-Rate-Limit", "36/200")
				return res, nil
			},
		)

		response := struct {
			Bar string `json:"bar"`
		}{}

		c := NewClient("6zL9vh8]B9H3BEecwL%Vzh^VwKR3C2CNZ3Bv%=fFxm$z)duY[U+2=3CydZrkQFnA")
		err := c.(*client).POST("/foo", &response, []byte(`{"bar":"foo"}`))
		assert.NoError(t, err)
		assert.Equal(t, "foo", response.Bar)
	})

	t.Run("failure with with expected API error", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s%s", apiEndpoint, "/foo"),
			func(req *http.Request) (*http.Response, error) {
				res := httpmock.NewStringResponse(http.StatusBadRequest, `{
  "error": {
    "id": "400",
    "name": "error_name",
    "detail": "Error detail"
  }
}`)
				res.Header.Add("X-Rate-Limit", "36/200")
				return res, nil
			},
		)

		response := struct {
			Foo string `json:"foo"`
		}{}

		c := NewClient("")
		err := c.(*client).POST("/foo", &response, []byte(`{"bar":"foo"}`))
		expectedErrStr := "api: error id=400 name=error_name detail=Error detail"
		assert.EqualError(t, err, expectedErrStr)
	})

	t.Run("failure with with unexpected API error", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s%s", apiEndpoint, "/foo"),
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(http.StatusInternalServerError, "Internal Server Error"), nil
			},
		)

		response := struct {
			Foo string `json:"foo"`
		}{}

		c := NewClient("")
		err := c.(*client).POST("/foo", &response, []byte(`{"bar":"foo"}`))
		expectedErrStr := "api: error id=500 name=unknown_api_error detail=Unknown API error"
		assert.EqualError(t, err, expectedErrStr)
	})

	t.Run("silent failure due to invalid response model", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s%s", apiEndpoint, "/foo"),
			func(req *http.Request) (*http.Response, error) {
				res := httpmock.NewStringResponse(http.StatusOK, `{"bar":"foo"}`)
				res.Header.Add("X-Rate-Limit", "36/200")
				return res, nil
			},
		)

		response := struct {
			Foo string `json:"foo"`
		}{}

		c := NewClient("")
		err := c.(*client).POST("/foo", &response, []byte(`{"bar":"foo"}`))
		assert.NoError(t, err)
		assert.Equal(t, struct {
			Foo string `json:"foo"`
		}{}, response)
	})

	t.Run("regression test existence of request header content-type = application/json", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s%s", apiEndpoint, "/foo"),
			func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, req.Header.Get("Content-Type"), "application/json")
				res := httpmock.NewStringResponse(http.StatusOK, `{"bar":"foo"}`)
				res.Header.Add("X-Rate-Limit", "36/200")
				return res, nil
			},
		)

		response := struct {
			Foo string `json:"foo"`
		}{}

		c := NewClient("")
		err := c.(*client).POST("/foo", &response, []byte(`{"bar":"foo"}`))
		assert.NoError(t, err)
		assert.Equal(t, struct {
			Foo string `json:"foo"`
		}{}, response)
	})
}

func TestClient_PUT(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPut, fmt.Sprintf("%s%s", apiEndpoint, "/foo"),
			func(req *http.Request) (*http.Response, error) {
				buf, err := io.ReadAll(req.Body)
				assert.NoError(t, err)
				assert.Equal(t, `{"bar":"foo"}`, string(buf))
				assert.Equal(t, "application/json", req.Header.Get("Accept"))
				assert.Equal(t, "Bearer 6zL9vh8]B9H3BEecwL%Vzh^VwKR3C2CNZ3Bv%=fFxm$z)duY[U+2=3CydZrkQFnA", req.Header.Get("Authorization"))

				res := httpmock.NewStringResponse(http.StatusOK, string(buf))
				res.Header.Add("X-Rate-Limit", "36/200")
				return res, nil
			},
		)

		response := struct {
			Bar string `json:"bar"`
		}{}

		c := NewClient("6zL9vh8]B9H3BEecwL%Vzh^VwKR3C2CNZ3Bv%=fFxm$z)duY[U+2=3CydZrkQFnA")
		err := c.(*client).PUT("/foo", &response, []byte(`{"bar":"foo"}`))
		assert.NoError(t, err)
		assert.Equal(t, "foo", response.Bar)
	})

	t.Run("failure with with expected API error", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPut, fmt.Sprintf("%s%s", apiEndpoint, "/foo"),
			func(req *http.Request) (*http.Response, error) {
				res := httpmock.NewStringResponse(http.StatusBadRequest, `{
  "error": {
    "id": "400",
    "name": "error_name",
    "detail": "Error detail"
  }
}`)
				res.Header.Add("X-Rate-Limit", "36/200")
				return res, nil
			},
		)

		response := struct {
			Foo string `json:"foo"`
		}{}

		c := NewClient("")
		err := c.(*client).PUT("/foo", &response, []byte(`{"bar":"foo"}`))
		expectedErrStr := "api: error id=400 name=error_name detail=Error detail"
		assert.EqualError(t, err, expectedErrStr)
	})

	t.Run("failure with with unexpected API error", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPut, fmt.Sprintf("%s%s", apiEndpoint, "/foo"),
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(http.StatusInternalServerError, "Internal Server Error"), nil
			},
		)

		response := struct {
			Foo string `json:"foo"`
		}{}

		c := NewClient("")
		err := c.(*client).PUT("/foo", &response, []byte(`{"bar":"foo"}`))
		expectedErrStr := "api: error id=500 name=unknown_api_error detail=Unknown API error"
		assert.EqualError(t, err, expectedErrStr)
	})

	t.Run("silent failure due to invalid response model", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPut, fmt.Sprintf("%s%s", apiEndpoint, "/foo"),
			func(req *http.Request) (*http.Response, error) {
				res := httpmock.NewStringResponse(http.StatusOK, `{"bar":"foo"}`)
				res.Header.Add("X-Rate-Limit", "36/200")
				return res, nil
			},
		)

		response := struct {
			Foo string `json:"foo"`
		}{}

		c := NewClient("")
		err := c.(*client).PUT("/foo", &response, []byte(`{"bar":"foo"}`))
		assert.NoError(t, err)
		assert.Equal(t, struct {
			Foo string `json:"foo"`
		}{}, response)
	})

	t.Run("regression test existence of request header content-type = application/json", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPut, fmt.Sprintf("%s%s", apiEndpoint, "/foo"),
			func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, req.Header.Get("Content-Type"), "application/json")
				res := httpmock.NewStringResponse(http.StatusOK, `{"bar":"foo"}`)
				res.Header.Add("X-Rate-Limit", "36/200")
				return res, nil
			},
		)

		response := struct {
			Foo string `json:"foo"`
		}{}

		c := NewClient("")
		err := c.(*client).PUT("/foo", &response, []byte(`{"bar":"foo"}`))
		assert.NoError(t, err)
		assert.Equal(t, struct {
			Foo string `json:"foo"`
		}{}, response)
	})
}

func TestClient_PATCH(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPatch, fmt.Sprintf("%s%s", apiEndpoint, "/foo"),
			func(req *http.Request) (*http.Response, error) {
				buf, err := io.ReadAll(req.Body)
				assert.NoError(t, err)
				assert.Equal(t, `{"bar":"foo"}`, string(buf))
				assert.Equal(t, "application/json", req.Header.Get("Accept"))
				assert.Equal(t, "Bearer 6zL9vh8]B9H3BEecwL%Vzh^VwKR3C2CNZ3Bv%=fFxm$z)duY[U+2=3CydZrkQFnA", req.Header.Get("Authorization"))

				res := httpmock.NewStringResponse(http.StatusOK, string(buf))
				res.Header.Add("X-Rate-Limit", "36/200")
				return res, nil
			},
		)

		response := struct {
			Bar string `json:"bar"`
		}{}

		c := NewClient("6zL9vh8]B9H3BEecwL%Vzh^VwKR3C2CNZ3Bv%=fFxm$z)duY[U+2=3CydZrkQFnA")
		err := c.(*client).PATCH("/foo", &response, []byte(`{"bar":"foo"}`))
		assert.NoError(t, err)
		assert.Equal(t, "foo", response.Bar)
	})

	t.Run("failure with with expected API error", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPatch, fmt.Sprintf("%s%s", apiEndpoint, "/foo"),
			func(req *http.Request) (*http.Response, error) {
				res := httpmock.NewStringResponse(http.StatusBadRequest, `{
  "error": {
    "id": "400",
    "name": "error_name",
    "detail": "Error detail"
  }
}`)
				res.Header.Add("X-Rate-Limit", "36/200")
				return res, nil
			},
		)

		response := struct {
			Foo string `json:"foo"`
		}{}

		c := NewClient("")
		err := c.(*client).PATCH("/foo", &response, []byte(`{"bar":"foo"}`))
		expectedErrStr := "api: error id=400 name=error_name detail=Error detail"
		assert.EqualError(t, err, expectedErrStr)
	})

	t.Run("failure with with unexpected API error", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPatch, fmt.Sprintf("%s%s", apiEndpoint, "/foo"),
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(http.StatusInternalServerError, "Internal Server Error"), nil
			},
		)

		response := struct {
			Foo string `json:"foo"`
		}{}

		c := NewClient("")
		err := c.(*client).PATCH("/foo", &response, []byte(`{"bar":"foo"}`))
		expectedErrStr := "api: error id=500 name=unknown_api_error detail=Unknown API error"
		assert.EqualError(t, err, expectedErrStr)
	})

	t.Run("silent failure due to invalid response model", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPatch, fmt.Sprintf("%s%s", apiEndpoint, "/foo"),
			func(req *http.Request) (*http.Response, error) {
				res := httpmock.NewStringResponse(http.StatusOK, `{"bar":"foo"}`)
				res.Header.Add("X-Rate-Limit", "36/200")
				return res, nil
			},
		)

		response := struct {
			Foo string `json:"foo"`
		}{}

		c := NewClient("")
		err := c.(*client).PATCH("/foo", &response, []byte(`{"bar":"foo"}`))
		assert.NoError(t, err)
		assert.Equal(t, struct {
			Foo string `json:"foo"`
		}{}, response)
	})

	t.Run("regression test existence of request header content-type = application/json", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPatch, fmt.Sprintf("%s%s", apiEndpoint, "/foo"),
			func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
				res := httpmock.NewStringResponse(http.StatusOK, `{"bar":"foo"}`)
				res.Header.Add("X-Rate-Limit", "36/200")
				return res, nil
			},
		)

		response := struct {
			Foo string `json:"foo"`
		}{}

		c := NewClient("")
		err := c.(*client).PATCH("/foo", &response, []byte(`{"bar":"foo"}`))
		assert.NoError(t, err)
		assert.Equal(t, struct {
			Foo string `json:"foo"`
		}{}, response)
	})
}
