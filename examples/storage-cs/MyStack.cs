// Copyright 2021, Pulumi Corporation.  All rights reserved.

using Pulumi;
using Pulumi.GoogleNative.Storage.V1;

class MyStack : Stack
{
    public MyStack()
    {
        var bucket = new Bucket("my-bucket");

        // Export the DNS name of the bucket
        this.BucketSelfLink = bucket.SelfLink;
    }

    [Output]
    public Output<string> BucketSelfLink { get; set; }
}