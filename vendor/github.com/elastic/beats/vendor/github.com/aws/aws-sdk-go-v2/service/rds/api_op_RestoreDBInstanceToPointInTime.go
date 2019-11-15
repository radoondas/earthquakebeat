// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RestoreDBInstanceToPointInTimeMessage
type RestoreDBInstanceToPointInTimeInput struct {
	_ struct{} `type:"structure"`

	// Indicates that minor version upgrades are applied automatically to the DB
	// instance during the maintenance window.
	AutoMinorVersionUpgrade *bool `type:"boolean"`

	// The Availability Zone (AZ) where the DB instance will be created.
	//
	// Default: A random, system-chosen Availability Zone.
	//
	// Constraint: You can't specify the AvailabilityZone parameter if the MultiAZ
	// parameter is set to true.
	//
	// Example: us-east-1a
	AvailabilityZone *string `type:"string"`

	// True to copy all tags from the restored DB instance to snapshots of the restored
	// DB instance, and otherwise false. The default is false.
	CopyTagsToSnapshot *bool `type:"boolean"`

	// The compute and memory capacity of the Amazon RDS DB instance, for example,
	// db.m4.large. Not all DB instance classes are available in all AWS Regions,
	// or for all database engines. For the full list of DB instance classes, and
	// availability for your engine, see DB Instance Class (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)
	// in the Amazon RDS User Guide.
	//
	// Default: The same DBInstanceClass as the original DB instance.
	DBInstanceClass *string `type:"string"`

	// The database name for the restored DB instance.
	//
	// This parameter is not used for the MySQL or MariaDB engines.
	DBName *string `type:"string"`

	// The name of the DB parameter group to associate with this DB instance. If
	// this argument is omitted, the default DBParameterGroup for the specified
	// engine is used.
	//
	// Constraints:
	//
	//    * If supplied, must match the name of an existing DBParameterGroup.
	//
	//    * Must be 1 to 255 letters, numbers, or hyphens.
	//
	//    * First character must be a letter.
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens.
	DBParameterGroupName *string `type:"string"`

	// The DB subnet group name to use for the new instance.
	//
	// Constraints: If supplied, must match the name of an existing DBSubnetGroup.
	//
	// Example: mySubnetgroup
	DBSubnetGroupName *string `type:"string"`

	// Indicates if the DB instance should have deletion protection enabled. The
	// database can't be deleted when this value is set to true. The default is
	// false. For more information, see Deleting a DB Instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).
	DeletionProtection *bool `type:"boolean"`

	// Specify the Active Directory Domain to restore the instance in.
	Domain *string `type:"string"`

	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service.
	DomainIAMRoleName *string `type:"string"`

	// The list of logs that the restored DB instance is to export to CloudWatch
	// Logs. The values in the list depend on the DB engine being used. For more
	// information, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon RDS User Guide.
	EnableCloudwatchLogsExports []string `type:"list"`

	// True to enable mapping of AWS Identity and Access Management (IAM) accounts
	// to database accounts, and otherwise false.
	//
	// You can enable IAM database authentication for the following database engines
	//
	//    * For MySQL 5.6, minor version 5.6.34 or higher
	//
	//    * For MySQL 5.7, minor version 5.7.16 or higher
	//
	// Default: false
	EnableIAMDatabaseAuthentication *bool `type:"boolean"`

	// The database engine to use for the new instance.
	//
	// Default: The same as source
	//
	// Constraint: Must be compatible with the engine of the source
	//
	// Valid Values:
	//
	//    * mariadb
	//
	//    * mysql
	//
	//    * oracle-ee
	//
	//    * oracle-se2
	//
	//    * oracle-se1
	//
	//    * oracle-se
	//
	//    * postgres
	//
	//    * sqlserver-ee
	//
	//    * sqlserver-se
	//
	//    * sqlserver-ex
	//
	//    * sqlserver-web
	Engine *string `type:"string"`

	// The amount of Provisioned IOPS (input/output operations per second) to be
	// initially allocated for the DB instance.
	//
	// Constraints: Must be an integer greater than 1000.
	//
	// SQL Server
	//
	// Setting the IOPS value for the SQL Server database engine is not supported.
	Iops *int64 `type:"integer"`

	// License model information for the restored DB instance.
	//
	// Default: Same as source.
	//
	// Valid values: license-included | bring-your-own-license | general-public-license
	LicenseModel *string `type:"string"`

	// Specifies if the DB instance is a Multi-AZ deployment.
	//
	// Constraint: You can't specify the AvailabilityZone parameter if the MultiAZ
	// parameter is set to true.
	MultiAZ *bool `type:"boolean"`

	// The name of the option group to be used for the restored DB instance.
	//
	// Permanent options, such as the TDE option for Oracle Advanced Security TDE,
	// can't be removed from an option group, and that option group can't be removed
	// from a DB instance once it is associated with a DB instance
	OptionGroupName *string `type:"string"`

	// The port number on which the database accepts connections.
	//
	// Constraints: Value must be 1150-65535
	//
	// Default: The same port as the original DB instance.
	Port *int64 `type:"integer"`

	// The number of CPU cores and the number of threads per core for the DB instance
	// class of the DB instance.
	ProcessorFeatures []ProcessorFeature `locationNameList:"ProcessorFeature" type:"list"`

	// Specifies the accessibility options for the DB instance. A value of true
	// specifies an Internet-facing instance with a publicly resolvable DNS name,
	// which resolves to a public IP address. A value of false specifies an internal
	// instance with a DNS name that resolves to a private IP address. For more
	// information, see CreateDBInstance.
	PubliclyAccessible *bool `type:"boolean"`

	// The date and time to restore from.
	//
	// Valid Values: Value must be a time in Universal Coordinated Time (UTC) format
	//
	// Constraints:
	//
	//    * Must be before the latest restorable time for the DB instance
	//
	//    * Can't be specified if UseLatestRestorableTime parameter is true
	//
	// Example: 2009-09-07T23:45:00Z
	RestoreTime *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The identifier of the source DB instance from which to restore.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing DB instance.
	SourceDBInstanceIdentifier *string `type:"string"`

	// The resource ID of the source DB instance from which to restore.
	SourceDbiResourceId *string `type:"string"`

	// Specifies the storage type to be associated with the DB instance.
	//
	// Valid values: standard | gp2 | io1
	//
	// If you specify io1, you must also include a value for the Iops parameter.
	//
	// Default: io1 if the Iops parameter is specified, otherwise standard
	StorageType *string `type:"string"`

	// A list of tags. For more information, see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	Tags []Tag `locationNameList:"Tag" type:"list"`

	// The name of the new DB instance to be created.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 letters, numbers, or hyphens
	//
	//    * First character must be a letter
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens
	//
	// TargetDBInstanceIdentifier is a required field
	TargetDBInstanceIdentifier *string `type:"string" required:"true"`

	// The ARN from the key store with which to associate the instance for TDE encryption.
	TdeCredentialArn *string `type:"string"`

	// The password for the given ARN from the key store in order to access the
	// device.
	TdeCredentialPassword *string `type:"string"`

	// A value that specifies that the DB instance class of the DB instance uses
	// its default processor features.
	UseDefaultProcessorFeatures *bool `type:"boolean"`

	// Specifies whether (true) or not (false) the DB instance is restored from
	// the latest backup time.
	//
	// Default: false
	//
	// Constraints: Can't be specified if RestoreTime parameter is provided.
	UseLatestRestorableTime *bool `type:"boolean"`

	// A list of EC2 VPC security groups to associate with this DB instance.
	//
	// Default: The default EC2 VPC security group for the DB subnet group's VPC.
	VpcSecurityGroupIds []string `locationNameList:"VpcSecurityGroupId" type:"list"`
}

