// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/piperenv"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type fortifyExecuteScanOptions struct {
	AuthToken                       string   `json:"authToken,omitempty"`
	CustomScanVersion               string   `json:"customScanVersion,omitempty"`
	GithubToken                     string   `json:"githubToken,omitempty"`
	AutoCreate                      bool     `json:"autoCreate,omitempty"`
	ModulePath                      string   `json:"modulePath,omitempty"`
	PythonRequirementsFile          string   `json:"pythonRequirementsFile,omitempty"`
	AutodetectClasspath             bool     `json:"autodetectClasspath,omitempty"`
	MustAuditIssueGroups            string   `json:"mustAuditIssueGroups,omitempty"`
	SpotAuditIssueGroups            string   `json:"spotAuditIssueGroups,omitempty"`
	PythonRequirementsInstallSuffix string   `json:"pythonRequirementsInstallSuffix,omitempty"`
	PythonVersion                   string   `json:"pythonVersion,omitempty"`
	UploadResults                   bool     `json:"uploadResults,omitempty"`
	Version                         string   `json:"version,omitempty"`
	BuildDescriptorFile             string   `json:"buildDescriptorFile,omitempty"`
	CommitID                        string   `json:"commitId,omitempty"`
	CommitMessage                   string   `json:"commitMessage,omitempty"`
	GithubAPIURL                    string   `json:"githubApiUrl,omitempty"`
	Owner                           string   `json:"owner,omitempty"`
	Repository                      string   `json:"repository,omitempty"`
	Memory                          string   `json:"memory,omitempty"`
	UpdateRulePack                  bool     `json:"updateRulePack,omitempty"`
	ReportDownloadEndpoint          string   `json:"reportDownloadEndpoint,omitempty"`
	PollingMinutes                  int      `json:"pollingMinutes,omitempty"`
	QuickScan                       bool     `json:"quickScan,omitempty"`
	Translate                       string   `json:"translate,omitempty"`
	Src                             []string `json:"src,omitempty"`
	Exclude                         []string `json:"exclude,omitempty"`
	APIEndpoint                     string   `json:"apiEndpoint,omitempty"`
	ReportType                      string   `json:"reportType,omitempty"`
	PythonAdditionalPath            []string `json:"pythonAdditionalPath,omitempty"`
	ArtifactURL                     string   `json:"artifactUrl,omitempty"`
	ConsiderSuspicious              bool     `json:"considerSuspicious,omitempty"`
	FprUploadEndpoint               string   `json:"fprUploadEndpoint,omitempty"`
	ProjectName                     string   `json:"projectName,omitempty"`
	Reporting                       bool     `json:"reporting,omitempty"`
	ServerURL                       string   `json:"serverUrl,omitempty"`
	PullRequestMessageRegexGroup    int      `json:"pullRequestMessageRegexGroup,omitempty"`
	DeltaMinutes                    int      `json:"deltaMinutes,omitempty"`
	SpotCheckMinimum                int      `json:"spotCheckMinimum,omitempty"`
	FprDownloadEndpoint             string   `json:"fprDownloadEndpoint,omitempty"`
	VersioningModel                 string   `json:"versioningModel,omitempty"`
	PythonInstallCommand            string   `json:"pythonInstallCommand,omitempty"`
	ReportTemplateID                int      `json:"reportTemplateId,omitempty"`
	FilterSetTitle                  string   `json:"filterSetTitle,omitempty"`
	PullRequestName                 string   `json:"pullRequestName,omitempty"`
	PullRequestMessageRegex         string   `json:"pullRequestMessageRegex,omitempty"`
	BuildTool                       string   `json:"buildTool,omitempty"`
	ProjectSettingsFile             string   `json:"projectSettingsFile,omitempty"`
	GlobalSettingsFile              string   `json:"globalSettingsFile,omitempty"`
	M2Path                          string   `json:"m2Path,omitempty"`
	VerifyOnly                      bool     `json:"verifyOnly,omitempty"`
	InstallArtifacts                bool     `json:"installArtifacts,omitempty"`
}

