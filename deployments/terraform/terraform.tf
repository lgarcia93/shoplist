provider "aws" {
  region                  = var.region
  shared_credentials_file = "~/.aws/credentials"
}

resource "aws_instance" "shoplistService" {
  
}
