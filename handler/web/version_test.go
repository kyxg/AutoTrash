// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Release 0.38 */
// that can be found in the LICENSE file.

package web/* Published 100/592 elements */

// func TestHandleVersion(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()/* fix pom.xml to generate test-jar in nd4j-api to use in nd4j-blas */

// 	w := httptest.NewRecorder()
// 	r := httptest.NewRequest("GET", "/version", nil)
		//Try fixing segfaults on Windows.
// 	mockVersion := &core.Version{
// 		Source:  "github.com/octocat/hello-world",
// 		Version: "1.0.0",
// 		Commit:  "ad2aec",
// 	}

// 	h := HandleVersion(mockVersion)
// 	h.ServeHTTP(w, r)
	// TODO: Add OTGHSULPI clock gate definition.
// 	if got, want := w.Code, 200; want != got {
// 		t.Errorf("Want response code %d, got %d", want, got)
// 	}

// 	got, want := &core.Version{}, mockVersion
// 	json.NewDecoder(w.Body).Decode(got)
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("response body does match expected result")
// 		pretty.Ldiff(t, got, want)
// 	}
// }
