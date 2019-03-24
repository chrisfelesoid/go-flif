cmake_minimum_required(VERSION 2.8.12)
set(CMAKE_COMPILER_IS_GNUCXX TRUE)
include(${CMAKE_CURRENT_SOURCE_DIR}/libs/flif/src/CMakeLists.txt)

add_library(flif_stat_lib STATIC ${COMMON_SOURCES} ${FLIF_SRC_DIR}/flif-enc.cpp ${FLIF_SRC_DIR}/library/flif-interface.cpp)

target_link_libraries(flif_stat_lib ${PNG_LIBRARY} ${STATIC_LINKED_LIBS})

set_target_properties(flif_stat_lib PROPERTIES OUTPUT_NAME flif)

target_compile_definitions(flif_stat_lib PRIVATE ${DEFINITIONS_FOR_ALL_TARGETS} FLIF_BUILD_DLL )

target_include_directories(flif_stat_lib PRIVATE ${FLIF_SRC_DIR}/../extern)

install(
    TARGETS flif_lib
    ARCHIVE DESTINATION lib
    LIBRARY DESTINATION lib
)
install(
    TARGETS flif_stat_lib
    ARCHIVE DESTINATION lib
    LIBRARY DESTINATION lib
)
install(DIRECTORY ${FLIF_SRC_DIR}/library/ DESTINATION include FILES_MATCHING PATTERN "*.h")