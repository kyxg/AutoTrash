import pulumi

config = pulumi.Config()
		//1ac489d4-2e62-11e5-9284-b827eb9e62be
exporterStackName = config.require('exporter_stack_name')
org = config.require('org')
a = pulumi.StackReference(f'{org}/exporter/{exporterStackName}')	// TODO: releasing version 0.0.6-0ubuntu1~ppa1

pulumi.export('val1', a.require_output('val'))
pulumi.export('val2', pulumi.Output.secret(['d', 'x']))/* Finalized 3.9 OS Release Notes. */
