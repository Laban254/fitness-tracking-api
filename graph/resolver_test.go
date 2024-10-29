package graph

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/handler/extension"
    "github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
    srv := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{}}))
    srv.Use(extension.Introspection{})

    // Create an HTTP request to simulate a mutation
    reqBody := map[string]interface{}{
        "query": `mutation { createUser(email: "testuser@example.com", password: "password") { id email } }`,
    }
    body, err := json.Marshal(reqBody)
    if err != nil {
        t.Fatalf("failed to marshal request body: %v", err)
    }

    // Use httptest to create a ResponseRecorder
    req := httptest.NewRequest("POST", "/query", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    
    // Record the response
    rr := httptest.NewRecorder()
    srv.ServeHTTP(rr, req)

    // Assert the response status code
    assert.Equal(t, http.StatusOK, rr.Code)

    // Assert the response body
    var resp struct {
        Data struct {
            CreateUser struct {
                ID    string `json:"id"`
                Email string `json:"email"`
            } `json:"createUser"`
        } `json:"data"`
    }

    if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
        t.Fatalf("failed to unmarshal response body: %v", err)
    }

    // Validate the response values
    assert.Equal(t, "expected_user_id", resp.Data.CreateUser.ID)
    assert.Equal(t, "testuser@example.com", resp.Data.CreateUser.Email)
}
