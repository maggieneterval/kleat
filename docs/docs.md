# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [appengine_provider.proto](#appengine_provider.proto)
    - [AppengineAccount](#proto.AppengineAccount)
    - [AppengineProvider](#proto.AppengineProvider)
  
  
  
  

- [artifact_template.proto](#artifact_template.proto)
    - [ArtifactTemplate](#proto.ArtifactTemplate)
  
  
  
  

- [artifacts.proto](#artifacts.proto)
    - [ArtifactProviders](#proto.ArtifactProviders)
  
  
  
  

- [aws_provider.proto](#aws_provider.proto)
    - [AwsAccount](#proto.AwsAccount)
    - [AwsBakeryDefaults](#proto.AwsBakeryDefaults)
    - [AwsBaseImage](#proto.AwsBaseImage)
    - [AwsBaseImageSettings](#proto.AwsBaseImageSettings)
    - [AwsFeatures](#proto.AwsFeatures)
    - [AwsFeatures.CloudFormation](#proto.AwsFeatures.CloudFormation)
    - [AwsLifecycleHook](#proto.AwsLifecycleHook)
    - [AwsProvider](#proto.AwsProvider)
    - [AwsRegion](#proto.AwsRegion)
    - [AwsVirtualizationSettings](#proto.AwsVirtualizationSettings)
  
  
  
  

- [azure_provider.proto](#azure_provider.proto)
    - [AzureAccount](#proto.AzureAccount)
    - [AzureBakeryDefaults](#proto.AzureBakeryDefaults)
    - [AzureBaseImage](#proto.AzureBaseImage)
    - [AzureBaseImageSettings](#proto.AzureBaseImageSettings)
    - [AzureProvider](#proto.AzureProvider)
  
  
  
  

- [bitbucket_artifact_provider.proto](#bitbucket_artifact_provider.proto)
    - [BitbucketArtifactAccount](#proto.BitbucketArtifactAccount)
    - [BitbucketArtifactProvider](#proto.BitbucketArtifactProvider)
  
  
  
  

- [clouddriver.proto](#clouddriver.proto)
    - [ClouddriverConfig](#proto.ClouddriverConfig)
  
  
  
  

- [cloudfoundry_provider.proto](#cloudfoundry_provider.proto)
    - [CloudFoundryAccount](#proto.CloudFoundryAccount)
    - [CloudFoundryProvider](#proto.CloudFoundryProvider)
  
  
  
  

- [dcos_provider.proto](#dcos_provider.proto)
    - [DcosAccount](#proto.DcosAccount)
    - [DcosAccountCluster](#proto.DcosAccountCluster)
    - [DcosAccountDockerRegistry](#proto.DcosAccountDockerRegistry)
    - [DcosCluster](#proto.DcosCluster)
    - [DcosClusterLoadBalancer](#proto.DcosClusterLoadBalancer)
    - [DcosProvider](#proto.DcosProvider)
  
  
  
  

- [docker_registry_provider.proto](#docker_registry_provider.proto)
    - [DockerRegistryAccount](#proto.DockerRegistryAccount)
    - [DockerRegistryProvider](#proto.DockerRegistryProvider)
  
  
  
  

- [ecs_provider.proto](#ecs_provider.proto)
    - [EcsAccount](#proto.EcsAccount)
    - [EcsProvider](#proto.EcsProvider)
  
  
  
  

- [front50.proto](#front50.proto)
    - [Front50Config](#proto.Front50Config)
    - [Front50Config.Spinnaker](#proto.Front50Config.Spinnaker)
  
  
  
  

- [gcs_artifact_provider.proto](#gcs_artifact_provider.proto)
    - [GcsArtifactAccount](#proto.GcsArtifactAccount)
    - [GcsArtifactProvider](#proto.GcsArtifactProvider)
  
  
  
  

- [github_artifact_provider.proto](#github_artifact_provider.proto)
    - [GitHubArtifactAccount](#proto.GitHubArtifactAccount)
    - [GitHubArtifactProvider](#proto.GitHubArtifactProvider)
  
  
  
  

- [gitlab_artifact_provider.proto](#gitlab_artifact_provider.proto)
    - [GitLabArtifactAccount](#proto.GitLabArtifactAccount)
    - [GitLabArtifactProvider](#proto.GitLabArtifactProvider)
  
  
  
  

- [gitrepo_artifact_provider.proto](#gitrepo_artifact_provider.proto)
    - [GitRepoArtifactAccount](#proto.GitRepoArtifactAccount)
    - [GitRepoArtifactProvider](#proto.GitRepoArtifactProvider)
  
  
  
  

- [google_provider.proto](#google_provider.proto)
    - [Consul](#proto.Consul)
    - [GoogleAccount](#proto.GoogleAccount)
    - [GoogleBakeryDefaults](#proto.GoogleBakeryDefaults)
    - [GoogleBaseImage](#proto.GoogleBaseImage)
    - [GoogleBaseImageSettings](#proto.GoogleBaseImageSettings)
    - [GoogleProvider](#proto.GoogleProvider)
    - [VirtualizationSettings](#proto.VirtualizationSettings)
  
  
  
  

- [halconfig.proto](#halconfig.proto)
    - [HalConfig](#proto.HalConfig)
    - [HalConfig.PersistentStorage](#proto.HalConfig.PersistentStorage)
    - [HalConfig.Providers](#proto.HalConfig.Providers)
  
  
  
  

- [helm_artifact_provider.proto](#helm_artifact_provider.proto)
    - [HelmArtifactAccount](#proto.HelmArtifactAccount)
    - [HelmArtifactProvider](#proto.HelmArtifactProvider)
  
  
  
  

- [http_artifact_provider.proto](#http_artifact_provider.proto)
    - [HttpArtifactAccount](#proto.HttpArtifactAccount)
    - [HttpArtifactProvider](#proto.HttpArtifactProvider)
  
  
  
  

- [huaweicloud_provider.proto](#huaweicloud_provider.proto)
    - [HuaweiCloudAccount](#proto.HuaweiCloudAccount)
    - [HuaweiCloudBakeryDefaults](#proto.HuaweiCloudBakeryDefaults)
    - [HuaweiCloudBaseImage](#proto.HuaweiCloudBaseImage)
    - [HuaweiCloudBaseImageSettings](#proto.HuaweiCloudBaseImageSettings)
    - [HuaweiCloudProvider](#proto.HuaweiCloudProvider)
    - [HuaweiCloudVirtualizationSettings](#proto.HuaweiCloudVirtualizationSettings)
  
  
  
  

- [kubernetes_provider.proto](#kubernetes_provider.proto)
    - [KubernetesAccount](#proto.KubernetesAccount)
    - [KubernetesAccountDockerRegistry](#proto.KubernetesAccountDockerRegistry)
    - [KubernetesCachingPolicy](#proto.KubernetesCachingPolicy)
    - [KubernetesCustomResource](#proto.KubernetesCustomResource)
    - [KubernetesProvider](#proto.KubernetesProvider)
  
  
  
  

- [maven_artifact_provider.proto](#maven_artifact_provider.proto)
    - [MavenArtifactAccount](#proto.MavenArtifactAccount)
    - [MavenArtifactProvider](#proto.MavenArtifactProvider)
  
  
  
  

- [oracle_artifact_provider.proto](#oracle_artifact_provider.proto)
    - [OracleArtifactAccount](#proto.OracleArtifactAccount)
    - [OracleArtifactProvider](#proto.OracleArtifactProvider)
  
  
  
  

- [oracle_provider.proto](#oracle_provider.proto)
    - [OracleAccount](#proto.OracleAccount)
    - [OracleBakeryDefaults](#proto.OracleBakeryDefaults)
    - [OracleBaseImage](#proto.OracleBaseImage)
    - [OracleBaseImageSettings](#proto.OracleBaseImageSettings)
    - [OracleProvider](#proto.OracleProvider)
    - [OracleVirtualizationSettings](#proto.OracleVirtualizationSettings)
  
  
  
  

- [permissions.proto](#permissions.proto)
    - [Permissions](#proto.Permissions)
  
  
  
  

- [persistent_storage.proto](#persistent_storage.proto)
    - [GCS](#proto.GCS)
  
  
  
  

- [s3_artifact_provider.proto](#s3_artifact_provider.proto)
    - [S3ArtifactAccount](#proto.S3ArtifactAccount)
    - [S3ArtifactProvider](#proto.S3ArtifactProvider)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="appengine_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## appengine_provider.proto



<a name="proto.AppengineAccount"></a>

### AppengineAccount
Configuration for an App Engine account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cachingIntervalSeconds | [int32](#int32) |  | The interval in seconds at which Spinnaker will poll for updates in your App Engine clusters. |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| gcloudReleaseTrack | [string](#string) |  | The gcloud release track (`ALPHA`, `BETA`, or `STABLE`) that Spinnaker will use when deploying to App Engine. |
| gitHttpsUsername | [string](#string) |  | A username to be used when connecting to a remote git repository server over HTTPS. If set, `gitHttpsPassword` must also be set. |
| gitHttpsPassword | [string](#string) |  | A password to be used when connecting to a remote git repository server over HTTPS. If set, `gitHttpsUsername` must also be set. |
| githubOAuthAccessToken | [string](#string) |  | An OAuth token provided by Github for connecting to a git repository over HTTPS. See https://help.github.com/articles/creating-an-access-token-for-command-line-use for more information. |
| jsonPath | [string](#string) |  | The path to a JSON service account that Spinnaker will use as credentials. This is only needed if Spinnaker is not deployed on a Google Compute Engine VM, or needs permissions not afforded to the VM it is running on. See https://cloud.google.com/compute/docs/access/service-accounts for more information. |
| localRepositoryDirectory | [string](#string) |  | A local directory to be used to stage source files for App Engine deployments within Clouddriver. Defaults to `/var/tmp/clouddriver`. |
| omitServices | [string](#string) | repeated | A list of regular expressions. Any service matching one of these regexes will be ignored by Spinnaker. |
| omitVersions | [string](#string) | repeated | A list of regular expressions. Any version matching one of these regexes will be ignored by Spinnaker. |
| project | [string](#string) |  | The Google Cloud Platform project this Spinnaker account will manage. |
| permissions | [Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated): List of required Fiat permission groups. Configure `permissions` instead. |
| services | [string](#string) | repeated | A list of regular expressions. Any service matching one of these regexes will be indexed by Spinnaker (unless the service also matches a regex in `omitServices`). |
| sshKnownHostsFilePath | [string](#string) |  | The path to a `known_hosts` file to be used when connecting with a remote git repository over SSH. |
| sshPrivateKeyFilePath | [string](#string) |  | The path to an SSH private key to be used when connecting with a remote git repository over SSH. If set, `sshPrivateKeyPassphrase` must also be set. |
| sshPrivateKeyPassphrase | [string](#string) |  | The passphrase to an SSH private key to be used when connecting with a remote git repository over SSH. If set, `sshPrivateKeyFilePath` must also be set. |
| sshTrustUnknownHosts | [bool](#bool) |  | Enabling this flag will allow Spinnaker to connect with a remote git repository over SSH without verifying the server&#39;s IP address against a `known_hosts` file. Defaults to false. |
| versions | [string](#string) | repeated | A list of regular expressions. Any version matching one of these regexes will be indexed by Spinnaker (unless the version also matches a regex in `omitVersions`). |
| name | [string](#string) |  | The name of the account. |






<a name="proto.AppengineProvider"></a>

### AppengineProvider
Configuration for the Google App Engine (GAE) provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [AppengineAccount](#proto.AppengineAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |





 

 

 

 



<a name="artifact_template.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## artifact_template.proto



<a name="proto.ArtifactTemplate"></a>

### ArtifactTemplate
Configuration for a Jinja template for Spinnaker to use for artifact
extraction. For more details, please read the documentation:
https://www.spinnaker.io/reference/artifacts/from-build-triggers/#artifacts-from-build-triggers


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the artifact template. |
| templatePath | [string](#string) |  | The path to the artifact template. |





 

 

 

 



<a name="artifacts.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## artifacts.proto



<a name="proto.ArtifactProviders"></a>

### ArtifactProviders



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bitbucket | [BitbucketArtifactProvider](#proto.BitbucketArtifactProvider) |  |  |
| gcs | [GcsArtifactProvider](#proto.GcsArtifactProvider) |  |  |
| github | [GitHubArtifactProvider](#proto.GitHubArtifactProvider) |  |  |
| gitlab | [GitLabArtifactProvider](#proto.GitLabArtifactProvider) |  |  |
| gitrepo | [GitRepoArtifactProvider](#proto.GitRepoArtifactProvider) |  |  |
| helm | [HelmArtifactProvider](#proto.HelmArtifactProvider) |  |  |
| http | [HttpArtifactProvider](#proto.HttpArtifactProvider) |  |  |
| maven | [MavenArtifactProvider](#proto.MavenArtifactProvider) |  |  |
| oracle | [OracleArtifactProvider](#proto.OracleArtifactProvider) |  |  |
| s3 | [S3ArtifactProvider](#proto.S3ArtifactProvider) |  |  |
| templates | [ArtifactTemplate](#proto.ArtifactTemplate) | repeated |  |





 

 

 

 



<a name="aws_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## aws_provider.proto



<a name="proto.AwsAccount"></a>

### AwsAccount
Configuration for an AWS account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accountId | [string](#string) |  | The AWS account ID to manage. See http://docs.aws.amazon.com/IAM/latest/UserGuide/console_account-alias.html for more information. |
| assumeRole | [string](#string) |  | If set, Spinnaker will configure a credentials provider that uses AWS Security Token Service to assume the specified role. Examples: `user/spinnaker`, `role/spinnakerManaged`. |
| defaultKeyPair | [string](#string) |  | The name of the AWS key-pair to use. See http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html for more information. |
| discovery | [string](#string) |  | The endpoint at which your Eureka discovery system is reachable. See https://github.com/Netflix/eureka for more information. Example: `http://.eureka.url.to.use:8080/eureka-server/v2`. Using will make Spinnaker use AWS regions in the hostname to access discovery so that you can have discovery for multiple regions. |
| edda | [string](#string) |  | The endpoint at which Edda is reachable. Edda is not a hard dependency of Spinnaker, but is helpful for reducing the request volume against AWS. See https://github.com/Netflix/edda for more information. |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| permissions | [Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated): List of required Fiat permission groups. Configure `permissions` instead. |
| lifecycleHooks | [AwsLifecycleHook](#proto.AwsLifecycleHook) | repeated | List of configured AWS lifecycle hooks. |
| regions | [AwsRegion](#proto.AwsRegion) | repeated | List of configured AWS regions. |
| name | [string](#string) |  | The name of the account. |






<a name="proto.AwsBakeryDefaults"></a>

### AwsBakeryDefaults
Configuration for Spinnaker&#39;s image bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| awsAccessKey | [string](#string) |  | The default access key used to communicate with AWS. |
| awsSecretKey | [string](#string) |  | The secret key used to communicate with AWS. |
| awsSubnetId | [string](#string) |  | If using VPC, the default ID of the subnet, such as `subnet-12345def`, where Packer will launch the EC2 instance. This field is required if you are using a non-default VPC. |
| awsVpcId | [string](#string) |  | If launching into a VPC subnet, Packer needs the VPC ID in order to create a temporary security group within the VPC. Requires `subnet_id` to be set. If this default value is left blank, Packer will try to get the VPC ID from `awsSubnetId`. |
| awsAssociatePublicIpAddress | [bool](#bool) |  | If using a non-default VPC, public IP addresses are not provided by default. If this is enabled, your new instance will get a Public IP. |
| defaultVirtualizationType | [string](#string) |  | The default type of virtualization for the AMI you are building. This option must match the supported virtualization type of `AwsVirtualizationSettings.sourceAmi`. Acceptable values: `pv`, `hvm`. |
| baseImages | [AwsBaseImageSettings](#proto.AwsBaseImageSettings) | repeated | List of configured base images. |
| templateFile | [string](#string) |  | This is the name of the packer template that will be used to bake images from this base image. The template file must be found in this list https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or supplied as described here: https://spinnaker.io/setup/bakery/. |






<a name="proto.AwsBaseImage"></a>

### AwsBaseImage
Base image configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | This is the identifier used by AWS to find this base image. |
| shortDescription | [string](#string) |  | A short description to help human operators identify the image. |
| detailedDescription | [string](#string) |  | A long description to help human operators identify the image. |
| packageType | [string](#string) |  | This is used to help Spinnaker&#39;s bakery download the build artifacts you supply it with. For example, specifying deb indicates that your artifacts will need to be fetched from a debian repository. |
| templateFile | [string](#string) |  | The name of the Packer template that will be used to bake images from this base image. The template file must be found in this list: https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or supplied as described here: https://spinnaker.io/setup/bakery/. |






<a name="proto.AwsBaseImageSettings"></a>

### AwsBaseImageSettings
Configuration for a base image for the AWS provider&#39;s bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseImage | [AwsBaseImage](#proto.AwsBaseImage) |  | Base image configuration. |
| virtualizationSettings | [AwsVirtualizationSettings](#proto.AwsVirtualizationSettings) |  | Base image virtualization settings. |






<a name="proto.AwsFeatures"></a>

### AwsFeatures
Configuration for AWS-specific features.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cloudFormation | [AwsFeatures.CloudFormation](#proto.AwsFeatures.CloudFormation) |  | Configuration for AWS CloudFormation. |






<a name="proto.AwsFeatures.CloudFormation"></a>

### AwsFeatures.CloudFormation
Configuration for AWS CloudFormation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether AWS CloudFormation is enabled. |






<a name="proto.AwsLifecycleHook"></a>

### AwsLifecycleHook
Configuration for AWS Auto Scaling Lifecycle Hooks. For more information, see:
https://docs.aws.amazon.com/autoscaling/ec2/userguide/lifecycle-hooks.html


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| defaultResult | [string](#string) |  | Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Acceptable values: `CONTINUE`, `ABANDON`. |
| heartbeatTimeout | [int32](#int32) |  | Set the heartbeat timeout in seconds for the lifecycle hook. Instances can remain in a wait state for a finite period of time. Must be greater than or equal to 30 and less than or equal to 7200. The default is 3600 (one hour). |
| lifecycleTransition | [string](#string) |  | Type of lifecycle transition. Acceptable values: `autoscaling:EC2_INSTANCE_LAUNCHING`, `autoscaling:EC2_INSTANCE_TERMINATING` |
| notificationTargetARN | [string](#string) |  | The ARN of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in the transition state for the lifecycle hook. This target can be either an SQS queue or an SNS topic. |
| roleARN | [string](#string) |  | The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target, for example, an Amazon SNS topic or an Amazon SQS queue. |






<a name="proto.AwsProvider"></a>

### AwsProvider
Configuration for the AWS provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [AwsAccount](#proto.AwsAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |
| accessKeyId | [string](#string) |  | Your AWS Access Key ID. Note that if you are baking AMIs with Rosco, you may also need to set `AwsBakeryDefaults.awsAccessKey`. |
| secretAccessKey | [string](#string) |  | Your AWS Secret Key. Note that if you are baking AMIs with Rosco, you may also need to set `AwsBakeryDefaults.awsSecretKey`. |
| defaultRegions | [AwsRegion](#proto.AwsRegion) | repeated | List of default regions. |
| features | [AwsFeatures](#proto.AwsFeatures) |  | Configuration for AWS-specific features. |
| bakeryDefaults | [AwsBakeryDefaults](#proto.AwsBakeryDefaults) |  | Configuration for Spinnaker&#39;s image bakery. |






<a name="proto.AwsRegion"></a>

### AwsRegion
An AWS region.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the region. |






<a name="proto.AwsVirtualizationSettings"></a>

### AwsVirtualizationSettings
Base image virtualization settings.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| region | [string](#string) |  | The name of the region in which to launch the EC2 instance to create the AMI. |
| virtualizationType | [string](#string) |  | The type of virtualization for the AMI you are building. This option must match the supported virtualization type of `sourceAmi`. Acceptable values: `pv`, `hvm`. |
| instanceType | [string](#string) |  | The EC2 instance type to use while building the AMI, such as `t2.small`. |
| sourceAmi | [string](#string) |  | The source AMI whose root volume will be copied and provisioned on the currently running instance. This must be an EBS-backed AMI with a root volume snapshot that you have access to. |
| sshUserName | [string](#string) |  | The username to connect to SSH with. Required if using SSH. |
| winRmUserName | [string](#string) |  | The username to use to connect to WinRM. |
| spotPrice | [string](#string) |  | The maximum hourly price to pay for a spot instance to create the AMI. Spot instances are a type of instance that EC2 starts when the current spot price is less than the maximum price you specify. Spot price will be updated based on available spot instance capacity and current spot instance requests. It may save you some costs. You can set this to `auto` for Packer to automatically discover the best spot price or to &#34;0&#34; to use an on demand instance (default). |
| spotPriceAutoProduct | [string](#string) |  | Required if `spotPrice` is set to `auto`. This tells Packer what sort of AMI you are launching to find the best spot price. This must be one of: `Linux/UNIX`, `SUSE Linux`, `Windows`, `Linux/UNIX (Amazon VPC)`, `SUSE Linux (Amazon VPC)`, `Windows (Amazon VPC)`. |





 

 

 

 



<a name="azure_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## azure_provider.proto



<a name="proto.AzureAccount"></a>

### AzureAccount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| appKey | [string](#string) |  | (Required) The `appKey` (password) of your service principal. |
| clientId | [string](#string) |  | (Required) The `clientId` (also called `appId`) of your service principal. |
| defaultKeyVault | [string](#string) |  | (Required) The name of a KeyVault that contains the user name, password, and ssh public key used to create VMs |
| defaultResourceGroup | [string](#string) |  | (Required) The default resource group to contain any non-application specific resources. |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| objectId | [string](#string) |  | The `objectId` of your service principal. This is only required if using Packer to bake Windows images. |
| packerResourceGroup | [string](#string) |  | The resource group to use if baking images with Packer. |
| packerStorageAccount | [string](#string) |  | The storage account to use if baking images with Packer. |
| permissions | [Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated): List of required Fiat permission groups. Configure `permissions` instead. |
| regions | [string](#string) | repeated | The Azure regions this Spinnaker account will manage. |
| subscriptionId | [string](#string) |  | (Required) The `subscriptionId` to which your service principal is assigned. |
| tenantId | [string](#string) |  | (Required) The `tenantId` to which your service principal is assigned. |
| useSshPublicKey | [bool](#bool) |  | If true, the SSH public key is used to provision the linux VM. If false, the password is used instead. |






<a name="proto.AzureBakeryDefaults"></a>

### AzureBakeryDefaults
Configuration for Spinnaker&#39;s image bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseImages | [AzureBaseImageSettings](#proto.AzureBaseImageSettings) | repeated | List of configured base images. |






<a name="proto.AzureBaseImage"></a>

### AzureBaseImage
Base image configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shortDescription | [string](#string) |  | A short description to help human operators identify the image. |
| detailedDescription | [string](#string) |  | A long description to help human operators identify the image. |
| publisher | [string](#string) |  | (Required) The Publisher name for your base image. See https://aka.ms/azspinimage to get a list of images. |
| offer | [string](#string) |  | (Required) The offer for your base image. See https://aka.ms/azspinimage to get a list of images. |
| sku | [string](#string) |  | (Required) The SKU for your base image. See https://aka.ms/azspinimage to get a list of images. |
| version | [string](#string) |  | The version of your base image. This defaults to `latest` if not specified. |
| packageType | [string](#string) |  | This is used to help Spinnaker&#39;s bakery download the build artifacts you supply it with. For example, specifying `deb` indicates that your artifacts will need to be fetched from a debian repository. |
| templateFile | [string](#string) |  | This is the name of the packer template that will be used to bake images from this base image. The template file must be found in this list: https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or supplied as described here: https://spinnaker.io/setup/bakery/. |






<a name="proto.AzureBaseImageSettings"></a>

### AzureBaseImageSettings
Configuration for a base image for the Azure provider&#39;s bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseImage | [AzureBaseImage](#proto.AzureBaseImage) |  | Base image configuration. |






<a name="proto.AzureProvider"></a>

### AzureProvider
Configuration for the Azure provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [AzureAccount](#proto.AzureAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |
| bakeryDefaults | [AzureBakeryDefaults](#proto.AzureBakeryDefaults) |  | Configuration for Spinnaker&#39;s image bakery. |





 

 

 

 



<a name="bitbucket_artifact_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## bitbucket_artifact_provider.proto



<a name="proto.BitbucketArtifactAccount"></a>

### BitbucketArtifactAccount
Configuration for a Bitbucket artifact account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| username | [string](#string) |  | The username of the account. Either `username` and `password` should be set, or `usernamePasswordFile` should be set. |
| password | [string](#string) |  | The password of the account. Either `username` and `password` should be set, or `usernamePasswordFile` should be set. |
| usernamePasswordFile | [string](#string) |  | The path to a file containing the username and password of the account in the format `${username}:${password}`. Either `username` and `password` should be set, or `usernamePasswordFile` should be set. |






<a name="proto.BitbucketArtifactProvider"></a>

### BitbucketArtifactProvider
Configuration for the Bitbucket artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the Bitbucket artifact provider is enabled. |
| accounts | [BitbucketArtifactAccount](#proto.BitbucketArtifactAccount) | repeated | The list of configured Bitbucket accounts. |





 

 

 

 



<a name="clouddriver.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## clouddriver.proto



<a name="proto.ClouddriverConfig"></a>

### ClouddriverConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes | [KubernetesProvider](#proto.KubernetesProvider) |  |  |
| google | [GoogleProvider](#proto.GoogleProvider) |  |  |
| appengine | [AppengineProvider](#proto.AppengineProvider) |  |  |
| aws | [AwsProvider](#proto.AwsProvider) |  |  |
| azure | [AzureProvider](#proto.AzureProvider) |  |  |
| cloudfoundry | [CloudFoundryProvider](#proto.CloudFoundryProvider) |  |  |
| dcos | [DcosProvider](#proto.DcosProvider) |  |  |
| dockerRegistry | [DockerRegistryProvider](#proto.DockerRegistryProvider) |  |  |
| ecs | [EcsProvider](#proto.EcsProvider) |  |  |
| huaweicloud | [HuaweiCloudProvider](#proto.HuaweiCloudProvider) |  |  |
| oracle | [OracleProvider](#proto.OracleProvider) |  |  |
| artifacts | [ArtifactProviders](#proto.ArtifactProviders) |  |  |





 

 

 

 



<a name="cloudfoundry_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cloudfoundry_provider.proto



<a name="proto.CloudFoundryAccount"></a>

### CloudFoundryAccount
Configuration for a Spinnaker Cloud Foundry account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| api | [string](#string) |  | (Required) Host of the Cloud Foundry Foundation API endpoint (e.g., `api.sys.somesystem.com`). |
| appsManagerUri | [string](#string) |  | HTTP(S) URL of the Apps Manager application for the Cloud Foundry Foundation (e.g., `https://apps.sys.somesystem.com`). |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| metricsUri | [string](#string) |  | HTTP(S) URL of the metrics application for the Cloud Foundry Foundation (e.g., `https://metrics.sys.somesystem.com`). |
| password | [string](#string) |  | (Required) Password for the account to use for this Cloud Foundry Foundation. |
| skipSslValidation | [bool](#bool) |  | (Default: `false`) Skip SSL server certificate validation of the API endpoint. |
| user | [string](#string) |  | (Required) User name for the account to use for this Cloud Foundry Foundation. |
| permissions | [Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated): List of required Fiat permission groups. Configure `permissions` instead. |






<a name="proto.CloudFoundryProvider"></a>

### CloudFoundryProvider
Configuration for the Cloud Foundry provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [CloudFoundryAccount](#proto.CloudFoundryAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |





 

 

 

 



<a name="dcos_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dcos_provider.proto



<a name="proto.DcosAccount"></a>

### DcosAccount
Credentials to authenticate against one or more DC/OS clusters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | (Required) The name of the account. |
| clusters | [DcosAccountCluster](#proto.DcosAccountCluster) | repeated | (Required) The clusters against which this account will authenticate. |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| dockerRegistries | [DcosAccountDockerRegistry](#proto.DcosAccountDockerRegistry) | repeated | (Required) The list of Docker registries to use with this DC/OS account. |
| permissions | [Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated) List of required Fiat permission groups. Configure `permissions` instead. |






<a name="proto.DcosAccountCluster"></a>

### DcosAccountCluster
Configuration for a DC/OS cluster associated with a `DcosAccount`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | (Required) The name of the cluster. Must match the name of a `DcosCluster` defined for this provider. |
| uid | [string](#string) |  | (Required) User or service account identifier. |
| serviceKeyFile | [string](#string) |  | Path to a file containing the secret key for service account authentication. If set, `password` should not be set. |
| password | [string](#string) |  | Password for a user account. If set, `serviceKeyFile` should not be set. |






<a name="proto.DcosAccountDockerRegistry"></a>

### DcosAccountDockerRegistry
Configuration for a Docker registry associated with a `DcosAccount`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accountName | [string](#string) |  | The name of the Docker registry. Must be the name of an account configured with the Docker registry provider. |






<a name="proto.DcosCluster"></a>

### DcosCluster
Configuration for a DC/OS cluster.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | (Required) The name of the cluster. |
| caCertFile | [string](#string) |  | Root certificate file to trust for connections to the cluster. |
| dcosUrl | [string](#string) |  | (Required) URL of the endpoint for the DC/OS cluster&#39;s admin router. |
| loadBalancer | [DcosClusterLoadBalancer](#proto.DcosClusterLoadBalancer) |  | Configuration for a DC/OS load balancer. |
| insecureSkipTlsVerify | [bool](#bool) |  | If `true`, disables verification of certificates from the cluster (insecure). |






<a name="proto.DcosClusterLoadBalancer"></a>

### DcosClusterLoadBalancer
Configuration for a DC/OS load balancer.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| image | [string](#string) |  | Marathon-lb image to use when creating a load balancer with Spinnaker. |
| serviceAccountSecret | [string](#string) |  | Name of the secret to use for allowing marathon-lb to authenticate with the cluster. Only necessary for clusters with strict or permissive security. |






<a name="proto.DcosProvider"></a>

### DcosProvider
Configuration for the DC/OS (Distributed Cloud Operating System) provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [DcosAccount](#proto.DcosAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |
| clusters | [DcosCluster](#proto.DcosCluster) | repeated | The list of configured clusters. |





 

 

 

 



<a name="docker_registry_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## docker_registry_provider.proto



<a name="proto.DockerRegistryAccount"></a>

### DockerRegistryAccount
A credential able to authenticate against a set of Docker repositories.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| address | [string](#string) |  | (Required) The registry address from which to pull and deploy images (e.g., `https://index.docker.io`). |
| cacheIntervalSeconds | [int32](#int32) |  | The number of seconds between polling the Docker registry. Certain registries are sensitive to over-polling, and larger intervals (e.g., 10 minutes = 600 seconds) are desirable if you experience rate limiting. Defaults to `30`. |
| cacheThreads | [int32](#int32) |  | The number of threads on which to cache all provided repositories. Really only useful if you have a ton of repos. Defaults to 1. |
| clientTimeoutMillis | [int32](#int32) |  | Timeout in milliseconds for provided repositories. Defaults to `60,000`. |
| email | [string](#string) |  | The email associated with your Docker registry. Often this only needs to be well-formed, rather than be a real address. |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| insecureRegistry | [bool](#bool) |  | If `true`, Spinnaker will treat the Docker registry as insecure and not validate the SSL certificate. Defaults to `false`. |
| paginateSize | [int32](#int32) |  | Pagination size for the Docker `repository _catalog` endpoint. Defaults to `100`. |
| password | [string](#string) |  | The Docker registry password. Only one of `password`, `passwordCommand`, and `passwordFile` should be specified. |
| passwordCommand | [string](#string) |  | Command to retrieve Docker token/password. The command must be available in the environment. Only one of `password`, `passwordCommand`, and `passwordFile` should be specified. |
| passwordFile | [string](#string) |  | The path to a file containing your Docker password in plaintext (not a Docker `config.json` file). Only one of `password`, `passwordCommand`, and `passwordFile` should be specified. |
| permissions | [Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated) List of required Fiat permission groups. Configure `permissions` instead. |
| repositories | [string](#string) | repeated | An optional list of repositories from which to cache images. If not provided, Spinnaker will attempt to read accessible repositories from the `registries _catalog` endpoint. |
| sortTagsByDate | [bool](#bool) |  | If `true`, Spinnaker will sort tags by creation date. Defaults to `false`. |
| trackDigests | [bool](#bool) |  | If `true`, Spinnaker will track digest changes. This is not recommended because it greatly increases queries to the registry, and most registries are flaky. Defaults to `false`. |
| username | [string](#string) |  | The username associated with this Docker registry. |






<a name="proto.DockerRegistryProvider"></a>

### DockerRegistryProvider
Configuration for the Docker Registry provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [DockerRegistryAccount](#proto.DockerRegistryAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |





 

 

 

 



<a name="ecs_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ecs_provider.proto



<a name="proto.EcsAccount"></a>

### EcsAccount
Configuration for an ECS account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| awsAccount | [string](#string) |  | (Required) Provide the name of the AWS account associated with this ECS account. See https://github.com/spinnaker/clouddriver/blob/master/clouddriver-ecs/README.md for more information. |
| permissions | [Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated) List of required Fiat permission groups. Configure `permissions` instead. |






<a name="proto.EcsProvider"></a>

### EcsProvider
Configuration for the ECS provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [EcsAccount](#proto.EcsAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |





 

 

 

 



<a name="front50.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## front50.proto



<a name="proto.Front50Config"></a>

### Front50Config



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spinnaker | [Front50Config.Spinnaker](#proto.Front50Config.Spinnaker) |  |  |






<a name="proto.Front50Config.Spinnaker"></a>

### Front50Config.Spinnaker



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| persistentStoreType | [string](#string) |  |  |
| gcs | [GCS](#proto.GCS) |  |  |





 

 

 

 



<a name="gcs_artifact_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gcs_artifact_provider.proto



<a name="proto.GcsArtifactAccount"></a>

### GcsArtifactAccount
Configuration for a GCS artifact account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| jsonPath | [string](#string) |  | The path to a JSON key for a GCP service account with which to authenticate. The service account must have the `roles/storage.admin` role enabled. |






<a name="proto.GcsArtifactProvider"></a>

### GcsArtifactProvider
Configuration for the Google Cloud Storage (GCS) artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the GCS artifact provider is enabled. |
| accounts | [GcsArtifactAccount](#proto.GcsArtifactAccount) | repeated | The list of configured GCS accounts. |





 

 

 

 



<a name="github_artifact_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github_artifact_provider.proto



<a name="proto.GitHubArtifactAccount"></a>

### GitHubArtifactAccount
Configuration for a GitHub artifact account. Either `username` and
`password`, `usernamePasswordFile`, `token`, or `tokenFile` should be
specified as means of authentication.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| username | [string](#string) |  | The GitHub account username. |
| password | [string](#string) |  | The GitHub account password. |
| usernamePasswordFile | [string](#string) |  | The path to a file containing the username and password of the account in the format `${username}:${password}`. |
| token | [string](#string) |  | The GitHub access token. |
| tokenFile | [string](#string) |  | The path to a file containing the GitHub access token. |






<a name="proto.GitHubArtifactProvider"></a>

### GitHubArtifactProvider
Configuration for the GitHub artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the GitHub artifact provider is enabled. |
| accounts | [GitHubArtifactAccount](#proto.GitHubArtifactAccount) | repeated | The list of configured GitHub accounts. |





 

 

 

 



<a name="gitlab_artifact_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gitlab_artifact_provider.proto



<a name="proto.GitLabArtifactAccount"></a>

### GitLabArtifactAccount
Configuration for a GitLab artifact account. Either `token` or
`tokenFile` should be specified as means of authentication.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| token | [string](#string) |  | The GitLab access token. |
| tokenFile | [string](#string) |  | The path to a file containing the GitLab access token. |






<a name="proto.GitLabArtifactProvider"></a>

### GitLabArtifactProvider
Configuration for the GitLab artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the GitLab artifact provider is enabled. |
| accounts | [GitLabArtifactAccount](#proto.GitLabArtifactAccount) | repeated | The list of configured GitLab accounts. |





 

 

 

 



<a name="gitrepo_artifact_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gitrepo_artifact_provider.proto



<a name="proto.GitRepoArtifactAccount"></a>

### GitRepoArtifactAccount
Configuration for a Git repo artifact account. An account maps to a
credential that is able to authenticate against a Git repository hosted by a
Git hosting service. Either `username` and `password`,
`usernamePasswordFile`, `token`, `tokenFile`, or `sshPrivateKeyFilePath` and
`sshPrivateKeyPassphrase` must be set.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| username | [string](#string) |  | The username of the account. |
| password | [string](#string) |  | The password of the account. |
| usernamePasswordFile | [string](#string) |  | The path to a file containing the username and password of the account in the format `${username}:${password}`. |
| token | [string](#string) |  | The Git repo access token. |
| tokenFile | [string](#string) |  | The path to a file containing the Git repo access token. |
| sshPrivateKeyFilePath | [string](#string) |  | The path to an SSH private key to be used when connecting with the Git repo over SSH. |
| sshPrivateKeyPassphrase | [string](#string) |  | The passphrase to an SSH private key to be used when connecting with the Git repo over SSH. |
| sshKnownHostsFilePath | [string](#string) |  | The path to a `known_hosts` file to be used when connecting with a Git repository over SSH. |
| sshTrustUnknownHosts | [bool](#bool) |  | If `true`, Spinnaker can connect with a Git repository over SSH without verifying the server&#39;s IP address against a `known_hosts` file. |






<a name="proto.GitRepoArtifactProvider"></a>

### GitRepoArtifactProvider
Configuration for the Git repo artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the Git repo artifact provider is enabled. |
| accounts | [GitRepoArtifactAccount](#proto.GitRepoArtifactAccount) | repeated | The list of configured Git Repo accounts. |





 

 

 

 



<a name="google_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google_provider.proto



<a name="proto.Consul"></a>

### Consul
Configuration for Consul.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether Consul is enabled. |
| agentEndpoint | [string](#string) |  | Reachable Consul node endpoint connected to the Consul cluster. Defaults to localhost. |
| agentPort | [int32](#int32) |  | Port consul is running on for every agent. Defaults to 8500. |
| datacenters | [string](#string) | repeated | List of data centers to cache and keep updated. Defaults to all. |






<a name="proto.GoogleAccount"></a>

### GoogleAccount
Configuration for a Spinnaker Google account. An account maps to a
credential that can authenticate against a GCP project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated): List of required Fiat permission groups. Configure `permissions` instead. |
| permissions | [Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| project | [string](#string) |  | The GCP project this Spinnaker account will manage. |
| jsonPath | [string](#string) |  | The path to a JSON service account that Spinnaker will use as credentials. This is only needed if Spinnaker is not deployed on a Google Compute Engine VM, or needs permissions not afforded to the VM it is running on. See https://cloud.google.com/compute/docs/access/service-accounts for more information. |
| alphaListed | [bool](#bool) |  | Enable this flag if your GCP project has access to alpha features and you want Spinnaker to take advantage of them. |
| imageProjects | [string](#string) | repeated | A list of GCP projects from which Spinnaker will be able to cache and deploy images. When this is omitted, it defaults to the current project. Each project must have granted the IAM role compute.imageUser to the service account associated with the JSON key used by this account, as well as to the Google APIs service account automatically created for the project being managed (should look similar to 12345678912@cloudservices.gserviceaccount.com). See https://cloud.google.com/compute/docs/images/sharing-images-across-projects for more information about sharing images across GCP projects. |
| consul | [Consul](#proto.Consul) |  | Configuration for Consul. |
| regions | [string](#string) | repeated | A list of regions for caching and mutating calls. This overwrites any default regions set on the provider. |
| userDataFile | [string](#string) |  | The path to user data template file. Spinnaker has the ability to inject userdata into generated instance templates. The mechanism is via a template file that is token replaced to provide some specifics about the deployment. See https://github.com/spinnaker/clouddriver/blob/master/clouddriver-aws/UserData.md for more information. |






<a name="proto.GoogleBakeryDefaults"></a>

### GoogleBakeryDefaults
Configuration for Spinnaker&#39;s image bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| templateFile | [string](#string) |  | The name of the Packer template that will be used to bake images from this base image. The template file must be found in this list: https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or supplied as described here: https://spinnaker.io/setup/bakery/. |
| baseImages | [GoogleBaseImageSettings](#proto.GoogleBaseImageSettings) | repeated | List of configured base images. |
| zone | [string](#string) |  | The default zone in which to bake an image. |
| network | [string](#string) |  | The Google Compute network ID or URL to use for the launched instance. Defaults to default. |
| useInternalIp | [bool](#bool) |  | If true, use the instance&#39;s internal IP instead of its external IP during baking. |
| networkProjectId | [string](#string) |  | The default project ID for the network and subnet to use for the VM baking your image. |






<a name="proto.GoogleBaseImage"></a>

### GoogleBaseImage
Base image configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | This is the identifier used by GCP to find this base image. |
| shortDescription | [string](#string) |  | A short description to help human operators identify the image. |
| detailedDescription | [string](#string) |  | A long description to help human operators identify the image. |
| packageType | [string](#string) |  | This is used to help Spinnaker&#39;s bakery download the build artifacts you supply it with. For example, specifying deb indicates that your artifacts will need to be fetched from a debian repository. |
| imageFamily | [bool](#bool) |  | If set to true, Deck will annotate the popup tooltip to indicate that the selected option represents an image family. |






<a name="proto.GoogleBaseImageSettings"></a>

### GoogleBaseImageSettings
Configuration for a base image for the Google provider&#39;s bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseImage | [GoogleBaseImage](#proto.GoogleBaseImage) |  | Base image configuration. |
| virtualizationSettings | [VirtualizationSettings](#proto.VirtualizationSettings) |  | Image source configuration. |






<a name="proto.GoogleProvider"></a>

### GoogleProvider
Configuration for the Google Compute Engine (GCE) provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [GoogleAccount](#proto.GoogleAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |
| bakeryDefaults | [GoogleBakeryDefaults](#proto.GoogleBakeryDefaults) |  | Configuration for Spinnaker&#39;s image bakery. |






<a name="proto.VirtualizationSettings"></a>

### VirtualizationSettings
Image source configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sourceImage | [string](#string) |  | The source image. If both sourceImage and sourceImageFamily are set, sourceImage will take precedence. |
| sourceImageFamily | [string](#string) |  | The source image family to create the image from. The newest, non-deprecated image is used. If both sourceImage and sourceImageFamily are set, sourceImage will take precedence. |





 

 

 

 



<a name="halconfig.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## halconfig.proto



<a name="proto.HalConfig"></a>

### HalConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| persistentStorage | [HalConfig.PersistentStorage](#proto.HalConfig.PersistentStorage) |  |  |
| providers | [HalConfig.Providers](#proto.HalConfig.Providers) |  |  |
| artifacts | [ArtifactProviders](#proto.ArtifactProviders) |  |  |






<a name="proto.HalConfig.PersistentStorage"></a>

### HalConfig.PersistentStorage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| persistentStoreType | [string](#string) |  |  |
| gcs | [GCS](#proto.GCS) |  |  |






<a name="proto.HalConfig.Providers"></a>

### HalConfig.Providers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes | [KubernetesProvider](#proto.KubernetesProvider) |  |  |
| google | [GoogleProvider](#proto.GoogleProvider) |  |  |
| appengine | [AppengineProvider](#proto.AppengineProvider) |  |  |
| aws | [AwsProvider](#proto.AwsProvider) |  |  |
| azure | [AzureProvider](#proto.AzureProvider) |  |  |
| cloudfoundry | [CloudFoundryProvider](#proto.CloudFoundryProvider) |  |  |
| dcos | [DcosProvider](#proto.DcosProvider) |  |  |
| dockerRegistry | [DockerRegistryProvider](#proto.DockerRegistryProvider) |  |  |
| ecs | [EcsProvider](#proto.EcsProvider) |  |  |
| huaweicloud | [HuaweiCloudProvider](#proto.HuaweiCloudProvider) |  |  |
| oracle | [OracleProvider](#proto.OracleProvider) |  |  |





 

 

 

 



<a name="helm_artifact_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## helm_artifact_provider.proto



<a name="proto.HelmArtifactAccount"></a>

### HelmArtifactAccount
Configuration for a Helm artifact account. Either `username` and `password`
or `usernamePasswordFile` should be specified as means of authentication.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| repository | [string](#string) |  | The Helm chart repository URL. |
| username | [string](#string) |  | A username for Helm chart repository basic auth. |
| password | [string](#string) |  | A password for Helm chart repository basic auth. |
| usernamePasswordFile | [string](#string) |  | The path to a file containing the username and password for Helm chart repository basic auth. Must be in the format `${username}:${password}`. |






<a name="proto.HelmArtifactProvider"></a>

### HelmArtifactProvider
Configuration for the Helm artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the Helm artifact provider is enabled. |
| accounts | [HelmArtifactAccount](#proto.HelmArtifactAccount) | repeated | The list of configured Helm accounts. |





 

 

 

 



<a name="http_artifact_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## http_artifact_provider.proto



<a name="proto.HttpArtifactAccount"></a>

### HttpArtifactAccount
Configuration for an HTTP artifact account. Either `username` and `password`
or `usernamePasswordFile` should be specified as means of authentication.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| username | [string](#string) |  | A username for HTTP basic auth. |
| password | [string](#string) |  | A password for HTTP basic auth. |
| usernamePasswordFile | [string](#string) |  | The path to a file containing the username and password for HTTP basic auth. Must be in the format `${username}:${password}`. |






<a name="proto.HttpArtifactProvider"></a>

### HttpArtifactProvider
Configuration for the HTTP artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the HTTP artifact provider is enabled. |
| accounts | [HttpArtifactAccount](#proto.HttpArtifactAccount) | repeated | The list of configured HTTP accounts. |





 

 

 

 



<a name="huaweicloud_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## huaweicloud_provider.proto



<a name="proto.HuaweiCloudAccount"></a>

### HuaweiCloudAccount
Configuration for a Huawei Cloud account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| accountType | [string](#string) |  | The type of account. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated) List of required Fiat permission groups. Configure `permissions` instead. |
| permissions | [Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| authUrl | [string](#string) |  | (Required) The auth URL of the cloud. |
| domainName | [string](#string) |  | (Required) The domain name of the cloud. |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| insecure | [bool](#bool) |  | If `true`, disables certificate validation on SSL connections. Needed if certificates are self-signed. Defaults to `false`. |
| password | [string](#string) |  | (Required) The password used to access the cloud. |
| projectName | [string](#string) |  | (Required) The name of the project within the cloud. |
| regions | [string](#string) | repeated | (Required) The region(s) of the cloud. |
| username | [string](#string) |  | (Required) The username used to access the cloud. |






<a name="proto.HuaweiCloudBakeryDefaults"></a>

### HuaweiCloudBakeryDefaults
Configuration for Spinnaker&#39;s image bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseImages | [HuaweiCloudBaseImageSettings](#proto.HuaweiCloudBaseImageSettings) | repeated | List of configured base images. |
| templateFile | [string](#string) |  | This is the name of the packer template that will be used to bake images from this base image. The template file must be found in this list: https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or supplied as described here: https://spinnaker.io/setup/bakery/. |
| authUrl | [string](#string) |  | (Required) The default auth URL in which images will be baked. |
| username | [string](#string) |  | (Required) The default username with which images will be baked. |
| password | [string](#string) |  | (Required) The default password with which images will be baked. |
| projectName | [string](#string) |  | The name of the default project in which images will be baked. |
| domainName | [string](#string) |  | (Required) The default domain name in which images will be baked. |
| insecure | [bool](#bool) |  | The security setting for connecting to the Huawei Cloud account. Defaults to `false`. |
| vpcId | [string](#string) |  | (Required) The VPC in which images will be baked. |
| subnetId | [string](#string) |  | (Required) The subnet in which images will be baked. |
| securityGroup | [string](#string) |  | (Required) The default security group in which images will be baked. |
| eipBandwidthSize | [int32](#int32) |  | (Required) The bandwidth size of EIP in which images will be baked. |






<a name="proto.HuaweiCloudBaseImage"></a>

### HuaweiCloudBaseImage
Huawei Cloud base image settings.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | The name of the base image. |
| packageType | [string](#string) |  | This is used to help Spinnaker&#39;s bakery download the build artifacts you supply it with. For example, specifying `deb` indicates that your artifacts will need to be fetched from a debian repository. |
| templateFile | [string](#string) |  | This is the name of the packer template that will be used to bake images from this base image. The template file must be found in this list: https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or supplied as described here: https://spinnaker.io/setup/bakery/. |
| shortDescription | [string](#string) |  | A short description to help human operators identify the image. |
| detailedDescription | [string](#string) |  | A long description to help human operators identify the image. |






<a name="proto.HuaweiCloudBaseImageSettings"></a>

### HuaweiCloudBaseImageSettings
Configuration for a base image for the Huawei Cloud provider&#39;s bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseImage | [HuaweiCloudBaseImage](#proto.HuaweiCloudBaseImage) |  | Base image configuration. |
| virtualizationSettings | [HuaweiCloudVirtualizationSettings](#proto.HuaweiCloudVirtualizationSettings) | repeated | Image source configuration. |






<a name="proto.HuaweiCloudProvider"></a>

### HuaweiCloudProvider
Configuration for the Huawei Cloud provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [HuaweiCloudAccount](#proto.HuaweiCloudAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |
| bakeryDefaults | [HuaweiCloudBakeryDefaults](#proto.HuaweiCloudBakeryDefaults) |  | Configuration for Spinnaker&#39;s image bakery. |






<a name="proto.HuaweiCloudVirtualizationSettings"></a>

### HuaweiCloudVirtualizationSettings
Huawei Cloud virtualization settings.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| region | [string](#string) |  | (Required) The region for the baking configuration. |
| instanceType | [string](#string) |  | (Required) The instance type for the baking configuration. |
| sourceImageId | [string](#string) |  | (Required) The source image ID for the baking configuration. |
| sshUserName | [string](#string) |  | (Required) The SSH username for the baking configuration. |
| eipType | [string](#string) |  | (Required) The EIP type for the baking configuration. See the API doc to get its value. |





 

 

 

 



<a name="kubernetes_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## kubernetes_provider.proto



<a name="proto.KubernetesAccount"></a>

### KubernetesAccount
Configuration for a Spinnaker Kubernetes account. An account maps to a
credential that can authenticate against your Kubernetes cluster.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| providerVersion | [string](#string) |  | Some providers support multiple versions/release tracks. This allows you to pick the version of the provider (not the resources it manages) to run within Spinnaker. |
| kinds | [string](#string) | repeated | A list of resource kinds this Spinnaker account can deploy and will cache. When no kinds are configured, this defaults to all kinds described here: https://spinnaker.io/reference/providers/kubernetes-v2/. This can only be set when omitKinds is empty or not set. |
| omitKinds | [string](#string) | repeated | A list of resource kinds this Spinnaker account cannot deploy to or cache. This can only be set when kinds is empty or not set. |
| context | [string](#string) |  | The kubernetes context to be managed by Spinnaker. See http://kubernetes.io/docs/user-guide/kubeconfig-file/#context for more information. When no context is configured for an account the `current-context` in your kubeconfig is assumed. |
| cacheThreads | [int32](#int32) |  | Number of caching agents for this kubernetes account. Each agent handles a subset of the namespaces available to this account. By default, only 1 agent caches all kinds for all namespaces in the account. |
| namespaces | [string](#string) | repeated | A list of namespaces this Spinnaker account can deploy to and will cache. When no namespaces are configured, this defaults to all namespaces. |
| omitNamespaces | [string](#string) | repeated | A list of namespaces this Spinnaker account cannot deploy to or cache. This can only be set when namespaces is empty or not set. |
| customResources | [KubernetesCustomResource](#proto.KubernetesCustomResource) | repeated | The list of custom resources Clouddriver will manage and make available for use in Patch and Delete (Manifest) stages. |
| cachingPolicies | [KubernetesCachingPolicy](#proto.KubernetesCachingPolicy) | repeated | The list of kind-specific caching policies. |
| dockerRegistries | [KubernetesAccountDockerRegistry](#proto.KubernetesAccountDockerRegistry) | repeated | The list of the Spinnaker docker registry account names this Spinnaker account can use as image sources. These docker registry accounts must be registered in your halconfig before you can add them here. |
| oAuthScopes | [string](#string) | repeated | The list of OAuth scopes used by kubectl to fetch an OAuth token. |
| kubeconfigFile | [string](#string) |  | The path to your kubeconfig file. By default, it will be under the Spinnaker user&#39;s home directory in the typical .kube/config location. todo: document new var/secrets convention. |
| permissions | [Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated): List of required Fiat permission groups. Configure `permissions` instead. |






<a name="proto.KubernetesAccountDockerRegistry"></a>

### KubernetesAccountDockerRegistry
Configuration for a Docker registry.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accountName | [string](#string) |  | The configured name of the Docker registry. |
| namespaces | [string](#string) | repeated | The list of Docker registry namespaces usable as image sources. |






<a name="proto.KubernetesCachingPolicy"></a>

### KubernetesCachingPolicy
Configuration for a kind-specific caching policy.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetesKind | [string](#string) |  | The Kubernetes kind to which the policy applies. |
| maxEntriesPerAgent | [int32](#int32) |  | The maximum number of resources an agent will cache of the specified Kubernetes kind. |






<a name="proto.KubernetesCustomResource"></a>

### KubernetesCustomResource
Configuration for a CRD to be managed by Spinnaker. If Spinnaker does not
have permission to list CRDs but you need Spinnaker to manage CRDs, you
need to explicitly register each CRD.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetesKind | [string](#string) |  | The Kubernetes kind of the custom resource. |
| spinnakerKind | [string](#string) |  | The Spinnaker kind to which you would like the custom resource to map. |
| deployPriority | [string](#string) |  | An integer representing the deployment priority of this resource. Resources with lower values are deployed before resources with higher values. |
| versioned | [bool](#bool) |  | Whether Spinnaker should manage versioning this resource. |
| namespaced | [bool](#bool) |  | Whether the resource is namespaced. |






<a name="proto.KubernetesProvider"></a>

### KubernetesProvider
Configuration for the Kubernetes provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [KubernetesAccount](#proto.KubernetesAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |





 

 

 

 



<a name="maven_artifact_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## maven_artifact_provider.proto



<a name="proto.MavenArtifactAccount"></a>

### MavenArtifactAccount
Configuration for a Maven artifact account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| repositoryUrl | [string](#string) |  | (Required) Full URI for the Maven repository (e.g., `http://some.host.com/repository/path`). |






<a name="proto.MavenArtifactProvider"></a>

### MavenArtifactProvider
Configuration for the Maven artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the Maven artifact provider is enabled. |
| accounts | [MavenArtifactAccount](#proto.MavenArtifactAccount) | repeated | The list of configured Maven accounts. |





 

 

 

 



<a name="oracle_artifact_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## oracle_artifact_provider.proto



<a name="proto.OracleArtifactAccount"></a>

### OracleArtifactAccount
Configuration for an Oracle artifact account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| fingerprint | [string](#string) |  | The fingerprint of the public key. |
| namespace | [string](#string) |  | The namespace in which the bucket and objects should be created. |
| privateKeyPassphrase | [string](#string) |  | The passphrase used for the private key, if it is encrypted. |
| region | [string](#string) |  | An Oracle region (e.g., `us-phoenix-1`). |
| sshPrivateKeyFilePath | [string](#string) |  | Path to the private key in PEM format. |
| tenancyId | [string](#string) |  | The OCID of the Oracle Tenancy to use. |
| userId | [string](#string) |  | The OCID of the Oracle User with which to authenticate. |






<a name="proto.OracleArtifactProvider"></a>

### OracleArtifactProvider
Configuration for the Oracle artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the Oracle artifact provider is enabled. |
| accounts | [OracleArtifactAccount](#proto.OracleArtifactAccount) | repeated | The list of configured Oracle artifact accounts. |





 

 

 

 



<a name="oracle_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## oracle_provider.proto



<a name="proto.OracleAccount"></a>

### OracleAccount
Configuration for an Oracle account. An account maps to an Oracle Cloud
Infrastructure (OCI) user.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated) List of required Fiat permission groups. Configure `permissions` instead. |
| permissions | [Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| compartmentId | [string](#string) |  | (Required) The OCID of the Oracle Compartment to use. |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| fingerprint | [string](#string) |  | (Required) Fingerprint of the public key. |
| privateKeyPassphrase | [string](#string) |  | Passphrase used for the private key, if it is encrypted. |
| region | [string](#string) |  | (Required) An Oracle region (e.g., `us-phoenix-1`). |
| sshPrivateKeyFilePath | [string](#string) |  | (Required) Path to the private key in PEM format. |
| tenancyId | [string](#string) |  | (Required) The OCID of the Oracle Tenancy to use. |
| userId | [string](#string) |  | (Required) The OCID of the Oracle User with which to authenticate. |






<a name="proto.OracleBakeryDefaults"></a>

### OracleBakeryDefaults
Configuration for Spinnaker&#39;s image bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| templateFile | [string](#string) |  | The name of the Packer template that will be used to bake images from this base image. The template file must be found in this list: https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or supplied as described here: https://spinnaker.io/setup/bakery/. |
| baseImages | [OracleBaseImageSettings](#proto.OracleBaseImageSettings) | repeated | List of configured base images. |
| availabilityDomain | [string](#string) |  | (Required) The name of the Availability Domain within which a new instance is launched and provisioned. |
| instanceShape | [string](#string) |  | (Required) The shape for a newly created instance. |
| subnetId | [string](#string) |  | (Required) The name of the subnet within which a new instance is launched and provisioned. |






<a name="proto.OracleBaseImage"></a>

### OracleBaseImage
Oracle base image configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | The name of the base image. |
| shortDescription | [string](#string) |  | A short description to help human operators identify the image. |
| detailedDescription | [string](#string) |  | A long description to help human operators identify the image. |
| packageType | [string](#string) |  | This is used to help Spinnaker&#39;s bakery download the build artifacts you supply it with. For example, specifying deb indicates that your artifacts will need to be fetched from a debian repository. |
| templateFile | [string](#string) |  | The name of the Packer template that will be used to bake images from this base image. The template file must be found in this list: https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or supplied as described here: https://spinnaker.io/setup/bakery/. |






<a name="proto.OracleBaseImageSettings"></a>

### OracleBaseImageSettings
Configuration for a base image for the Oracle provider&#39;s bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseImage | [OracleBaseImage](#proto.OracleBaseImage) |  | Oracle base image configuration. |
| virtualizationSettings | [OracleVirtualizationSettings](#proto.OracleVirtualizationSettings) |  | Oracle virtualization settings. |






<a name="proto.OracleProvider"></a>

### OracleProvider
Configuration for the Oracle provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [OracleAccount](#proto.OracleAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |
| bakeryDefaults | [OracleBakeryDefaults](#proto.OracleBakeryDefaults) |  | Configuration for Spinnaker&#39;s image bakery. |






<a name="proto.OracleVirtualizationSettings"></a>

### OracleVirtualizationSettings
Oracle virtualization settings.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseImageId | [string](#string) |  | (Required) The OCID of the base image ID for the baking configuration. |
| sshUserName | [string](#string) |  | (Required) The ssh username for the baking configuration. |





 

 

 

 



<a name="permissions.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## permissions.proto



<a name="proto.Permissions"></a>

### Permissions
A Fiat permissions configuration object.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| READ | [string](#string) | repeated | A user must have at least one of these roles in order to view this account&#39;s cloud resources. |
| WRITE | [string](#string) | repeated | A user must have at least one of these roles in order to make changes to this account&#39;s cloud resources. |
| EXECUTE | [string](#string) | repeated | A user must have at least one of these roles in order to execute pipelines. |





 

 

 

 



<a name="persistent_storage.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## persistent_storage.proto



<a name="proto.GCS"></a>

### GCS



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| jsonPath | [string](#string) |  | A path to a JSON service account with permission to read and write to the bucket to be used as a backing store. |
| project | [string](#string) |  | The Google Cloud Platform project you are using to host the GCS bucket as a backing store. |
| bucket | [string](#string) |  | The name of a storage bucket that your specified account has access to. If not specified, a random name will be chosen. If you specify a globally unique bucket name that does not exist yet, Halyard will create that bucket for you. |
| rootFolder | [string](#string) |  | The root folder in the chosen bucket to place all of Spinnaker&#39;s persistent data in. |
| bucketLocation | [string](#string) |  | This is only required if the bucket you specify does not exist yet. |





 

 

 

 



<a name="s3_artifact_provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## s3_artifact_provider.proto



<a name="proto.S3ArtifactAccount"></a>

### S3ArtifactAccount
Configuration for an S3 artifact account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| apiEndpoint | [string](#string) |  | The S3 API endpoint. Only required when using an S3 clone such as Minio. |
| apiRegion | [string](#string) |  | The S3 API region. Only required when using an S3 clone such as Minio. |
| awsAccessKeyId | [string](#string) |  | The AWS Access Key ID. If not provided, Spinnaker will try to find AWS credentials as described at http://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/credentials.html#credentials-default. |
| awsSecretAccessKey | [string](#string) |  | The AWS Secret Key. |
| region | [string](#string) |  | The S3 region. |






<a name="proto.S3ArtifactProvider"></a>

### S3ArtifactProvider
Configuration for the S3 artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the S3 artifact provider is enabled. |
| accounts | [S3ArtifactAccount](#proto.S3ArtifactAccount) | repeated | The list of configured S3 artifact accounts. |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

