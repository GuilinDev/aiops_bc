from typing import Any

def modify_config(service_name: str, key: str, value: Any) -> str:
    """
    修改服务配置
    
    Args:
        service_name: 服务名称
        key: 配置键
        value: 配置值
    
    Returns:
        str: 操作结果信息
    """
    return f"已将服务 {service_name} 的 {key} 修改为 {value}"

def restart_service(service_name: str) -> str:
    """
    重启服务
    
    Args:
        service_name: 服务名称
    
    Returns:
        str: 操作结果信息
    """
    return f"服务 {service_name} 已重启"

def apply_manifest(resource_type: str, image: str) -> str:
    """
    部署资源
    
    Args:
        resource_type: 资源类型（如deployment, service等）
        image: 容器镜像
    
    Returns:
        str: 操作结果信息
    """
    return f"已部署 {resource_type}，使用镜像 {image}" 
