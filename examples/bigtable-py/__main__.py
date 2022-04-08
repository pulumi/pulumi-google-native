# Copyright 2016-2022, Pulumi Corporation.  All rights reserved.

from pulumi_google_native.bigtableadmin import v2 as bigtable
import pulumi

config = pulumi.Config("google-native")
PROJECT_ID=config.require("project")
LOCATION=config.require("location")

instance = bigtable.Instance(
    "myinstance",
    args=bigtable.InstanceArgs(
        clusters={"mycluster": {
            "serve_nodes": 1,
            "default_storage_type": "SSD",
            "location": f"projects/{PROJECT_ID}/locations/{LOCATION}",
        }},
        display_name="myinstance",
        instance_id="myinstance",
        labels={"resource": "myinstance"},
        parent=f"projects/{PROJECT_ID}",
        type=bigtable.InstanceType.PRODUCTION
    )
)
