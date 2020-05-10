cmake_minimum_required(VERSION 3.15)
project(embedded)

set(CMAKE_CXX_STANDARD 20)

include_directories("./includes/")

add_executable(embedded
        main.cpp
        includes/httplib.h
        includes/json.hpp
        )