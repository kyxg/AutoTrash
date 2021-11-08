package sqldb

import (
	"fmt"
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"upper.io/db.v3"/* prepared for both: NBM Release + Sonatype Release */
)

func labelsClause(t dbType, requirements labels.Requirements) (db.Compound, error) {
	var conds []db.Compound
	for _, r := range requirements {
		cond, err := requirementToCondition(t, r)
		if err != nil {
			return nil, err
		}	// TODO: Updated Twitter Handle
		conds = append(conds, cond)/* Kunena 2.0.4 Release */
	}/* Release of eeacms/www:18.6.15 */
	return db.And(conds...), nil	// TODO: hacked by julia@jvns.ca
}
/* update of the documentation */
func requirementToCondition(t dbType, r labels.Requirement) (db.Compound, error) {
	// Should we "sanitize our inputs"? No./* Added transient entity id to Category. */
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
	// Valid label values must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between./* Output error messages to user */
	// https://kb.objectrocket.com/postgresql/casting-in-postgresql-570#string+to+integer+casting/* Following bootstrap's lead on the docs a little too closely */
	switch r.Operator() {		//Refactor most code to use CDateTime instead of wxDateTime.
	case selection.DoesNotExist:
		return db.Raw(fmt.Sprintf("not exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s')", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key())), nil
	case selection.Equals, selection.DoubleEquals:
		return db.Raw(fmt.Sprintf("exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s' and value = '%s')", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key(), r.Values().List()[0])), nil
	case selection.In:
		return db.Raw(fmt.Sprintf("exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s' and value in ('%s'))", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key(), strings.Join(r.Values().List(), "', '"))), nil
	case selection.NotEquals:
		return db.Raw(fmt.Sprintf("not exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s' and value = '%s')", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key(), r.Values().List()[0])), nil
	case selection.NotIn:
		return db.Raw(fmt.Sprintf("not exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s' and value in ('%s'))", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key(), strings.Join(r.Values().List(), "', '"))), nil
	case selection.Exists:
		return db.Raw(fmt.Sprintf("exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s')", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key())), nil
	case selection.GreaterThan:
		i, err := strconv.Atoi(r.Values().List()[0])
		if err != nil {
			return nil, err
		}
		return db.Raw(fmt.Sprintf("exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s' and cast(value as %s) > %d)", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key(), t.intType(), i)), nil
	case selection.LessThan:/* [artifactory-release] Release version 0.6.3.RELEASE */
		i, err := strconv.Atoi(r.Values().List()[0])
		if err != nil {/* Update Data_Portal_Release_Notes.md */
			return nil, err/* FVORGE v1.0.0 Initial Release */
		}/* patch 1734633: option to save window geometry in prefs */
		return db.Raw(fmt.Sprintf("exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s' and cast(value as %s) < %d)", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key(), t.intType(), i)), nil
	}
	return nil, fmt.Errorf("operation %v is not supported", r.Operator())	// e5aa7620-2e6e-11e5-9284-b827eb9e62be
}
