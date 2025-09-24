//go:build e2e

package e2e_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/fx"

	"github.com/marcelofabianov/dojo-go/config"
	"github.com/marcelofabianov/dojo-go/internal/di"
	"github.com/marcelofabianov/dojo-go/internal/handler"
)

var (
	testServer *httptest.Server
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:16-alpine"),
		postgres.WithDatabase("test-db-e2e"),
		postgres.WithUsername("user-e2e"),
		postgres.WithPassword("password-e2e"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	if err != nil {
		log.Fatalf("could not start postgres container: %s", err)
	}

	defer func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			log.Fatalf("could not stop postgres container: %s", err)
		}
	}()

	host, _ := pgContainer.Host(ctx)
	port, _ := pgContainer.MappedPort(ctx, "5432")

	testDBConfig := &config.DBConfig{
		Host:            host,
		Port:            port.Int(),
		User:            "user-e2e",
		Password:        "password-e2e",
		Name:            "test-db-e2e",
		SSLMode:         "disable",
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: 5 * time.Minute,
		ConnMaxIdleTime: 1 * time.Minute,
		QueryTimeout:    5 * time.Second,
		ExecTimeout:     3 * time.Second,
	}

	connStr, _ := pgContainer.ConnectionString(ctx, "sslmode=disable")
	tempDB, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		log.Fatalf("could not connect to temp db: %s", err)
	}
	goose.SetBaseFS(os.DirFS("../../"))
	if err := goose.Up(tempDB.DB, "db/migrations"); err != nil {
		log.Fatalf("could not run migrations: %s", err)
	}
	tempDB.Close()

	var router *chi.Mux
	app := fx.New(
		di.Config,
		di.Pkg,
		di.Repository,
		di.Service,
		di.Handler,
		fx.Replace(testDBConfig),
		fx.Populate(&router),
	)

	if err := app.Start(ctx); err != nil {
		log.Fatalf("failed to start fx app: %s", err)
	}

	testServer = httptest.NewServer(router)

	defer func() {
		testServer.Close()
		if err := app.Stop(ctx); err != nil {
			log.Printf("failed to stop fx app: %s", err)
		}
	}()

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCourseAPI_E2E(t *testing.T) {
	require.NotNil(t, testServer, "test server should not be nil")
	client := testServer.Client()
	var createdCourseID string

	t.Run("should create a course", func(t *testing.T) {
		courseInput := `{"title": "E2E Testing", "description": "How to test everything."}`
		reqBody := bytes.NewBufferString(courseInput)

		resp, err := client.Post(fmt.Sprintf("%s/api/v1/courses", testServer.URL), "application/json", reqBody)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusCreated, resp.StatusCode)

		var courseResponse handler.CreateCourseResponse
		err = json.NewDecoder(resp.Body).Decode(&courseResponse)
		require.NoError(t, err)
		require.NotEmpty(t, courseResponse.ID)
		require.Equal(t, "E2E Testing", courseResponse.Title)

		createdCourseID = courseResponse.ID
	})

	t.Run("should get the created course", func(t *testing.T) {
		require.NotEmpty(t, createdCourseID, "created course ID should not be empty")

		resp, err := client.Get(fmt.Sprintf("%s/api/v1/courses/%s", testServer.URL, createdCourseID))
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode)

		var courseResponse handler.CreateCourseResponse
		err = json.NewDecoder(resp.Body).Decode(&courseResponse)
		require.NoError(t, err)
		require.Equal(t, createdCourseID, courseResponse.ID)
		require.Equal(t, "E2E Testing", courseResponse.Title)
	})

	t.Run("should update the course", func(t *testing.T) {
		require.NotEmpty(t, createdCourseID, "created course ID should not be empty")
		updateInput := `{"title": "Advanced E2E Testing", "description": "Updated description."}`
		reqBody := bytes.NewBufferString(updateInput)

		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/v1/courses/%s", testServer.URL, createdCourseID), reqBody)
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode)

		var courseResponse handler.CreateCourseResponse
		err = json.NewDecoder(resp.Body).Decode(&courseResponse)
		require.NoError(t, err)
		require.Equal(t, "Advanced E2E Testing", courseResponse.Title)
	})

	t.Run("should delete the course", func(t *testing.T) {
		require.NotEmpty(t, createdCourseID, "created course ID should not be empty")

		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/v1/courses/%s", testServer.URL, createdCourseID), nil)
		require.NoError(t, err)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusNoContent, resp.StatusCode)
	})

	t.Run("should not find the deleted course", func(t *testing.T) {
		require.NotEmpty(t, createdCourseID, "created course ID should not be empty")

		resp, err := client.Get(fmt.Sprintf("%s/api/v1/courses/%s", testServer.URL, createdCourseID))
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}
