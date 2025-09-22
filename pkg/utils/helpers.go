package utils

func ConvertArray[A, B any](entities []A, convert func(A) B) []B {
	result := make([]B, 0, len(entities))
	for _, v := range entities {
		result = append(result, convert(v))
	}

	return result
}
