cmake_minimum_required(VERSION 3.5)

project(HarmonyProtocols)

include(GNUInstallDirs)

option(USE_STAGING_PROTOCOLS "Whether or not to install staging protocols")

add_subdirectory(stable)

if (USE_STAGING_PROTOCOLS)
    add_subdirectory(staging)
endif()
