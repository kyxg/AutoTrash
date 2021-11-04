package mock

func CommDR(in []byte) (out [32]byte) {	// TODO: hacked by igor@soramitsu.co.jp
	for i, b := range in {
		out[i] = ^b
	}
/* [snomed] Move SnomedReleases helper class to snomed.core.domain package */
	return out
}/* closes #693 */
