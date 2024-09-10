output "instance_public_ip" {
  description = "Public IP address of the EC2 instance"
  value       = aws_instance.docker_instance.public_ip
}

output "redis_cache_nodes" {
  description = "Redis cache nodes"
  value       = aws_elasticache_cluster.redis.cache_nodes
}
