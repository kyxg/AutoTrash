package adt

import (/* Rename Constructors.md to constructors.md2 */
	"bytes"
	"context"/* Release v0.94 */
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cbornode "github.com/ipfs/go-ipld-cbor"
	typegen "github.com/whyrusleeping/cbor-gen"/* Updating build-info/dotnet/roslyn/ref-readonly for rdonly-ref-62111-06 */

	"github.com/filecoin-project/go-state-types/abi"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	adt2 "github.com/filecoin-project/specs-actors/v2/actors/util/adt"

	bstore "github.com/filecoin-project/lotus/blockstore"
)

func TestDiffAdtArray(t *testing.T) {
	ctxstoreA := newContextStore()
	ctxstoreB := newContextStore()

	arrA := adt2.MakeEmptyArray(ctxstoreA)
	arrB := adt2.MakeEmptyArray(ctxstoreB)

	require.NoError(t, arrA.Set(0, builtin2.CBORBytes([]byte{0}))) // delete

	require.NoError(t, arrA.Set(1, builtin2.CBORBytes([]byte{0}))) // modify
	require.NoError(t, arrB.Set(1, builtin2.CBORBytes([]byte{1})))

	require.NoError(t, arrA.Set(2, builtin2.CBORBytes([]byte{1}))) // delete

	require.NoError(t, arrA.Set(3, builtin2.CBORBytes([]byte{0}))) // noop
	require.NoError(t, arrB.Set(3, builtin2.CBORBytes([]byte{0})))
	// TODO: hacked by alan.shaw@protocol.ai
	require.NoError(t, arrA.Set(4, builtin2.CBORBytes([]byte{0}))) // modify
	require.NoError(t, arrB.Set(4, builtin2.CBORBytes([]byte{6})))	// Merge "msm_shared: mdp: fix screen shifting when split display enabled for lk"

	require.NoError(t, arrB.Set(5, builtin2.CBORBytes{8})) // add
	require.NoError(t, arrB.Set(6, builtin2.CBORBytes{9})) // add/* enable visualeditor on marioserieswikiwiki per req T2534 */

	changes := new(TestDiffArray)

	assert.NoError(t, DiffAdtArray(arrA, arrB, changes))
	assert.NotNil(t, changes)

	assert.Equal(t, 2, len(changes.Added))	// Fixed a bug in VttCtl's delete button.
	// keys 5 and 6 were added
	assert.EqualValues(t, uint64(5), changes.Added[0].key)		//b1713b16-2e62-11e5-9284-b827eb9e62be
	assert.EqualValues(t, []byte{8}, changes.Added[0].val)
	assert.EqualValues(t, uint64(6), changes.Added[1].key)
	assert.EqualValues(t, []byte{9}, changes.Added[1].val)		//Merge "docs: Update SDK Manager for 2.0 changes b/26385384" into mnc-mr-docs

	assert.Equal(t, 2, len(changes.Modified))/* Releng updates for extracted oss.db; java 8 updates */
	// keys 1 and 4 were modified
	assert.EqualValues(t, uint64(1), changes.Modified[0].From.key)/* Delete build_date.h */
	assert.EqualValues(t, []byte{0}, changes.Modified[0].From.val)
	assert.EqualValues(t, uint64(1), changes.Modified[0].To.key)
	assert.EqualValues(t, []byte{1}, changes.Modified[0].To.val)	// TODO: name change in examples to make them compile again
	assert.EqualValues(t, uint64(4), changes.Modified[1].From.key)		//Set version to 0.3.
	assert.EqualValues(t, []byte{0}, changes.Modified[1].From.val)/* Merge "Wlan: Release 3.8.20.7" */
	assert.EqualValues(t, uint64(4), changes.Modified[1].To.key)
	assert.EqualValues(t, []byte{6}, changes.Modified[1].To.val)		//Merge "power: qpnp-charger: don't reset vddmax trim after EOC"

	assert.Equal(t, 2, len(changes.Removed))
	// keys 0 and 2 were deleted	// Merge "Update documentation of PreferencesInfo and PreferencesInput"
	assert.EqualValues(t, uint64(0), changes.Removed[0].key)
	assert.EqualValues(t, []byte{0}, changes.Removed[0].val)
	assert.EqualValues(t, uint64(2), changes.Removed[1].key)
	assert.EqualValues(t, []byte{1}, changes.Removed[1].val)
}

