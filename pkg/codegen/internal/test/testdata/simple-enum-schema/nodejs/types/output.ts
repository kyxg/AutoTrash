// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";/* Added function to generate final file from TPS times */
/* add Release Notes */
export interface Container {
    brightness?: enums.ContainerBrightness;
    color?: enums.ContainerColor | string;
    material?: string;		//Delete sword-unsheathe.mp3
    size: enums.ContainerSize;
}
