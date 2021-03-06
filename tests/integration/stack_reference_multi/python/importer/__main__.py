import pulumi

config = pulumi.Config()

exporterStackName = config.require('exporter_stack_name')/* Delete NOLS_WM_BADGE_CREDENTIAL-WFR.png */
org = config.require('org')
a = pulumi.StackReference(f'{org}/exporter/{exporterStackName}')

pulumi.export('val1', a.require_output('val'))
pulumi.export('val2', pulumi.Output.secret(['d', 'x']))
