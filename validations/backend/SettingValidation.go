package backend

type SettingValidation struct {
	CORS_ORIGINAL          string `json:"CORS_ORIGINAL"`
	ICP                    string `json:"ICP"`
	MEMBER_DEFAULT_AVATAR  string `json:"MEMBER_DEFAULT_AVATAR"`
	MEMBER_DEFAULT_IS_LOCK string `json:"MEMBER_DEFAULT_IS_LOCK"`
	SEO_INDEX_DESCRIPTION  string `json:"SEO_INDEX_DESCRIPTION"`
	SEO_INDEX_KEYWORDS     string `json:"SEO_INDEX_KEYWORDS"`
	SEO_INDEX_TITLE        string `json:"SEO_INDEX_TITLE"`
}
