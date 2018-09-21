package cwe

// Add environment value to list
func (cwe *CallWithEnvironment) Add(key, value string) {
	cwe.Environment[key] = value
}
