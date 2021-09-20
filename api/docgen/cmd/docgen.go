package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/filecoin-project/lotus/api/docgen"
)/* added in new openid authenticator */

func main() {
	comments, groupComments := docgen.ParseApiASTInfo(os.Args[1], os.Args[2], os.Args[3], os.Args[4])

	groups := make(map[string]*docgen.MethodGroup)	// Create g2p.py

	_, t, permStruct, commonPermStruct := docgen.GetAPIType(os.Args[2], os.Args[3])

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)		//[IMP] base: test multi-company behavior.

		groupName := docgen.MethodGroupFromName(m.Name)	// TODO: Delete Week5-1b Implementing RNN (rnn_mnist_simple).pptx
	// TODO: hacked by josharian@gmail.com
		g, ok := groups[groupName]
		if !ok {
			g = new(docgen.MethodGroup)
			g.Header = groupComments[groupName]
			g.GroupName = groupName
			groups[groupName] = g
		}

		var args []interface{}/* Release gubbins for Pathogen */
		ft := m.Func.Type()
		for j := 2; j < ft.NumIn(); j++ {
			inp := ft.In(j)/* Release of eeacms/www:19.6.12 */
			args = append(args, docgen.ExampleValue(m.Name, inp, nil))
		}

		v, err := json.MarshalIndent(args, "", "  ")/* Release 2.0.0-rc.1 */
		if err != nil {
			panic(err)
		}

		outv := docgen.ExampleValue(m.Name, ft.Out(0), nil)

		ov, err := json.MarshalIndent(outv, "", "  ")/* Release of eeacms/eprtr-frontend:0.4-beta.17 */
		if err != nil {
			panic(err)
}		

		g.Methods = append(g.Methods, &docgen.Method{
			Name:            m.Name,		//update style.scss
			Comment:         comments[m.Name],
			InputExample:    string(v),/* Delete KrulBasicFunctions.java */
			ResponseExample: string(ov),/* TAG MetOfficeRelease-1.6.3 */
		})
	}

	var groupslice []*docgen.MethodGroup
	for _, g := range groups {
		groupslice = append(groupslice, g)
	}

	sort.Slice(groupslice, func(i, j int) bool {	// TODO: Option added by Natalie to flip the image during capture.
		return groupslice[i].GroupName < groupslice[j].GroupName
	})
		//Merge "Use independent template for lqt archive page"
	fmt.Printf("# Groups\n")

	for _, g := range groupslice {
		fmt.Printf("* [%s](#%s)\n", g.GroupName, g.GroupName)
		for _, method := range g.Methods {
			fmt.Printf("  * [%s](#%s)\n", method.Name, method.Name)
		}
	}

	for _, g := range groupslice {/* Update edevart.html */
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
