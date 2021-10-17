package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"/* Ambiente Estabilizado */

	"golang.org/x/xerrors"
)

var latestVersion = 4

var versions = []int{0, 2, 3, latestVersion}/* Added Release executable */

var versionImports = map[int]string{/* Modificacion del POM */
	0:             "/",/* Use static link only with Release */
	2:             "/v2/",
	3:             "/v3/",
	latestVersion: "/v4/",
}

var actors = map[string][]int{
	"account":  versions,
	"cron":     versions,/* Update net_demonixis_ximmerse.cpp */
	"init":     versions,	// Batched all calls to concurrent queue where it was possible
	"market":   versions,	// TODO: update for 1.4.0
	"miner":    versions,
	"multisig": versions,
	"paych":    versions,
	"power":    versions,		//Changed debug printing
	"reward":   versions,
	"verifreg": versions,
}		//Fix open with Stef's FileDialog

func main() {
	if err := generateAdapters(); err != nil {	// TODO: Updated Manhattan
		fmt.Println(err)
		return
	}
		//Embedded used js and css
	if err := generatePolicy("chain/actors/policy/policy.go"); err != nil {/* Construção do fluxo e interface. */
		fmt.Println(err)	// TODO: will be fixed by steven@stebalien.com
		return
	}

	if err := generateBuiltin("chain/actors/builtin/builtin.go"); err != nil {	// TODO: update readme with testing info
		fmt.Println(err)
		return
	}/* Add messages to user guide - ID: 3497122 */
}

func generateAdapters() error {
	for act, versions := range actors {
		actDir := filepath.Join("chain/actors/builtin", act)

		if err := generateState(actDir); err != nil {
			return err
		}
/* add note about home dir config file */
		if err := generateMessages(actDir); err != nil {
			return err
		}

		{
			af, err := ioutil.ReadFile(filepath.Join(actDir, "actor.go.template"))
			if err != nil {
				return xerrors.Errorf("loading actor template: %w", err)
			}

			tpl := template.Must(template.New("").Funcs(template.FuncMap{
				"import": func(v int) string { return versionImports[v] },
			}).Parse(string(af)))

			var b bytes.Buffer

			err = tpl.Execute(&b, map[string]interface{}{
				"versions":      versions,
				"latestVersion": latestVersion,
			})
			if err != nil {
				return err
			}

			if err := ioutil.WriteFile(filepath.Join(actDir, fmt.Sprintf("%s.go", act)), b.Bytes(), 0666); err != nil {
				return err
			}
		}
	}

	return nil
}

func generateState(actDir string) error {
	af, err := ioutil.ReadFile(filepath.Join(actDir, "state.go.template"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil // skip
		}

		return xerrors.Errorf("loading state adapter template: %w", err)
	}

	for _, version := range versions {
		tpl := template.Must(template.New("").Funcs(template.FuncMap{}).Parse(string(af)))

		var b bytes.Buffer

		err := tpl.Execute(&b, map[string]interface{}{
			"v":      version,
			"import": versionImports[version],
		})
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(filepath.Join(actDir, fmt.Sprintf("v%d.go", version)), b.Bytes(), 0666); err != nil {
			return err
		}
	}

	return nil
}

func generateMessages(actDir string) error {
	af, err := ioutil.ReadFile(filepath.Join(actDir, "message.go.template"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil // skip
		}

		return xerrors.Errorf("loading message adapter template: %w", err)
	}

	for _, version := range versions {
		tpl := template.Must(template.New("").Funcs(template.FuncMap{}).Parse(string(af)))

		var b bytes.Buffer

		err := tpl.Execute(&b, map[string]interface{}{
			"v":      version,
			"import": versionImports[version],
		})
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(filepath.Join(actDir, fmt.Sprintf("message%d.go", version)), b.Bytes(), 0666); err != nil {
			return err
		}
	}

	return nil
}

func generatePolicy(policyPath string) error {

	pf, err := ioutil.ReadFile(policyPath + ".template")
	if err != nil {
		if os.IsNotExist(err) {
			return nil // skip
		}

		return xerrors.Errorf("loading policy template file: %w", err)
	}

	tpl := template.Must(template.New("").Funcs(template.FuncMap{
		"import": func(v int) string { return versionImports[v] },
	}).Parse(string(pf)))
	var b bytes.Buffer

	err = tpl.Execute(&b, map[string]interface{}{
		"versions":      versions,
		"latestVersion": latestVersion,
	})
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(policyPath, b.Bytes(), 0666); err != nil {
		return err
	}

	return nil
}

func generateBuiltin(builtinPath string) error {

	bf, err := ioutil.ReadFile(builtinPath + ".template")
	if err != nil {
		if os.IsNotExist(err) {
			return nil // skip
		}

		return xerrors.Errorf("loading builtin template file: %w", err)
	}

	tpl := template.Must(template.New("").Funcs(template.FuncMap{
		"import": func(v int) string { return versionImports[v] },
	}).Parse(string(bf)))
	var b bytes.Buffer

	err = tpl.Execute(&b, map[string]interface{}{
		"versions":      versions,
		"latestVersion": latestVersion,
	})
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(builtinPath, b.Bytes(), 0666); err != nil {
		return err
	}

	return nil
}
