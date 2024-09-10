# AWS Infrastructure with Terraform

This project uses Terraform to provision an EC2 instance and an ElastiCache Redis cluster on AWS.

## Resources Created

1. EC2 Instance
   - Used for running Docker containers
   - Instance type: t3.micro (configurable)
   - AMI: Amazon Linux 2023

![EC2 Instance](./images/ec2.png)

2. ElastiCache Redis Cluster
   - Single-node Redis cluster
   - Node type: cache.t3.micro (configurable)

![ElastiCache Redis Cluster](./images/redis.png)

3. Security Group
   - Allows SSH access to the EC2 instance

## Prerequisites

- AWS CLI installed and configured
- Terraform installed
- An AWS account with appropriate permissions

## Configuration

The main configuration is done through the following files:

- `main.tf`: Contains the main Terraform configuration
- `variables.tf`: Defines input variables
- `terraform.tfvars`: Sets values for the input variables (create this file if it doesn't exist)

## Usage

1. Initialize Terraform:
   ```
   terraform init
   ```

2. Review the planned changes:
   ```
   terraform plan
   ```

3. Apply the configuration:
   ```
   terraform apply
   ```

![Terraform Apply](./images/terraform_apply.png)

4. To destroy the resources:
   ```
   terraform destroy
   ```

## Resource Lifecycle

By default, the EC2 instance and Redis cluster will continue running until explicitly destroyed. They will not automatically terminate after 2 hours.

To automatically destroy resources after a set time:

1. Use a scheduled task or cron job to run `terraform destroy` after the desired duration.
2. Implement a custom solution using AWS Lambda and CloudWatch Events to terminate resources.
3. Use AWS Auto Scaling group for EC2 with a scheduled action to terminate instances.

Remember, running resources incur costs. Always monitor your AWS usage and destroy resources when they're no longer needed.

## Accessing Resources

- EC2 Instance: Use SSH with the key pair specified in the configuration
- Redis Cluster: Access details can be found in the AWS ElastiCache console or Terraform output

## Notes

- Ensure your AWS credentials are properly configured
- Review and adjust the security group settings as needed
- Monitor your AWS usage to avoid unexpected costs

