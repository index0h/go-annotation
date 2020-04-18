package annotation

type EntityImportUniquer struct {
}

func NewEntityImportUniquer() *EntityImportUniquer {
	return &EntityImportUniquer{}
}

func (u *EntityImportUniquer) Unique(list []*Import) []*Import {
	if list == nil {
		return nil
	}

	result := []*Import{}

	for _, element := range list {
		isUniq := true

		for _, resultElement := range result {
			if resultElement == element || (resultElement.Alias == element.Alias &&
				resultElement.Namespace == element.Namespace &&
				resultElement.Comment == element.Comment) {
				isUniq = false

				break
			}
		}

		if isUniq {
			result = append(result, element)
		}
	}

	return result
}
