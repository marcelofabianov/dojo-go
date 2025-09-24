//go:build integration

package repository

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/marcelofabianov/dojo-go/internal/model"
)

var (
	db *sqlx.DB
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	// 1. Iniciar o contêiner do PostgreSQL
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:16-alpine"),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	if err != nil {
		log.Fatalf("could not start postgres container: %s", err)
	}

	// Garante que o contêiner será encerrado após os testes
	defer func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			log.Fatalf("could not stop postgres container: %s", err)
		}
	}()

	// 2. Obter a connection string e conectar
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatalf("could not get connection string: %s", err)
	}

	db, err = sqlx.Connect("pgx", connStr)
	if err != nil {
		log.Fatalf("could not connect to database: %s", err)
	}

	// 3. Rodar as migrations
	goose.SetBaseFS(os.DirFS("../../"))
	if err := goose.Up(db.DB, "db/migrations"); err != nil {
		log.Fatalf("could not run migrations: %s", err)
	}

	// 4. Executar os testes
	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestCourseRepository_Integration(t *testing.T) {
	require.NotNil(t, db, "database connection should not be nil")

	repo := NewPostgresCourseRepository(db)
	ctx := context.Background()

	newCourse, err := model.NewCourse(model.NewCourseInput{
		Title:       "Integration Testing 101",
		Description: "A deep dive into integration testing.",
	})
	require.NoError(t, err)

	t.Run("Create", func(t *testing.T) {
		err := repo.CreateCourse(ctx, newCourse)
		require.NoError(t, err)
	})

	t.Run("Get", func(t *testing.T) {
		retrievedCourse, err := repo.GetCourseByID(ctx, newCourse.ID)
		require.NoError(t, err)
		require.NotNil(t, retrievedCourse)
		require.Equal(t, newCourse.Title, retrievedCourse.Title)
		require.Equal(t, newCourse.Description, retrievedCourse.Description)
	})

	t.Run("Update", func(t *testing.T) {
		newCourse.Title = "Advanced Integration Testing"
		err := repo.UpdateCourse(ctx, newCourse)
		require.NoError(t, err)

		updatedCourse, err := repo.GetCourseByID(ctx, newCourse.ID)
		require.NoError(t, err)
		require.Equal(t, "Advanced Integration Testing", updatedCourse.Title)
	})

	t.Run("Delete", func(t *testing.T) {
		err := repo.DeleteCourseByID(ctx, newCourse.ID)
		require.NoError(t, err)
	})

	t.Run("Verify Deletion", func(t *testing.T) {
		deletedCourse, err := repo.GetCourseByID(ctx, newCourse.ID)
		require.Error(t, err)
		require.Nil(t, deletedCourse)
		require.ErrorIs(t, err, model.ErrCourseNotFound)
	})
}
