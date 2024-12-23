package extras

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"path"
)

func BuildFrontendAuth(prj *project.Project) error {
	prj.Logger.Debug("Building extra FrontendAuth")

	err := buildFrontendAuthService(prj)
	if err != nil {
		return err
	}

	// Add stores
	err = buildUserStore(prj)
	if err != nil {
		return err
	}

	// Add requests
	err = buildLoginRequest(prj)
	if err != nil {
		return err
	}

	// Patch middleware
	err = buildMiddleware(prj)
	if err != nil {
		return err
	}

	// Add user model
	err = buildUserModel(prj)
	if err != nil {
		return err
	}

	// Add extra test helpers
	err = buildMockUsersStore(prj)
	if err != nil {
		return err
	}

	if prj.Extras.HasExtra(project.SQLC) {
		err = helpers.GenerateSQLCDefinitions(prj)
		if err != nil {
			return err
		}
	}

	err = helpers.RunGoTidy(prj)
	if err != nil {
		return err
	}

	return nil
}

//go:embed templates/frontendauth/frontend_auth_service_template.go.tmpl
var frontendAuthServiceTemplate string

//go:embed templates/frontendauth/frontend_auth_service_test_template.go.tmpl
var frontendAuthServiceTestTemplate string

func buildFrontendAuthService(prj *project.Project) error {
	prj.Logger.Debug("Building extra FrontendAuthService")

	modelModule := "models"
	if prj.Extras.HasExtra(project.SQLC) {
		modelModule = "repository"
	}

	data := map[string]interface{}{
		"packageName": prj.PackageName,
		"modelModule": modelModule,
	}

	err := helpers.WriteTemplateToFile(prj, path.Join("services", "frontend_auth_service.go"), frontendAuthServiceTemplate, data)
	if err != nil {
		return err
	}

	err = helpers.WriteTemplateToFile(prj, path.Join("services", "frontend_auth_service_test.go"), frontendAuthServiceTestTemplate, data)
	if err != nil {
		return err
	}

	return nil
}

//go:embed templates/frontendauth/sqlc_postgres_user_store_template.go.tmpl
var sqlcPostgresUserStoreTemplate string

//go:embed templates/frontendauth/sqlc_mysql_user_store_template.go.tmpl
var sqlcMysqlUserStoreTemplate string

//go:embed templates/frontendauth/mysql_user_store_template.go.tmpl
var mysqlUserStoreTemplate string

//go:embed templates/frontendauth/postgres_user_store_template.go.tmpl
var postgresUserStoreTemplate string

//go:embed templates/frontendauth/mysql_user_store_test_template.go.tmpl
var mysqlUserStoreTestTemplate string

//go:embed templates/frontendauth/postgres_user_store_test_template.go.tmpl
var postgresUserStoreTestTemplate string

func buildUserStore(prj *project.Project) error {
	prj.Logger.Debug("Building extra User")

	modelName := "models"
	if prj.Extras.HasExtra(project.SQLC) {
		modelName = "repository"
	}

	data := map[string]interface{}{
		"packageName": prj.PackageName,
		"modelModule": modelName,
	}

	// todo add store for project without sqlc
	if prj.Extras.HasExtra(project.SQLC) {
		templateString := sqlcPostgresUserStoreTemplate
		if prj.Database == project.MySQL {
			templateString = sqlcMysqlUserStoreTemplate
		}

		err := helpers.WriteTemplateToFile(prj, path.Join("stores", "user_store.go"), templateString, data)
		if err != nil {
			return err
		}
	} else {
		templateString := mysqlUserStoreTemplate
		if prj.Database == project.Postgres {
			templateString = postgresUserStoreTemplate
		}

		err := helpers.WriteTemplateToFile(prj, path.Join("stores", "user_store.go"), templateString, data)
		if err != nil {
			return err
		}
	}

	templateString := postgresUserStoreTestTemplate
	if prj.Database == project.MySQL {
		templateString = mysqlUserStoreTestTemplate
	}

	err := helpers.WriteTemplateToFile(prj, path.Join("stores", "user_store_test.go"), templateString, data)
	if err != nil {
		return err
	}

	return nil
}

//go:embed templates/frontendauth/mysql_user_query_template.sql.tmpl
var mysqlUserQueryTemplate string

//go:embed templates/frontendauth/postgres_user_query_template.sql.tmpl
var postgresUserQueryTemplate string

//go:embed templates/frontendauth/create_users_table_template_mysql.sql.tmpl
var createUsersTableTemplateMysql string

//go:embed templates/frontendauth/create_users_table_template_postgres.sql.tmpl
var createUsersTableTemplatePostgres string

//go:embed templates/frontendauth/user_model.go.tmpl
var userModelTemplate string

func buildUserModel(prj *project.Project) error {
	prj.Logger.Debug("Building extra User")

	// Add user model
	if prj.Extras.HasExtra(project.SQLC) {
		templateString := mysqlUserQueryTemplate
		if prj.Database == project.Postgres {
			templateString = postgresUserQueryTemplate
		}

		err := helpers.WriteTemplateToFile(prj, path.Join("queries", "user.sql"), templateString, nil)
		if err != nil {
			return err
		}
	} else {
		err := helpers.WriteTemplateToFile(prj, path.Join("models", "user.go"), userModelTemplate, nil)
		if err != nil {
			return err
		}
	}

	databaseTemplate := createUsersTableTemplatePostgres
	if prj.Database == project.MySQL {
		databaseTemplate = createUsersTableTemplateMysql
	}

	err := helpers.WriteTemplateToFile(prj, path.Join("migrations", project.GenerateTimestamp()+"_create_users_table.sql"), databaseTemplate, nil)
	if err != nil {
		return err
	}

	return nil
}

//go:embed templates/frontendauth/login_request_template.go.tmpl
var loginRequestTemplate string

func buildLoginRequest(prj *project.Project) error {
	prj.Logger.Debug("Building extra LoginRequest")

	err := helpers.WriteTemplateToFile(prj, path.Join("requests", "login_request.go"), loginRequestTemplate, nil)
	if err != nil {
		return err
	}

	return nil
}

//go:embed templates/frontendauth/frontend_auth_middleware_template.go.tmpl
var middlewareTemplate string

func buildMiddleware(prj *project.Project) error {
	prj.Logger.Debug("Building extra Middleware")

	err := helpers.WriteTemplateToFile(prj, path.Join("middleware", "frontend_auth_middleware.go"), middlewareTemplate, nil)
	if err != nil {
		return err
	}

	return nil
}

//go:embed templates/frontendauth/mock_usersstore_template.go.tmpl
var mock_usersstore_template string

func buildMockUsersStore(prj *project.Project) error {
	prj.Logger.Debug("Generating test_helpers/mock_usersstore.go")

	modelName := "models"
	if prj.Extras.HasExtra(project.SQLC) {
		modelName = "repository"
	}

	data := map[string]interface{}{
		"packageName": prj.PackageName,
		"modelModule": modelName,
	}

	err := helpers.WriteTemplateToFile(prj, path.Join("test_helpers", "mock_user_store.go"), mock_usersstore_template, data)
	if err != nil {
		return err
	}

	prj.Logger.Debug("Generated test_helpers/mock_usersstore.go")
	return nil
}
