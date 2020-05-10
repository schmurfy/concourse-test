job "concourse-test" {
  datacenters = ["fra1"]
  
  type = "service"
  
  group "app" {
    count = 1
    
    task "gateway" {
      driver = "docker"
      
      config {
        image = "schmurfy/concourse-test-gateway:latest"
        force_pull = true
        
        dns_servers = ["${attr.unique.network.ip-address}"]


        port_map {
          http = 8888
        }
      }
      
      resources {
        cpu    = 50 # MHz
        memory = 20 # MB
        network {
          # mbits = 10
          port "http" {}
        }
      }
      
      service {
        name = "gateway-http"
        port = "http"
        
        address_mode = "driver"
        
        tags = [
          "traefik.tags=service",
          "traefik.http.routers.concourse-test.rule=Host(`app.xerxes.co.za`)",
          "traefik.http.routers.concourse-test.entryPoints=https",
          "traefik.http.routers.concourse-test.service=gateway-http",
          "traefik.http.routers.concourse-test.tls=true",
          "traefik.http.routers.concourse-test.tls.certresolver=myresolver",
        ]
      }
      
      env {
        LISTEN_ADDR = ":8888"
        SERVICE_ADDR = "service-grpc.service.consul:8000"
      }
    }
    
    
    task "service" {
      driver = "docker"
      
      config {
        image = "schmurfy/concourse-test-service:latest"
        force_pull = true
        
        dns_servers = ["${attr.unique.network.ip-address}"]
        
        port_map {
          grpc = 8001
        }
      }
      
      resources {
        cpu    = 50 # MHz
        memory = 20 # MB
        network {
          # mbits = 10
          port "grpc" {}
        }
      }
      
      service {
        name = "service-grpc"
        port = "grpc"
        address_mode = "driver"
      }
      
      env {
        LISTEN_ADDR = ":8001"
        REDIS_ADDR = "redis.service.consul:6379"
      }
    }

    
  }
}
