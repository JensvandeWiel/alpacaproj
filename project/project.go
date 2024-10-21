package project

import (
	"errors"
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
	"path/filepath"
)

type ExtraOptions []ExtraOption

func (e ExtraOptions) HasExtra(extra ExtraOption) bool {
	for _, opt := range e {
		if opt == extra {
			return true
		}
	}
	return false
}

type Project struct {
	Logger       *slog.Logger   `yaml:"-"`
	Name         string         `yaml:"name"`
	PackageName  string         `yaml:"package_name"`
	Path         string         `yaml:"-"`
	Database     DatabaseDriver `yaml:"database"`
	HasFrontend  bool           `yaml:"has_frontend"`
	FrontendType FrontendType   `yaml:"frontend_type"`
	Extras       ExtraOptions   `yaml:"extras"`
}

type ExtraOption string
type DatabaseDriver string
type FrontendType string

func (f FrontendType) HasFrontend() bool {
	return f != None
}

func (f FrontendType) IsInertia() bool {
	return f == InertiaReact || f == InertiaVue || f == InertiaSvelte
}

const (
	MySQL    DatabaseDriver = "mysql"
	Postgres DatabaseDriver = "postgres"

	None          FrontendType = "none"
	InertiaReact  FrontendType = "inertia+react"
	InertiaVue    FrontendType = "inertia+vue"
	InertiaSvelte FrontendType = "inertia+svelte"

	Svelte5      ExtraOption = "svelte5"
	SQLC         ExtraOption = "sqlc"
	FrontendAuth ExtraOption = "frontend-auth"
)

var (
	ErrorInvalidDatabase = errors.New("invalid/unsupported database driver")
	ErrorInvalidFrontend = errors.New("invalid/unsupported frontend type")
)

// ParseProjectName parses the project name and returns the absolute path
func ParseProjectName(projName string) (name string, absPath string, err error) {
	if projName == "." {
		absPath, _ := os.Getwd()
		return filepath.Base(absPath), absPath, nil
	}
	if filepath.IsAbs(projName) {
		return filepath.Base(projName), projName, nil
	}
	absPath, err = filepath.Abs(projName)
	return filepath.Base(absPath), absPath, err
}

func parseDatabase(database string) (DatabaseDriver, error) {
	switch database {
	case "mysql":
		return MySQL, nil
	case "postgres":
		return Postgres, nil
	default:
		return "", ErrorInvalidDatabase
	}
}

func parseFrontend(frontend string) (FrontendType, error) {
	switch frontend {
	case "none":
		return None, nil
	case "inertia+react":
		return InertiaReact, nil
	case "inertia+vue":
		return InertiaVue, nil
	case "inertia+svelte":
		return InertiaSvelte, nil
	default:
		return None, ErrorInvalidFrontend
	}
}

var ErrInvalidExtra = errors.New("invalid extra option")

func parseExtra(extra string) (ExtraOption, error) {
	switch extra {
	case "svelte5":
		return Svelte5, nil
	case "sqlc":
		return SQLC, nil
	case "frontend-auth":
		return FrontendAuth, nil
	default:
		return "", ErrInvalidExtra
	}
}

func parseExtras(extras []string) ([]ExtraOption, error) {
	var opts []ExtraOption
	for _, extra := range extras {
		opt, err := parseExtra(extra)
		if err != nil {
			return nil, err
		}
		opts = append(opts, opt)
	}
	return opts, nil
}

func NewProject(projName string, database string, frontend string, packageName string, extras []string, logger *slog.Logger) (*Project, error) {
	name, absPath, err := ParseProjectName(projName)
	if err != nil {
		return nil, err
	}

	dbDriver, err := parseDatabase(database)
	if err != nil {
		return nil, err
	}

	frontendType, err := parseFrontend(frontend)
	if err != nil {
		return nil, err
	}

	extraOpts, err := parseExtras(extras)
	if err != nil {
		return nil, err
	}

	if packageName == "" {
		packageName = name
	}

	logger.Debug("Parsed project", slog.String("name", name), slog.String("path", absPath), slog.String("database", string(dbDriver)), slog.String("frontend", string(frontendType)))

	return &Project{
		Logger:       logger,
		PackageName:  packageName,
		Name:         name,
		Path:         absPath,
		Database:     dbDriver,
		HasFrontend:  frontendType.HasFrontend(),
		FrontendType: frontendType,
		Extras:       extraOpts,
	}, nil
}

func (p *Project) SaveConfig() error {
	file, err := os.OpenFile(filepath.Join(p.Path, "alpaca.yaml"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := yaml.NewEncoder(file)
	return enc.Encode(p)
}

func LoadProject(path string, isVerbose bool) (*Project, error) {
	lvl := slog.LevelInfo
	if isVerbose {
		lvl = slog.LevelDebug
	}

	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: lvl,
	}))

	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(absPath, "alpaca.yaml"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dec := yaml.NewDecoder(file)
	var proj Project
	err = dec.Decode(&proj)
	if err != nil {
		return nil, err
	}

	// Set the Path field to the absolute path
	proj.Path = absPath

	// Set the logger
	proj.Logger = l

	return &proj, nil
}
