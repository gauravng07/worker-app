#!/usr/bin/env bash

main() {
  install_go_libraries
  build
  echo "========   DONE   ========"
}

build() {
  echo ===================================
  echo "         Building App            "
  echo ===================================
  make build
}

install_go_libraries() {
  echo ===================================
  echo "   Installing GO dependencies    "
  echo ===================================

  ## Use install_dep function to install the libraries
  go mod download
}

main
