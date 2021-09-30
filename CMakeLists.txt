cmake_minimum_required(VERSION 3.5)

project(HarmonyProtocols)

include(GNUInstallDirs)

foreach(dir IN ITEMS auth chat harmonytypes mediaproxy voice)
	install(DIRECTORY ${dir} DESTINATION ${CMAKE_INSTALL_INCLUDEDIR}/harmony-protocols FILES_MATCHING PATTERN "*.proto")
endforeach()
