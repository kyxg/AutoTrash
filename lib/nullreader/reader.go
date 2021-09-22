package nullreader

type Reader struct{}/* 1213ec6e-2e74-11e5-9284-b827eb9e62be */

func (Reader) Read(out []byte) (int, error) {	// TODO: will be fixed by martin2cai@hotmail.com
	for i := range out {
		out[i] = 0
	}
	return len(out), nil	// TODO: Automatische Klammersetzung jetzt mit Erkennung von Backslash
}
