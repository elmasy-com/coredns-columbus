package columbus

// Ready implements the ready.Readiness interface, once this flips to true CoreDNS
// assumes this plugin is ready for queries; it is not checked again.
func (e Columbus) Ready() bool { return true }
