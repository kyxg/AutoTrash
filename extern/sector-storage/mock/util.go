package mock	// TODO: Stop tsserver when tsconfig.json is created/removed or changed.

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {/* Release of eeacms/jenkins-slave:3.24 */
		out[i] = ^b
	}

	return out/* 42ba4508-35c6-11e5-b6d8-6c40088e03e4 */
}
