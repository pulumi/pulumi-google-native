[![Deploy](https://get.pulumi.com/new/button.svg)](https://app.pulumi.com/new?template=https://github.com/pulumi/pulumi-google-native/blob/master/examples/webserver-ts/README.md)

# Web Server Using Compute Engine

Starting point for building the Pulumi web server sample in Google Cloud.

## Running the App

1.  Create a new stack:

    ```
    $ pulumi stack init dev
    ```

1.  Configure the project:

    ```
    $ pulumi config set google-native:project YOURGOOGLECLOUDPROJECT
    $ pulumi config set google-native:zone us-central1-a
    ```

1.  Restore NPM dependencies:

    ```
    $ npm install
    ```

1.  Run `pulumi up` to preview and deploy changes:

    ``` 
    $ pulumi up
    Previewing changes:
    ...
    Performing changes:
     Type                                 Name                                       Status      
+   pulumi:pulumi:Stack                   webserver-ts-dev                           created
+   ├─ google-native:compute/v1:Network   network                                    created
+   ├─ google-native:compute/v1:Firewall  firewall                                   created
+   └─ google-native:compute/v1:Instance  instance                                   created
    ---outputs:---
    instanceIP  : "35.224.37.178"
    instanceLink: "....webserver-instance"
    info: 4 changes performed:
    + 4 resources created
      Update duration: 57.918455655s
      ```

1.  Curl the HTTP server:

    ```
    $ curl $(pulumi stack output instanceIP)
    Hello, World!
    ```

1.  SSH into the server:

    ```
    $ gcloud compute ssh $(pulumi stack output instanceName)
    Warning: Permanently added 'compute.967481934451185713' (ECDSA) to the list of known hosts.
    The programs included with the Debian GNU/Linux system are free software;
    the exact distribution terms for each program are described in the
    individual files in /usr/share/doc/*/copyright.
    Debian GNU/Linux comes with ABSOLUTELY NO WARRANTY, to the extent
    permitted by applicable law.
    luke@instance-8ad9bd8:~$
    ```

1. Cleanup

    ```
    $ pulumi destroy
    $ pulumi stack rm
    ```
