foreach(dir IN ITEMS auth channels harmonytypes mediaproxy emote profile batch sync)
	install(DIRECTORY ${dir} DESTINATION ${CMAKE_INSTALL_INCLUDEDIR}/harmony-protocols FILES_MATCHING PATTERN "*.proto")
endforeach()
