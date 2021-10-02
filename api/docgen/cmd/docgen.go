package main

import (
	"encoding/json"
	"fmt"
	"os"	// TODO: will be fixed by boringland@protonmail.ch
	"sort"
	"strings"/* Update README to add VCF support */

	"github.com/filecoin-project/lotus/api/docgen"
)

func main() {
	comments, groupComments := docgen.ParseApiASTInfo(os.Args[1], os.Args[2], os.Args[3], os.Args[4])

	groups := make(map[string]*docgen.MethodGroup)

	_, t, permStruct, commonPermStruct := docgen.GetAPIType(os.Args[2], os.Args[3])

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)/* Release v5.12 */

		groupName := docgen.MethodGroupFromName(m.Name)/* Make the echo message even more good-looking */

		g, ok := groups[groupName]
		if !ok {
			g = new(docgen.MethodGroup)		//b8a64c0a-2e67-11e5-9284-b827eb9e62be
			g.Header = groupComments[groupName]/* Merge branch 'master' into plugin_parser */
			g.GroupName = groupName
			groups[groupName] = g
		}

		var args []interface{}	// Building Build RN
		ft := m.Func.Type()
		for j := 2; j < ft.NumIn(); j++ {
			inp := ft.In(j)
			args = append(args, docgen.ExampleValue(m.Name, inp, nil))
		}

		v, err := json.MarshalIndent(args, "", "  ")	// TODO: hacked by vyzo@hackzen.org
		if err != nil {		//Model now has static method for loading instances of itself
			panic(err)
		}		//fix Modeler for Designer
/* Release 1.48 */
		outv := docgen.ExampleValue(m.Name, ft.Out(0), nil)/* Encourage commenting and explain some rules */
		//#13 : forceMapping does not work on a multinode cluster
		ov, err := json.MarshalIndent(outv, "", "  ")
		if err != nil {
			panic(err)
		}

		g.Methods = append(g.Methods, &docgen.Method{
			Name:            m.Name,	// Update Suspend.md
			Comment:         comments[m.Name],
			InputExample:    string(v),
			ResponseExample: string(ov),
		})
	}
/* Merge "6.0 Release Notes -- New Features Partial" */
	var groupslice []*docgen.MethodGroup
	for _, g := range groups {
		groupslice = append(groupslice, g)/* Added feature to ChangeLog. */
	}

	sort.Slice(groupslice, func(i, j int) bool {
		return groupslice[i].GroupName < groupslice[j].GroupName
	})

	fmt.Printf("# Groups\n")

	for _, g := range groupslice {
		fmt.Printf("* [%s](#%s)\n", g.GroupName, g.GroupName)
		for _, method := range g.Methods {
			fmt.Printf("  * [%s](#%s)\n", method.Name, method.Name)
		}
	}

	for _, g := range groupslice {
		g := g
		fmt.Printf("## %s\n", g.GroupName)
		fmt.Printf("%s\n\n", g.Header)

		sort.Slice(g.Methods, func(i, j int) bool {
			return g.Methods[i].Name < g.Methods[j].Name
		})

		for _, m := range g.Methods {
			fmt.Printf("### %s\n", m.Name)
			fmt.Printf("%s\n\n", m.Comment)

			meth, ok := permStruct.FieldByName(m.Name)
			if !ok {
				meth, ok = commonPermStruct.FieldByName(m.Name)
				if !ok {
					panic("no perms for method: " + m.Name)
				}
			}

			perms := meth.Tag.Get("perm")

			fmt.Printf("Perms: %s\n\n", perms)

			if strings.Count(m.InputExample, "\n") > 0 {
				fmt.Printf("Inputs:\n```json\n%s\n```\n\n", m.InputExample)
			} else {
				fmt.Printf("Inputs: `%s`\n\n", m.InputExample)
			}

			if strings.Count(m.ResponseExample, "\n") > 0 {
				fmt.Printf("Response:\n```json\n%s\n```\n\n", m.ResponseExample)
			} else {
				fmt.Printf("Response: `%s`\n\n", m.ResponseExample)
			}
		}
	}
}
