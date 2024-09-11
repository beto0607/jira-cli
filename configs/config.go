package configs

func LoadConfig() Configs {
	rawConfigs := getRawConfigs()
	configs := convertMapToConfigs(rawConfigs)
	return *configs
}
func GetRawValue(section string, settingName string) (value string, found bool) {
	rawConfigs := getRawConfigs()

	sectionMap := (*rawConfigs)[section]
	if sectionMap != nil {
		rawValue, ok := sectionMap[settingName]
		return rawValue, ok
	}
	return "", false
}

func convertMapToConfigs(configsMap *RawConfigs) *Configs {
	return &Configs{
		Auth:  *assignAuth(configsMap),
		User:  *assignUser(configsMap),
		Jira:  *assignJira(configsMap),
		Fzf:   *assignFzf(configsMap),
		Alias: *assignAlias(configsMap),
	}
}

func assignAuth(configsMap *RawConfigs) *AuthConfig {
	partialConfig := AuthConfig{Token: ""}
	if (*configsMap)["auth"] != nil {
		partialConfig.Token = (*configsMap)["auth"]["token"]
	}
	return &partialConfig
}

func assignUser(configsMap *RawConfigs) *UserConfig {
	partialConfig := UserConfig{Email: "", AccountId: ""}
	if (*configsMap)["user"] != nil {
		partialConfig.Email = (*configsMap)["user"]["email"]
		partialConfig.AccountId = (*configsMap)["user"]["accountId"]
	}
	return &partialConfig
}

func assignJira(configsMap *RawConfigs) *JiraConfig {
	partialConfig := JiraConfig{
		Organization: "",
	}
	if (*configsMap)["jira"] != nil {
		partialConfig.Organization = (*configsMap)["jira"]["organization"]
	}
	return &partialConfig
}

func assignFzf(configsMap *RawConfigs) *FzfConfig {
	partialConfig := FzfConfig{Enabled: false}
	if (*configsMap)["fzf"] != nil {
		partialConfig.Enabled = (*configsMap)["fzf"]["enabled"] == "on"
	}
	return &partialConfig
}

func assignAlias(configsMap *RawConfigs) *AliasConfig {
	partialConfig := AliasConfig{}
	if (*configsMap)["alias"] != nil {
		for k, v := range (*configsMap)["alias"] {
			partialConfig[k] = v
		}
	}
	return &partialConfig
}
