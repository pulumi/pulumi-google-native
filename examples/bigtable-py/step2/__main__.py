import pulumi
from pulumi_google_native.bigtableadmin import v2 as bigtable

config = pulumi.Config("google-native")
PROJECT_ID=config.require("project")
LOCATION=config.require("location")

instance = bigtable.Instance(
    "myinstance",
    args=bigtable.InstanceArgs(
        clusters={"mycluster": {
            "serve_nodes": 3, # Bump serve_nodes. This should result in replacement
            "default_storage_type": "SSD",
            "location": f"projects/{PROJECT_ID}/locations/{LOCATION}",
        }},
        display_name="myinstance",
        instance_id="myinstance",
        labels={"resource": "myinstance", "label1": "blah1"},
        parent=f"projects/{PROJECT_ID}",
        type=bigtable.InstanceType.PRODUCTION
    ),
)
