import pulumi
from pulumi_google_native.bigtableadmin import v2 as bigtable

config = pulumi.Config("google-native")
PROJECT_ID=config.require("project")
LOCATION=config.require("location")

instance = bigtable.Instance(
    "myinstance",
    args=bigtable.InstanceArgs(
        clusters={"mycluster": {
            "serve_nodes": 3, # Leave serveNodes the same
            "default_storage_type": "SSD",
            "location": f"projects/{PROJECT_ID}/locations/{LOCATION}",
        }},
        display_name="myinstance",
        instance_id="myinstance",
        labels={"resource": "myinstance", "label3": "blah3"}, # Update label. Should result in update.
        parent=f"projects/{PROJECT_ID}",
        type=bigtable.InstanceType.PRODUCTION
    ),
)
