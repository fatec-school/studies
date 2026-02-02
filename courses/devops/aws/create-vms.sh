#!/bin/bash

KEY_NAME="my-ssh-key"
SG_NAME="my-security-group"

VPC_ID=$(aws ec2 describe-vpcs --filters "Name=isDefault,Values=true" --query 'Vpcs[0].VpcId' --output text)

SUBNET_ID=$(aws ec2 describe-subnets --filters "Name=default-for-az,Values=true" --query 'Subnets[0].SubnetId' --output text)

AMI_ID=$(aws ec2 describe-images --owners amazon --filters "Name=name,Values=amzn2-ami-hvm-*-x86_64-gp2" "Name=state,Values=available" --query 'Images | sort_by(@, &CreationDate)[-1].ImageId' --output text)


# Check if AWS CLI is installed
if ! command -v aws &> /dev/null; then
    echo "AWS CLI could not be found. Please install it to proceed."
    exit 1
fi

# Check if the VPC ID was retrieved successfully
if [ -z "$VPC_ID" ]; then
    echo "Failed to retrieve VPC ID. Exiting..."
    exit 1
fi

# Check if the Subnet ID was retrieved successfully
if [ -z "$SUBNET_ID" ]; then
    echo "Failed to retrieve Subnet ID. Exiting..."
    exit 1
fi

# Check if the AMI ID was retrieved successfully
if [ -z "$AMI_ID" ]; then
    echo "Failed to retrieve AMI ID. Exiting..."
    exit 1
fi

# Verify if the key pair already exists
aws ec2 describe-key-pairs --key-names "$KEY_NAME" > /dev/null 2>&1

if [ $? -ne 0 ]; then
    aws ec2 create-key-pair --key-name "$KEY_NAME" --query 'KeyMaterial' --output text > "${KEY_NAME}.pem"
    chmod 400 "${KEY_NAME}.pem"
    echo "Created key pair '$KEY_NAME' and saved to '${KEY_NAME}.pem'."
else
    echo "Key pair '$KEY_NAME' already exists."
fi

# Create security group
SG_ID=$(aws ec2 create-security-group --group-name "$SG_NAME" --vpc-id "$VPC_ID" --query 'GroupId' --output text)

# Allow 22 port (SSH) 
aws ec2 authorize-security-group-ingress --group_id "$SG_ID" --protocol tcp --port 22 --cidr 0.0.0.0/0 

# Launch EC2 instance
INSTANCE_ID=$(aws ec2 run-instances --image-id "$AMI_ID" --count 1 --instance-type t2.micro --key-name "$KEY_NAME" --security-group-ids "$SG_NAME" --subnet-id "$SUBNET_ID" --query 'Instances[0].InstanceId' --output text)

if [ $? -eq 0 ]; then
    echo "EC2 instance launched with ID: $INSTANCE_ID"
else
    echo "Failed to launch EC2 instance."
    exit 1
fi