// String returns the string representation
func (s RestoreDBInstanceToPointInTimeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreDBInstanceToPointInTimeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestoreDBInstanceToPointInTimeInput"}

	if s.TargetDBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetDBInstanceIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RestoreDBInstanceToPointInTimeResult
type RestoreDBInstanceToPointInTimeOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon RDS DB instance.
	//
	// This data type is used as a response element in the DescribeDBInstances action.
	DBInstance *DBInstance `type:"structure"`
}

// String returns the string representation
func (s RestoreDBInstanceToPointInTimeOutput) String() string {
	return awsutil.Prettify(s)
}

const opRestoreDBInstanceToPointInTime = "RestoreDBInstanceToPointInTime"

// RestoreDBInstanceToPointInTimeRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Restores a DB instance to an arbitrary point in time. You can restore to
// any point in time before the time identified by the LatestRestorableTime
// property. You can restore to a point up to the number of days specified by
// the BackupRetentionPeriod property.
//
// The target database is created with most of the original configuration, but
// in a system-selected Availability Zone, with the default security group,
// the default subnet group, and the default DB parameter group. By default,
// the new DB instance is created as a single-AZ deployment except when the
// instance is a SQL Server instance that has an option group that is associated
// with mirroring; in this case, the instance becomes a mirrored deployment
// and not a single-AZ deployment.
//
// This command doesn't apply to Aurora MySQL and Aurora PostgreSQL. For Aurora,
// use RestoreDBClusterToPointInTime.
//
//    // Example sending a request using RestoreDBInstanceToPointInTimeRequest.
//    req := client.RestoreDBInstanceToPointInTimeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RestoreDBInstanceToPointInTime
func (c *Client) RestoreDBInstanceToPointInTimeRequest(input *RestoreDBInstanceToPointInTimeInput) RestoreDBInstanceToPointInTimeRequest {
	op := &aws.Operation{
		Name:       opRestoreDBInstanceToPointInTime,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestoreDBInstanceToPointInTimeInput{}
	}

	req := c.newRequest(op, input, &RestoreDBInstanceToPointInTimeOutput{})
	return RestoreDBInstanceToPointInTimeRequest{Request: req, Input: input, Copy: c.RestoreDBInstanceToPointInTimeRequest}
}

// RestoreDBInstanceToPointInTimeRequest is the request type for the
// RestoreDBInstanceToPointInTime API operation.
type RestoreDBInstanceToPointInTimeRequest struct {
	*aws.Request
	Input *RestoreDBInstanceToPointInTimeInput
	Copy  func(*RestoreDBInstanceToPointInTimeInput) RestoreDBInstanceToPointInTimeRequest
}

// Send marshals and sends the RestoreDBInstanceToPointInTime API request.
func (r RestoreDBInstanceToPointInTimeRequest) Send(ctx context.Context) (*RestoreDBInstanceToPointInTimeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreDBInstanceToPointInTimeResponse{
		RestoreDBInstanceToPointInTimeOutput: r.Request.Data.(*RestoreDBInstanceToPointInTimeOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreDBInstanceToPointInTimeResponse is the response type for the
// RestoreDBInstanceToPointInTime API operation.
type RestoreDBInstanceToPointInTimeResponse struct {
	*RestoreDBInstanceToPointInTimeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreDBInstanceToPointInTime request.
func (r *RestoreDBInstanceToPointInTimeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
