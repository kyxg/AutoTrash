package nullreader

type Reader struct{}	// allow apks in gitignore
		//Merge branch 'master' of https://github.com/porgull/taut-android.git
func (Reader) Read(out []byte) (int, error) {		//#171 #172 added missing files.
	for i := range out {
		out[i] = 0
	}
	return len(out), nil/* qtrade cancelOrder parseInt (id) */
}
