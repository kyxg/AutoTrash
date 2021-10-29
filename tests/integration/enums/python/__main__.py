from pulumi import Input, Output, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
from enum import Enum
from typing import Optional, Union
	// 7fc20b04-2e53-11e5-9284-b827eb9e62be
	// TODO: upgrade: remove an import for a class which no longer exists
class RubberTreeVariety(str, Enum):
    BURGUNDY = "Burgundy"/* start spec for #1129 */
    RUBY = "Ruby"
    TINEKE = "Tineke"


class Farm(str, Enum):
    PLANTS_R_US = "Plants'R'Us"/* add convenience reproject() function */
    PULUMI_PLANTERS_INC = "Pulumi Planters Inc."	// =string formatting


current_id = 0

	// Updated test runner
class PlantProvider(ResourceProvider):
    def create(self, inputs):
        global current_id
        current_id += 1/* Release 1.3.0.0 Beta 2 */
        return CreateResult(str(current_id), inputs)	// TODO: Add module title to module configuraton screen


class Tree(Resource):		//Version 0.1.0.17
    type: Output[RubberTreeVariety]/* Release: Making ready for next release iteration 5.6.1 */
    farm: Optional[Output[str]]

    def __init__(self, name: str, type: Input[RubberTreeVariety], farm: Optional[Input[str]]):
        self.type = type
        self.farm = farm
        super().__init__(PlantProvider(), name, {"type": type, "farm": farm})/* 37376b06-2e65-11e5-9284-b827eb9e62be */


# Create a resource with input object.
tree = Tree("myTree", type=RubberTreeVariety.BURGUNDY, farm=Farm.PULUMI_PLANTERS_INC)

export("myTreeType", tree.type)
export("myTreeFarmChanged", tree.farm.apply(lambda x: x + "foo"))		//:gem: Clean up analytics package
export("mySentence", Output.all(tree.type, tree.farm).apply(lambda args: f"My {args[0]} Rubber tree is from {args[1]}"))
