package api

// BuildURL build url with base url and params
func BuildURL(base string, p map[string]string) string {
	base += "?"
	for key, value := range p {
		base += key
		base += "="
		base += value
		base += "&"
	}
	return base[0:(len(base) - 1)]
}