type fortifyExecuteScanInflux struct {
	fortify_data struct {
		fields struct {
			projectName       string
			projectVersion    string
			violations        int
			corporateTotal    int
			corporateAudited  int
			auditAllTotal     int
			auditAllAudited   int
			spotChecksTotal   int
			spotChecksAudited int
			spotChecksGap     int
			suspicious        int
			exploitable       int
			suppressed        int
		}
		tags struct {
		}
	}
}

func (i *fortifyExecuteScanInflux) persist(path, resourceName string) {
	measurementContent := []struct {
		measurement string
		valType     string
		name        string
		value       interface{}
	}{
		{valType: config.InfluxField, measurement: "fortify_data", name: "projectName", value: i.fortify_data.fields.projectName},
		{valType: config.InfluxField, measurement: "fortify_data", name: "projectVersion", value: i.fortify_data.fields.projectVersion},
		{valType: config.InfluxField, measurement: "fortify_data", name: "violations", value: i.fortify_data.fields.violations},
		{valType: config.InfluxField, measurement: "fortify_data", name: "corporateTotal", value: i.fortify_data.fields.corporateTotal},
		{valType: config.InfluxField, measurement: "fortify_data", name: "corporateAudited", value: i.fortify_data.fields.corporateAudited},
		{valType: config.InfluxField, measurement: "fortify_data", name: "auditAllTotal", value: i.fortify_data.fields.auditAllTotal},
		{valType: config.InfluxField, measurement: "fortify_data", name: "auditAllAudited", value: i.fortify_data.fields.auditAllAudited},
		{valType: config.InfluxField, measurement: "fortify_data", name: "spotChecksTotal", value: i.fortify_data.fields.spotChecksTotal},
		{valType: config.InfluxField, measurement: "fortify_data", name: "spotChecksAudited", value: i.fortify_data.fields.spotChecksAudited},
		{valType: config.InfluxField, measurement: "fortify_data", name: "spotChecksGap", value: i.fortify_data.fields.spotChecksGap},
		{valType: config.InfluxField, measurement: "fortify_data", name: "suspicious", value: i.fortify_data.fields.suspicious},
		{valType: config.InfluxField, measurement: "fortify_data", name: "exploitable", value: i.fortify_data.fields.exploitable},
		{valType: config.InfluxField, measurement: "fortify_data", name: "suppressed", value: i.fortify_data.fields.suppressed},
	}

	errCount := 0
	for _, metric := range measurementContent {
		err := piperenv.SetResourceParameter(path, resourceName, filepath.Join(metric.measurement, fmt.Sprintf("%vs", metric.valType), metric.name), metric.value)
		if err != nil {
			log.Entry().WithError(err).Error("Error persisting influx environment.")
			errCount++
		}
	}
	if errCount > 0 {
		log.Entry().Fatal("failed to persist Influx environment")
	}
}

// FortifyExecuteScanCommand This step executes a Fortify scan on the specified project to perform static code analysis and check the source code for security flaws.
func FortifyExecuteScanCommand() *cobra.Command {
	const STEP_NAME = "fortifyExecuteScan"

	metadata := fortifyExecuteScanMetadata()
	var stepConfig fortifyExecuteScanOptions
	var startTime time.Time
	var influx fortifyExecuteScanInflux

	var createFortifyExecuteScanCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "This step executes a Fortify scan on the specified project to perform static code analysis and check the source code for security flaws.",
		Long: `This step executes a Fortify scan on the specified project to perform static code analysis and check the source code for security flaws.

The Fortify step triggers a scan locally on your Jenkins within a docker container so finally you have to supply a docker image with a Fortify SCA
and Java plus Maven or alternatively Python installed into it for being able to perform any scans.`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			startTime = time.Now()
			log.SetStepName(STEP_NAME)
			log.SetVerbose(GeneralConfig.Verbose)

			path, _ := os.Getwd()
			fatalHook := &log.FatalHook{CorrelationID: GeneralConfig.CorrelationID, Path: path}
			log.RegisterHook(fatalHook)

			err := PrepareConfig(cmd, &metadata, STEP_NAME, &stepConfig, config.OpenPiperFile)
			if err != nil {
				log.SetErrorCategory(log.ErrorConfiguration)
				return err
			}
			log.RegisterSecret(stepConfig.AuthToken)
			log.RegisterSecret(stepConfig.GithubToken)

			if len(GeneralConfig.HookConfig.SentryConfig.Dsn) > 0 {
				sentryHook := log.NewSentryHook(GeneralConfig.HookConfig.SentryConfig.Dsn, GeneralConfig.CorrelationID)
				log.RegisterHook(&sentryHook)
			}

			return nil
		},
		Run: func(_ *cobra.Command, _ []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				config.RemoveVaultSecretFiles()
				influx.persist(GeneralConfig.EnvRootPath, "influx")
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetryData.ErrorCategory = log.GetErrorCategory().String()
				telemetry.Send(&telemetryData)
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, STEP_NAME)
			fortifyExecuteScan(stepConfig, &telemetryData, &influx)
			telemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addFortifyExecuteScanFlags(createFortifyExecuteScanCmd, &stepConfig)
	return createFortifyExecuteScanCmd
}

