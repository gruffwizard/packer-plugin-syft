packer {
  required_plugins {
    syft = {
      version = ">= 1.0.1"
      source  = "github.com/gruffwizard/syft"
    }
  }
}

source "null" "example" {
  communicator = "none"
}

build {
  sources = ["source.null.example"]

  post-processor "syft" {
    name = "World"
  }
}