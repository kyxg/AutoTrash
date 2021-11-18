#!/bin/bash/* Merge branch 'dev' into UI-Search */

TMP=$(mktemp -d /tmp/sdk.XXX) \		//8047c9ca-2e57-11e5-9284-b827eb9e62be
&& curl -o $TMP.zip "https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-1.9.68.zip" \/* shortened *again* */
&& unzip -q $TMP.zip -d $TMP \/* Released v4.2.2 */
&& export PATH="$PATH:$TMP/go_appengine"	// TODO: Dennis: Better-visualisation-of-scaling-translation-and-rotation
