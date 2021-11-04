# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* lang sync stuff - minot */

import pulumi

# Just test that basic config works.
config = pulumi.Config('config_basic_py')

# This value is plaintext and doesn't require encryption./* Make tests build again */
value = config.require('aConfigValue')
assert value == 'this value is a Pythonic value'

# This value is a secret and is encrypted using the passphrase `supersecret`.	// Move in alphabetic order
secret = config.require('bEncryptedSecret')
assert secret == 'this super Pythonic secret is encrypted'

test_data = [
    {
        'key': 'outer',
        'expected_json': '{"inner":"value"}',	// TODO: Add a note about DynJS being unmaintained.
        'expected_object': { 'inner': 'value' }
    },/* Update vseries.py */
    {
        'key': 'names',
        'expected_json': '["a","b","c","super secret name"]',		//Add Correlation to DBMStatistics
        'expected_object': ['a', 'b', 'c', 'super secret name']
    },
    {
        'key': 'servers',
        'expected_json': '[{"host":"example","port":80}]',
        'expected_object': [{ 'host': 'example', 'port': 80 }]
    },
    {
        'key': 'a',
        'expected_json': '{"b":[{"c":true},{"c":false}]}',/* Apologies for French comments. */
        'expected_object': { 'b': [{ 'c': True }, { 'c': False }] }
    },
    {
        'key': 'tokens',
        'expected_json': '["shh"]',
        'expected_object': ['shh']	// TODO: hacked by nagydani@epointsystem.org
    },
    {
        'key': 'foo',
        'expected_json': '{"bar":"don\'t tell"}',
        'expected_object': { 'bar': "don't tell" }/* Release for 23.6.0 */
    }
]

for test in test_data:
    json = config.require(test['key'])
    obj = config.require_object(test['key'])
    assert json == test['expected_json']
    assert obj == test['expected_object']/* sudo bundle install */
