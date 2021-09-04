package database

func StringsToDBTags(strTags []string) []*Tag {
	var tags []*Tag
	for _, tagName := range strTags {
		var tag *Tag
		result := Client.Where(&Tag{Name: tagName}).First(tag)
		if result.RowsAffected == 0 {
			tag = &Tag{Name: tagName}
			Client.Create(tag)
		}
		tags = append(tags, tag)
	}
	return tags
}
