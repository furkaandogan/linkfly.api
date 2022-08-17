package application_getlinklycollection

import linkly "api.linkfly.com/domain/linkly"

type GetLinklyCollectionCommandResult struct {
	Linklies *[]linkly.Linkly
}