func addFortifyExecuteScanFlags(cmd *cobra.Command, stepConfig *fortifyExecuteScanOptions) {
	cmd.Flags().StringVar(&stepConfig.AuthToken, "authToken", os.Getenv("PIPER_authToken"), "The FortifyToken to use for authentication")
	cmd.Flags().StringVar(&stepConfig.CustomScanVersion, "customScanVersion", os.Getenv("PIPER_customScanVersion"), "Custom version of the Fortify project used as source.")
	cmd.Flags().StringVar(&stepConfig.GithubToken, "githubToken", os.Getenv("PIPER_githubToken"), "GitHub personal access token as per https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line")
	cmd.Flags().BoolVar(&stepConfig.AutoCreate, "autoCreate", false, "Whether Fortify project and project version shall be implicitly auto created in case they cannot be found in the backend")
	cmd.Flags().StringVar(&stepConfig.ModulePath, "modulePath", `./`, "Allows providing the path for the module to scan")
	cmd.Flags().StringVar(&stepConfig.PythonRequirementsFile, "pythonRequirementsFile", os.Getenv("PIPER_pythonRequirementsFile"), "The requirements file used in `buildTool: 'pip'` to populate the build environment with the necessary dependencies")
	cmd.Flags().BoolVar(&stepConfig.AutodetectClasspath, "autodetectClasspath", true, "Whether the classpath is automatically determined via build tool i.e. maven or pip or not at all")
	cmd.Flags().StringVar(&stepConfig.MustAuditIssueGroups, "mustAuditIssueGroups", `Corporate Security Requirements, Audit All`, "Comma separated list of issue groups that must be audited completely")
	cmd.Flags().StringVar(&stepConfig.SpotAuditIssueGroups, "spotAuditIssueGroups", `Spot Checks of Each Category`, "Comma separated list of issue groups that are spot checked and for which `spotCheckMinimum` audited issues are enforced")
	cmd.Flags().StringVar(&stepConfig.PythonRequirementsInstallSuffix, "pythonRequirementsInstallSuffix", os.Getenv("PIPER_pythonRequirementsInstallSuffix"), "The suffix for the command used to install the requirements file in `buildTool: 'pip'` to populate the build environment with the necessary dependencies")
	cmd.Flags().StringVar(&stepConfig.PythonVersion, "pythonVersion", `python3`, "Python version to be used in `buildTool: 'pip'`")
	cmd.Flags().BoolVar(&stepConfig.UploadResults, "uploadResults", true, "Whether results shall be uploaded or not")
	cmd.Flags().StringVar(&stepConfig.Version, "version", os.Getenv("PIPER_version"), "Version used in conjunction with [`versioningModel`](#versioningModel) to identify the Fortify project to be created and used for results aggregation.")
	cmd.Flags().StringVar(&stepConfig.BuildDescriptorFile, "buildDescriptorFile", `./pom.xml`, "Path to the build descriptor file addressing the module/folder to be scanned.")
	cmd.Flags().StringVar(&stepConfig.CommitID, "commitId", os.Getenv("PIPER_commitId"), "Set the Git commit ID for identifying artifacts throughout the scan.")
	cmd.Flags().StringVar(&stepConfig.CommitMessage, "commitMessage", os.Getenv("PIPER_commitMessage"), "Set the Git commit message for identifying pull request merges throughout the scan.")
	cmd.Flags().StringVar(&stepConfig.GithubAPIURL, "githubApiUrl", `https://api.github.com`, "Set the GitHub API URL.")
	cmd.Flags().StringVar(&stepConfig.Owner, "owner", os.Getenv("PIPER_owner"), "Set the GitHub organization.")
	cmd.Flags().StringVar(&stepConfig.Repository, "repository", os.Getenv("PIPER_repository"), "Set the GitHub repository.")
	cmd.Flags().StringVar(&stepConfig.Memory, "memory", `-Xmx4G -Xms512M`, "The amount of memory granted to the translate/scan executions")
	cmd.Flags().BoolVar(&stepConfig.UpdateRulePack, "updateRulePack", true, "Whether the rule pack shall be updated and pulled from Fortify SSC before scanning or not")
	cmd.Flags().StringVar(&stepConfig.ReportDownloadEndpoint, "reportDownloadEndpoint", `/transfer/reportDownload.html`, "Fortify SSC endpoint for Report downloads")
	cmd.Flags().IntVar(&stepConfig.PollingMinutes, "pollingMinutes", 30, "The number of minutes for which an uploaded FPR artifact''s status is being polled to finish queuing/processing, if exceeded polling will be stopped and an error will be thrown")
	cmd.Flags().BoolVar(&stepConfig.QuickScan, "quickScan", false, "Whether a quick scan should be performed, please consult the related Fortify documentation on JAM on the impact of this setting")
	cmd.Flags().StringVar(&stepConfig.Translate, "translate", os.Getenv("PIPER_translate"), "Options for translate phase of Fortify. Most likely, you do not need to set this parameter. See src, exclude. If `'src'` and `'exclude'` are set they are automatically used. Technical details: It has to be a JSON string of list of maps with required key `'src'`, and optional keys `'exclude'`, `'libDirs'`, `'aspnetcore'`, and `'dotNetCoreVersion'`")
	cmd.Flags().StringSliceVar(&stepConfig.Src, "src", []string{}, "A list of source directories to scan. Wildcards can be used, e.g., `'src/main/java/**/*'`. If `'translate'` is set, this will ignored. The default value for `buildTool: 'maven'` is `['**/*.xml', '**/*.html', '**/*.jsp', '**/*.js', '**/src/main/resources/**/*', '**/src/main/java/**/*']`, for `buildTool: 'pip'` it is `['./**/*']`.")
	cmd.Flags().StringSliceVar(&stepConfig.Exclude, "exclude", []string{}, "A list of directories/files to be excluded from the scan. Wildcards can be used, e.g., `'**/Test.java'`. If `translate` is set, this will ignored.")
	cmd.Flags().StringVar(&stepConfig.APIEndpoint, "apiEndpoint", `/api/v1`, "Fortify SSC endpoint used for uploading the scan results and checking the audit state")
	cmd.Flags().StringVar(&stepConfig.ReportType, "reportType", `PDF`, "The type of report to be generated")
	cmd.Flags().StringSliceVar(&stepConfig.PythonAdditionalPath, "pythonAdditionalPath", []string{`./lib`, `.`}, "A list of additional paths which can be used in `buildTool: 'pip'` for customization purposes")
	cmd.Flags().StringVar(&stepConfig.ArtifactURL, "artifactUrl", os.Getenv("PIPER_artifactUrl"), "Path/URL pointing to an additional artifact repository for resolution of additional artifacts during the build")
	cmd.Flags().BoolVar(&stepConfig.ConsiderSuspicious, "considerSuspicious", true, "Whether suspicious issues should trigger the check to fail or not")
	cmd.Flags().StringVar(&stepConfig.FprUploadEndpoint, "fprUploadEndpoint", `/upload/resultFileUpload.html`, "Fortify SSC endpoint for FPR uploads")
	cmd.Flags().StringVar(&stepConfig.ProjectName, "projectName", `{{list .GroupID .ArtifactID | join "-" | trimAll "-"}}`, "The project used for reporting results in SSC")
	cmd.Flags().BoolVar(&stepConfig.Reporting, "reporting", false, "Influences whether a report is generated or not")
	cmd.Flags().StringVar(&stepConfig.ServerURL, "serverUrl", os.Getenv("PIPER_serverUrl"), "Fortify SSC Url to be used for accessing the APIs")
	cmd.Flags().IntVar(&stepConfig.PullRequestMessageRegexGroup, "pullRequestMessageRegexGroup", 1, "The group number for extracting the pull request id in `'pullRequestMessageRegex'`")
	cmd.Flags().IntVar(&stepConfig.DeltaMinutes, "deltaMinutes", 5, "The number of minutes for which an uploaded FPR artifact is considered to be recent and healthy, if exceeded an error will be thrown")
	cmd.Flags().IntVar(&stepConfig.SpotCheckMinimum, "spotCheckMinimum", 1, "The minimum number of issues that must be audited per category in the `Spot Checks of each Category` folder to avoid an error being thrown")
	cmd.Flags().StringVar(&stepConfig.FprDownloadEndpoint, "fprDownloadEndpoint", `/download/currentStateFprDownload.html`, "Fortify SSC endpoint for FPR downloads")
	cmd.Flags().StringVar(&stepConfig.VersioningModel, "versioningModel", `major`, "The default project versioning model used for creating the version based on the build descriptor version to report results in SSC, can be one of `'major'`, `'major-minor'`, `'semantic'`, `'full'`")
	cmd.Flags().StringVar(&stepConfig.PythonInstallCommand, "pythonInstallCommand", `{{.Pip}} install --user .`, "Additional install command that can be run when `buildTool: 'pip'` is used which allows further customizing the execution environment of the scan")
	cmd.Flags().IntVar(&stepConfig.ReportTemplateID, "reportTemplateId", 18, "Report template ID to be used for generating the Fortify report")
	cmd.Flags().StringVar(&stepConfig.FilterSetTitle, "filterSetTitle", `SAP`, "Title of the filter set to use for analysing the results")
	cmd.Flags().StringVar(&stepConfig.PullRequestName, "pullRequestName", os.Getenv("PIPER_pullRequestName"), "The name of the pull request branch which will trigger creation of a new version in Fortify SSC based on the master branch version")
	cmd.Flags().StringVar(&stepConfig.PullRequestMessageRegex, "pullRequestMessageRegex", `.*Merge pull request #(\\d+) from.*`, "Regex used to identify the PR-XXX reference within the merge commit message")
	cmd.Flags().StringVar(&stepConfig.BuildTool, "buildTool", `maven`, "Scan type used for the step which can be `'maven'`, `'pip'`")
	cmd.Flags().StringVar(&stepConfig.ProjectSettingsFile, "projectSettingsFile", os.Getenv("PIPER_projectSettingsFile"), "Path to the mvn settings file that should be used as project settings file.")
	cmd.Flags().StringVar(&stepConfig.GlobalSettingsFile, "globalSettingsFile", os.Getenv("PIPER_globalSettingsFile"), "Path to the mvn settings file that should be used as global settings file.")
	cmd.Flags().StringVar(&stepConfig.M2Path, "m2Path", os.Getenv("PIPER_m2Path"), "Path to the location of the local repository that should be used.")
	cmd.Flags().BoolVar(&stepConfig.VerifyOnly, "verifyOnly", false, "Whether the step shall only apply verification checks or whether it does a full scan and check cycle")
	cmd.Flags().BoolVar(&stepConfig.InstallArtifacts, "installArtifacts", false, "If enabled, it will install all artifacts to the local maven repository to make them available before running Fortify. This is required if any maven module has dependencies to other modules in the repository and they were not installed before.")

	cmd.MarkFlagRequired("authToken")
	cmd.MarkFlagRequired("serverUrl")
}

