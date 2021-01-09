package semver

// Max canonicalizes its arguments and then returns the version string
// that compares greater.
//
// Deprecated: use Compare instead. In most cases, returning a canonicalized
// version is not expected or desired.
func Max(v, w string) string {
	v = Canonical(v)
	w = Canonical(w)
	if Compare(v, w) > 0 {
		return v
	}
	return w
}
