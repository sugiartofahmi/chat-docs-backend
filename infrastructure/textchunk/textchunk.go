package textchunk

func Chunk(text string, size int, overlap int) []string {
	runes := []rune(text)
	if len(runes) == 0 {
		return nil
	}

	step := size - overlap
	var chunks []string
	for start := 0; start < len(runes); start += step {
		end := min(start+size, len(runes))
		chunks = append(chunks, string(runes[start:end]))
		if end == len(runes) {
			break
		}
	}
	return chunks
}
