variable "aws_region" {
  default = "us-east-1"
}

variable "ami_id" {
  default = "ami-0aa7d40eeae50c9a9"  # Amazon Linux 2023 AMI (HVM), SSD Volume Type
}

variable "instance_type" {
  default = "t3.micro"
}

variable "key_name" {
  description = "Name of the SSH key pair to use for the instance"
  type        = string
}

variable "redis_cluster_id" {
  default = "my-redis-cluster"
}

variable "redis_node_type" {
  default = "cache.t3.micro"
}

variable "redis_num_cache_nodes" {
  default = 1
}

variable "redis_parameter_group_name" {
  default = "default.redis7"
}