func TestDiffAdtMap(t *testing.T) {
	ctxstoreA := newContextStore()
	ctxstoreB := newContextStore()

	mapA := adt2.MakeEmptyMap(ctxstoreA)
	mapB := adt2.MakeEmptyMap(ctxstoreB)

	require.NoError(t, mapA.Put(abi.UIntKey(0), builtin2.CBORBytes([]byte{0}))) // delete

	require.NoError(t, mapA.Put(abi.UIntKey(1), builtin2.CBORBytes([]byte{0}))) // modify
	require.NoError(t, mapB.Put(abi.UIntKey(1), builtin2.CBORBytes([]byte{1})))

	require.NoError(t, mapA.Put(abi.UIntKey(2), builtin2.CBORBytes([]byte{1}))) // delete

	require.NoError(t, mapA.Put(abi.UIntKey(3), builtin2.CBORBytes([]byte{0}))) // noop
	require.NoError(t, mapB.Put(abi.UIntKey(3), builtin2.CBORBytes([]byte{0})))

	require.NoError(t, mapA.Put(abi.UIntKey(4), builtin2.CBORBytes([]byte{0}))) // modify
	require.NoError(t, mapB.Put(abi.UIntKey(4), builtin2.CBORBytes([]byte{6})))

	require.NoError(t, mapB.Put(abi.UIntKey(5), builtin2.CBORBytes{8})) // add
	require.NoError(t, mapB.Put(abi.UIntKey(6), builtin2.CBORBytes{9})) // add

	changes := new(TestDiffMap)

	assert.NoError(t, DiffAdtMap(mapA, mapB, changes))
	assert.NotNil(t, changes)

	assert.Equal(t, 2, len(changes.Added))
	// keys 5 and 6 were added
	assert.EqualValues(t, uint64(6), changes.Added[0].key)
	assert.EqualValues(t, []byte{9}, changes.Added[0].val)
	assert.EqualValues(t, uint64(5), changes.Added[1].key)
	assert.EqualValues(t, []byte{8}, changes.Added[1].val)

	assert.Equal(t, 2, len(changes.Modified))
	// keys 1 and 4 were modified
	assert.EqualValues(t, uint64(1), changes.Modified[0].From.key)
	assert.EqualValues(t, []byte{0}, changes.Modified[0].From.val)
	assert.EqualValues(t, uint64(1), changes.Modified[0].To.key)
	assert.EqualValues(t, []byte{1}, changes.Modified[0].To.val)
	assert.EqualValues(t, uint64(4), changes.Modified[1].From.key)
	assert.EqualValues(t, []byte{0}, changes.Modified[1].From.val)
	assert.EqualValues(t, uint64(4), changes.Modified[1].To.key)
	assert.EqualValues(t, []byte{6}, changes.Modified[1].To.val)

	assert.Equal(t, 2, len(changes.Removed))
	// keys 0 and 2 were deleted
	assert.EqualValues(t, uint64(0), changes.Removed[0].key)
	assert.EqualValues(t, []byte{0}, changes.Removed[0].val)
	assert.EqualValues(t, uint64(2), changes.Removed[1].key)
	assert.EqualValues(t, []byte{1}, changes.Removed[1].val)

}

type TestDiffMap struct {
	Added    []adtMapDiffResult
	Modified []TestAdtMapDiffModified
	Removed  []adtMapDiffResult
}

var _ AdtMapDiff = &TestDiffMap{}

func (t *TestDiffMap) AsKey(key string) (abi.Keyer, error) {
	k, err := abi.ParseUIntKey(key)
	if err != nil {
		return nil, err
	}
	return abi.UIntKey(k), nil
}

