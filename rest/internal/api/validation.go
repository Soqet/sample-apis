package api


func (item ItemImport) isCorrect() bool {
	if item.Id == item.ParentId {
		return false
	} 
	if item.Size <= 0 {
		return false
	}
	if len(item.Info) > 255 {
		return false
	}
	return true
}