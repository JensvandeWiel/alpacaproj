package extras

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/project"
	"os"
	"os/exec"
	"path"
	"text/template"
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

	//Add extra test helpers
	err = buildMockUsersStore(prj)
	if err != nil {
		return err
	}

	if prj.Extras.HasExtra(project.SQLC) {
		command := exec.Command("sqlc", "generate")
		command.Dir = prj.Path
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		err = command.Run()
		if err != nil {
			return err
		}
	}

	command := exec.Command("go", "mod", "tidy")
	command.Dir = prj.Path
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err = command.Run()
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

	tmpl, err := template.New("frontendAuthService").Parse(frontendAuthServiceTemplate)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "services", "frontend_auth_service.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	modelModule := "models"
	if prj.Extras.HasExtra(project.SQLC) {
		modelModule = "repository"
	}

	data := map[string]interface{}{
		"packageName": prj.PackageName,
		"modelModule": modelModule,
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}

	tmpl, err = template.New("frontendAuthServiceTest").Parse(frontendAuthServiceTestTemplate)
	if err != nil {
		return err
	}

	file, err = os.OpenFile(path.Join(prj.Path, "services", "frontend_auth_service_test.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	err = tmpl.Execute(file, data)
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

	//todo add store for project without sqlc
	if prj.Extras.HasExtra(project.SQLC) {
		templateString := sqlcPostgresUserStoreTemplate
		if prj.Database == project.MySQL {
			templateString = sqlcMysqlUserStoreTemplate
		}

		tmpl, err := template.New("userStore").Parse(templateString)

		if err != nil {
			return err
		}

		file, err := os.OpenFile(path.Join(prj.Path, "stores", "user_store.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return err
		}

		err = tmpl.Execute(file, data)
		if err != nil {
			return err
		}

		file.Close()
	} else {
		templateString := mysqlUserStoreTemplate
		if prj.Database == project.Postgres {
			templateString = postgresUserStoreTemplate
		}

		tmpl, err := template.New("userStore").Parse(templateString)
		if err != nil {
			return err
		}

		file, err := os.OpenFile(path.Join(prj.Path, "stores", "user_store.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return err
		}

		err = tmpl.Execute(file, data)
		if err != nil {
			return err
		}

		file.Close()
	}

	templateString := postgresUserStoreTestTemplate
	if prj.Database == project.MySQL {
		templateString = mysqlUserStoreTestTemplate
	}

	tmpl, err := template.New("userStoreTest").Parse(templateString)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "stores", "user_store_test.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}

	file.Close()

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

		tmpl, err := template.New("userQuery").Parse(templateString)
		if err != nil {
			return err
		}

		file, err := os.OpenFile(path.Join(prj.Path, "queries", "user.sql"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return err
		}

		err = tmpl.Execute(file, nil)
		if err != nil {
			return err
		}

		file.Close()
	} else {
		tmpl, err := template.New("userModel").Parse(userModelTemplate)
		if err != nil {
			return err
		}

		file, err := os.OpenFile(path.Join(prj.Path, "models", "user.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return err
		}

		err = tmpl.Execute(file, nil)
		if err != nil {
			return err
		}

		file.Close()
	}

	databaseTemplate := createUsersTableTemplatePostgres
	if prj.Database == project.MySQL {
		databaseTemplate = createUsersTableTemplateMysql
	}

	tmpl, err := template.New("createUsersTableMysql").Parse(databaseTemplate)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "migrations", project.GenerateTimestamp()+"_create_users_table.sql"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, nil)

	return nil
}

//go:embed templates/frontendauth/login_request_template.go.tmpl
var loginRequestTemplate string

func buildLoginRequest(prj *project.Project) error {
	prj.Logger.Debug("Building extra LoginRequest")

	tmpl, err := template.New("loginRequest").Parse(loginRequestTemplate)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "requests", "login_request.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	err = tmpl.Execute(file, nil)
	if err != nil {
		return err
	}

	return nil
}

//go:embed templates/frontendauth/frontend_auth_middleware_template.go.tmpl
var middlewareTemplate string

func buildMiddleware(prj *project.Project) error {
	prj.Logger.Debug("Building extra Middleware")

	tmpl, err := template.New("middleware").Parse(middlewareTemplate)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "middleware", "frontend_auth_middleware.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	err = tmpl.Execute(file, nil)

	return nil
}

//go:embed templates/frontendauth/mock_usersstore_template.go.tmpl
var mock_usersstore_template string

func buildMockUsersStore(prj *project.Project) error {
	prj.Logger.Debug("Generating test_helpers/mock_usersstore.go")

	tmpl, err := template.New("mock_usersstore").Parse(mock_usersstore_template)

	file, err := os.OpenFile(path.Join(prj.Path, "test_helpers", "mock_user_store.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	modelName := "models"
	if prj.Extras.HasExtra(project.SQLC) {
		modelName = "repository"
	}

	data := map[string]interface{}{
		"packageName": prj.PackageName,
		"modelModule": modelName,
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}
	prj.Logger.Debug("Generated test_helpers/mock_usersstore.go")
	return nil
}