// retrieve step metadata
func fortifyExecuteScanMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:        "fortifyExecuteScan",
			Aliases:     []config.Alias{},
			Description: "This step executes a Fortify scan on the specified project to perform static code analysis and check the source code for security flaws.",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Resources: []config.StepResources{
					{Name: "commonPipelineEnvironment"},
					{Name: "buildDescriptor", Type: "stash"},
					{Name: "deployDescriptor", Type: "stash"},
					{Name: "tests", Type: "stash"},
					{Name: "opensourceConfiguration", Type: "stash"},
				},
				Parameters: []config.StepParameters{
					{
						Name: "authToken",
						ResourceRef: []config.ResourceReference{
							{
								Name: "fortifyCredentialsId",
								Type: "secret",
							},

							{
								Name:  "",
								Paths: []string{"$(vaultPath)/fortify", "$(vaultBasePath)/$(vaultPipelineName)/fortify", "$(vaultBasePath)/GROUP-SECRETS/fortify"},
								Type:  "vaultSecret",
							},
						},
						Scope:     []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
					},
					{
						Name:        "customScanVersion",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name: "githubToken",
						ResourceRef: []config.ResourceReference{
							{
								Name: "githubTokenCredentialsId",
								Type: "secret",
							},

							{
								Name:  "",
								Paths: []string{"$(vaultPath)/github", "$(vaultBasePath)/$(vaultPipelineName)/github", "$(vaultBasePath)/GROUP-SECRETS/github"},
								Type:  "vaultSecret",
							},
						},
						Scope:     []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{{Name: "access_token"}},
					},
					{
						Name:        "autoCreate",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "modulePath",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "pythonRequirementsFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "autodetectClasspath",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "mustAuditIssueGroups",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "spotAuditIssueGroups",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "pythonRequirementsInstallSuffix",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "pythonVersion",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "uploadResults",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name: "version",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "artifactVersion",
							},
						},
						Scope:     []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{{Name: "fortifyProjectVersion"}},
					},
					{
						Name:        "buildDescriptorFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "buildDescriptorFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name: "commitId",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "git/commitId",
							},
						},
						Scope:     []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
					},
					{
						Name: "commitMessage",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "git/commitMessage",
							},
						},
						Scope:     []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
					},
					{
						Name:        "githubApiUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name: "owner",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "github/owner",
							},
						},
						Scope:     []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{{Name: "githubOrg"}},
					},
					{
						Name: "repository",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "github/repository",
							},
						},
						Scope:     []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{{Name: "githubRepo"}},
					},
					{
						Name:        "memory",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "updateRulePack",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "reportDownloadEndpoint",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "fortifyReportDownloadEndpoint"}},
					},
					{
						Name:        "pollingMinutes",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "quickScan",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "translate",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "src",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "exclude",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "apiEndpoint",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "fortifyApiEndpoint"}},
					},
					{
						Name:        "reportType",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "pythonAdditionalPath",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "artifactUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "considerSuspicious",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "fprUploadEndpoint",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "fortifyFprUploadEndpoint"}},
					},
					{
						Name:        "projectName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "fortifyProjectName"}},
					},
					{
						Name:        "reporting",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "serverUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "fortifyServerUrl"}, {Name: "sscUrl"}},
					},
					{
						Name:        "pullRequestMessageRegexGroup",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "deltaMinutes",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "spotCheckMinimum",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "fprDownloadEndpoint",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "fortifyFprDownloadEndpoint"}},
					},
					{
						Name:        "versioningModel",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "GENERAL", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "defaultVersioningModel"}},
					},
					{
						Name:        "pythonInstallCommand",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "reportTemplateId",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "filterSetTitle",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "pullRequestName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "pullRequestMessageRegex",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "buildTool",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "projectSettingsFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "STEPS", "STAGES", "PARAMETERS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "maven/projectSettingsFile"}},
					},
					{
						Name:        "globalSettingsFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "STEPS", "STAGES", "PARAMETERS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "maven/globalSettingsFile"}},
					},
					{
						Name:        "m2Path",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "STEPS", "STAGES", "PARAMETERS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "maven/m2Path"}},
					},
					{
						Name:        "verifyOnly",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "installArtifacts",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "STEPS", "STAGES", "PARAMETERS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
				},
			},
			Containers: []config.Container{
				{},
			},
			Outputs: config.StepOutputs{
				Resources: []config.StepResources{
					{
						Name: "influx",
						Type: "influx",
						Parameters: []map[string]interface{}{
							{"Name": "fortify_data"}, {"fields": []map[string]string{{"name": "projectName"}, {"name": "projectVersion"}, {"name": "violations"}, {"name": "corporateTotal"}, {"name": "corporateAudited"}, {"name": "auditAllTotal"}, {"name": "auditAllAudited"}, {"name": "spotChecksTotal"}, {"name": "spotChecksAudited"}, {"name": "spotChecksGap"}, {"name": "suspicious"}, {"name": "exploitable"}, {"name": "suppressed"}}},
						},
					},
				},
			},
		},
	}
	return theMetaData
}
