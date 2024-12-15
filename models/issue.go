package models

type AvatarUrls struct {
	I16x16 string `json:"16X16"`
	I24x24 string `json:"24X24"`
	I32x32 string `json:"32X32"`
	I48x48 string `json:"48X48"`
}
type Issue struct {
	Fields struct {
		Watcher struct {
			IsWatching bool   `json:"isWatching"`
			Self       string `json:"self"`
			WatchCount int
		}
		Attachment []struct {
			Author struct {
				AccountId   string
				AccountType string
				Active      bool
				AvatarUrls  AvatarUrls
				DisplayName string
				Key         string
				Name        string
				Self        string
			}
			Content   string
			Created   string
			Filename  string
			Id        string
			MimeType  string
			Self      string
			Size      int
			Thumbnail string
		}
		SubTasks []struct {
			Id           string
			OutwardIssue struct {
				Fields struct {
					Status struct {
						IconUrl string
						Name    string
					}
				}
				Id   string
				Key  string
				Self string
			}
		}
		Description struct {
			Type    string
			Version int
			Content []Content
		}
		Project struct {
			AvatarUrls AvatarUrls
			Id         string
			Insight    struct {
				LastIssueUpdateTime string
				TotalIssueCount     int
			}
			Key             string
			Name            string
			ProjectCategory struct {
				Description string
				Id          string
				Name        string
				Self        string
			}
			Self       string
			Simplified bool
			Style      string
		}
		Comment []struct {
			Author CommentAuthor
			Body   struct {
				Type    string
				Version int
				Content []Content
			}
			Created     string
			Id          string
			Self        string
			UpdaeAuthor CommentAuthor
			Updated     string
			Visibility  struct {
				Identifier string
				Type       string
				Value      string
			}
		}
		IssueLinks []struct {
		}
	}
}

type CommentAuthor struct {
	AccountId   string `json:"accountId"`
	Active      bool   `json:"active"`
	DisplayName string `json:"displayName"`
	Self        string `json:"self"`
}
type Content struct {
	Type    string `json:"type"`
	Content []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"content"`
}
