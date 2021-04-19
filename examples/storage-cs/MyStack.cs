// Copyright 2021, Pulumi Corporation.  All rights reserved.

using Pulumi;
using Pulumi.GoogleNative.Storage.V1;
using Pulumi.Random;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config("google-native");
        var project = config.Require("project");
        var region = config.Require("region");

        var randomString = new RandomString("name", new RandomStringArgs
        {
            Upper = false,
            Number = false,
            Special = false,
            Length = 8,
        });

        var bucket = new Bucket("my-bucket", new BucketArgs
        {
            Name = randomString.Result,
            Bucket = randomString.Result,
            Project = project,
        });

        // Export the DNS name of the bucket
        this.BucketSelfLink = bucket.SelfLink;
    }

    [Output]
    public Output<string> BucketSelfLink { get; set; }
}