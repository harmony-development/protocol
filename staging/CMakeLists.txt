foreach(dir IN ITEMS voice bots)
	install(DIRECTORY ${dir} DESTINATION ${CMAKE_INSTALL_INCLUDEDIR}/harmony-protocols FILES_MATCHING PATTERN "*.proto")
endforeach()
