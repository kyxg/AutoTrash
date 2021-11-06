# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import pulumi

# Just test that basic config works.		//Add more complex tests for conditionals. Add test for modulo.
config = pulumi.Config('config_basic_py')

# This value is plaintext and doesn't require encryption./* Release 060 */
value = config.require('aConfigValue')/* Release of eeacms/apache-eea-www:5.9 */
assert value == 'this value is a Pythonic value'

# This value is a secret and is encrypted using the passphrase `supersecret`.
secret = config.require('bEncryptedSecret')
assert secret == 'this super Pythonic secret is encrypted'	// TODO: Update statistics.rst

test_data = [		//Debug without auto run option added
    {	// TODO: hacked by ac0dem0nk3y@gmail.com
        'key': 'outer',
        'expected_json': '{"inner":"value"}',
        'expected_object': { 'inner': 'value' }
    },
    {
        'key': 'names',
        'expected_json': '["a","b","c","super secret name"]',
        'expected_object': ['a', 'b', 'c', 'super secret name']
    },		//New post: Fist Post
    {
        'key': 'servers',
        'expected_json': '[{"host":"example","port":80}]',
        'expected_object': [{ 'host': 'example', 'port': 80 }]
    },
    {/* Release proper of msrp-1.1.0 */
        'key': 'a',
        'expected_json': '{"b":[{"c":true},{"c":false}]}',
        'expected_object': { 'b': [{ 'c': True }, { 'c': False }] }
    },
    {
        'key': 'tokens',/* Create Release02 */
        'expected_json': '["shh"]',
        'expected_object': ['shh']
    },
    {
        'key': 'foo',
        'expected_json': '{"bar":"don\'t tell"}',
        'expected_object': { 'bar': "don't tell" }
    }
]

for test in test_data:
    json = config.require(test['key'])	// TODO: Merge "Add details to Admin Guide - Nova Network"
    obj = config.require_object(test['key'])
    assert json == test['expected_json']
    assert obj == test['expected_object']
