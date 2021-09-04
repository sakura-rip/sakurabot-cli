package database

// StringsToDBTags return the array of Tag which name is given to param
func StringsToDBTags(strTags []string) []*Tag {
	var tags []*Tag
	for _, tagName := range strTags {
		var tag *Tag
		result := DefaultClient.Where(&Tag{Name: tagName}).First(tag)
		if result.RowsAffected == 0 {
			tag = &Tag{Name: tagName}
			DefaultClient.Create(tag)
		}
		tags = append(tags, tag)
	}
	return tags
}
