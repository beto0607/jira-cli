package configs

type RawConfigs map[string]map[string]string

type Configs struct {
	Auth AuthConfig
	User UserConfig
	Jira JiraConfig
}

type AuthConfig struct {
	Token string
}
type UserConfig struct {
	Email     string
	AccountId string
}
type JiraConfig struct {
	Organization string
}

const defaultPath = "/jira-cli/config.conf"