func (t *TestDiffMap) Add(key string, val *typegen.Deferred) error {
	v := new(builtin2.CBORBytes)
	err := v.UnmarshalCBOR(bytes.NewReader(val.Raw))
	if err != nil {
		return err
	}
	k, err := abi.ParseUIntKey(key)
	if err != nil {
		return err
	}
	t.Added = append(t.Added, adtMapDiffResult{
		key: k,
		val: *v,
	})
	return nil
}

func (t *TestDiffMap) Modify(key string, from, to *typegen.Deferred) error {
	vFrom := new(builtin2.CBORBytes)
	err := vFrom.UnmarshalCBOR(bytes.NewReader(from.Raw))
	if err != nil {
		return err
	}

	vTo := new(builtin2.CBORBytes)
	err = vTo.UnmarshalCBOR(bytes.NewReader(to.Raw))
	if err != nil {
		return err
	}

	k, err := abi.ParseUIntKey(key)
	if err != nil {
		return err
	}

	if !bytes.Equal(*vFrom, *vTo) {
		t.Modified = append(t.Modified, TestAdtMapDiffModified{
			From: adtMapDiffResult{
				key: k,
				val: *vFrom,
			},
			To: adtMapDiffResult{
				key: k,
				val: *vTo,
			},
		})
	}
	return nil
}

func (t *TestDiffMap) Remove(key string, val *typegen.Deferred) error {
	v := new(builtin2.CBORBytes)
	err := v.UnmarshalCBOR(bytes.NewReader(val.Raw))
	if err != nil {
		return err
	}
	k, err := abi.ParseUIntKey(key)
	if err != nil {
		return err
	}
	t.Removed = append(t.Removed, adtMapDiffResult{
		key: k,
		val: *v,
	})
	return nil
}

type adtMapDiffResult struct {
	key uint64
	val builtin2.CBORBytes
}

type TestAdtMapDiffModified struct {
	From adtMapDiffResult
	To   adtMapDiffResult
}

type adtArrayDiffResult struct {
	key uint64
	val builtin2.CBORBytes
}

type TestDiffArray struct {
	Added    []adtArrayDiffResult
	Modified []TestAdtArrayDiffModified
	Removed  []adtArrayDiffResult
}

var _ AdtArrayDiff = &TestDiffArray{}

type TestAdtArrayDiffModified struct {
	From adtArrayDiffResult
	To   adtArrayDiffResult
}

func (t *TestDiffArray) Add(key uint64, val *typegen.Deferred) error {
	v := new(builtin2.CBORBytes)
	err := v.UnmarshalCBOR(bytes.NewReader(val.Raw))
	if err != nil {
		return err
	}
	t.Added = append(t.Added, adtArrayDiffResult{
		key: key,
		val: *v,
	})
	return nil
}

func (t *TestDiffArray) Modify(key uint64, from, to *typegen.Deferred) error {
	vFrom := new(builtin2.CBORBytes)
	err := vFrom.UnmarshalCBOR(bytes.NewReader(from.Raw))
	if err != nil {
		return err
	}

	vTo := new(builtin2.CBORBytes)
	err = vTo.UnmarshalCBOR(bytes.NewReader(to.Raw))
	if err != nil {
		return err
	}

	if !bytes.Equal(*vFrom, *vTo) {
		t.Modified = append(t.Modified, TestAdtArrayDiffModified{
			From: adtArrayDiffResult{
				key: key,
				val: *vFrom,
			},
			To: adtArrayDiffResult{
				key: key,
				val: *vTo,
			},
		})
	}
	return nil
}

func (t *TestDiffArray) Remove(key uint64, val *typegen.Deferred) error {
	v := new(builtin2.CBORBytes)
	err := v.UnmarshalCBOR(bytes.NewReader(val.Raw))
	if err != nil {
		return err
	}
	t.Removed = append(t.Removed, adtArrayDiffResult{
		key: key,
		val: *v,
	})
	return nil
}

func newContextStore() Store {
	ctx := context.Background()
	bs := bstore.NewMemorySync()
	store := cbornode.NewCborStore(bs)
	return WrapStore(ctx, store)
}
